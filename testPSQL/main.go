package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type City struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	State string `db:"state"`
}

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository(db *sqlx.DB) *CityRepository {

	return &CityRepository{db: db}
}

func (r *CityRepository) Create(city *City) error {

	query := `INSERT INTO cities (name, state) VALUES ($1, $2)`
	_, err := r.db.Exec(query, city.Name, city.State)
	fmt.Println("create city")
	return err
}

func (r *CityRepository) Delete(id int) error {
	query := `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CityRepository) Update(city *City) error {
	query := `UPDATE cities SET name = $1, state = $2 WHERE id = $3`
	_, err := r.db.Exec(query, city.Name, city.State, city.Id)
	return err
}

func (r *CityRepository) List() ([]City, error) {
	var cities []City
	fmt.Println("проверка кеша")

	fmt.Println("вывод из бд")
	query := `SELECT * FROM cities`
	err := r.db.Select(&cities, query)
	if err != nil {
		return nil, err
	}

	return cities, nil

}

func CreateTables(db *sqlx.DB) error {
	fmt.Println("table create")
	rows, err := db.Query(`SELECT * FROM cities LIMIT 1`)
	if err == nil {
		rows.Close()
		return nil
	}
	_, err = db.Exec(`CREATE TABLE cities (id SERIAL PRIMARY KEY,name VARCHAR(30) NOT NULL,state VARCHAR(30) NOT NULL)`)
	if err != nil {
		log.Println(err)
	}

	return err
}

func main() {

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "users"
	dbPassword := "secret"
	dbName := "postgres"
	sslmode := "disable"

	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslmode
	fmt.Println("start connect")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = CreateTables(db)
	log.Println(err)
	repo := NewCityRepository(db)

	city := City{
		Id:    1,
		Name:  "Snk",
		State: "Prs",
	}

	err = repo.Create(&city)
	if err != nil {
		log.Println(err)
	}
	list, err := repo.List()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(list)

	list1, err := repo.List()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(list1)

}
