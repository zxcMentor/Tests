package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
)

/*
Data Layer
*/
type Repository struct {
	db      *sqlx.DB
	builder squirrel.StatementBuilderType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Repository) AddUser(user User) error {
	query, args, _ := r.builder.
		Select("COUNT(*)").
		From("users").
		Where(squirrel.Eq{"email": user.Email}).
		ToSql()

	var count int
	err := r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email уже используется")
	}
	if user.Age < 18 {
		return errors.New("возраст пользователя должен быть не меньше 18 лет")
	}

	_, args, _ = r.builder.
		Insert("users").
		Columns("email", "password", "name", "age").
		Values(user.Email, user.Password, user.Name, user.Age).
		ToSql()

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *Repository) GetAllUsers() ([]User, error) {
	query, _, _ := r.builder.
		Select("email", "name", "age").
		From("users").
		ToSql()

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Email, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

/*

	Business Logic Layer

*/

type Service interface {
	GetUser(email string) User
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type RealService struct {
	repo *Repository
}

func (rs *RealService) GetUser(email string) User {

	return User{ /*что то там пароли логины*/ }
}

type Cache struct {
	store map[string]User
}

func NewCache() *Cache {
	return &Cache{store: make(map[string]User)}
}

func (c *Cache) Get(email string) (User, bool) {
	user, found := c.store[email]
	return user, found
}

func (c *Cache) Set(email string, user User) {
	c.store[email] = user
}

type CacheProxy struct {
	realService Service
	cache       *Cache
}

func NewCacheProxy(service Service) *CacheProxy {
	return &CacheProxy{
		realService: service,
		cache:       NewCache(),
	}
}

func (p *CacheProxy) GetUser(email string) User {

	return User{}
}

/*

 Presentation Layer

*/

func NewRouter(service Service) *http.ServeMux {
	router := http.NewServeMux()

	http.HandleFunc("/getUser", getUserHandler(service))
	http.HandleFunc("/updateUser", updateUserHandler(service))

	return router
}

func getUserHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		user := service.GetUser(email)
		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	}
}

func updateUserHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// логика обработки обновления пользователя

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User updated successfully")
	}
}

/*

 main

*/

func main() {

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=viva dbname=postgres password=1234 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := NewRepository(db)
	realService := &RealService{repo: repo}
	cacheProxy := NewCacheProxy(realService)

	router := NewRouter(cacheProxy)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
