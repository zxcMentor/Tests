package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"runtime"
)

// incrementLikes увеличивает счетчик лайков на 1.
// заметьте что передается context, который мы будем использовать для отмены операции
// в случае ошибки, либо завершения контекста.
// также заметьте, что мы не возвращаем значение, а только ошибку.
// это потому что в Redis операция увеличения счетчика возвращает новое значение счетчика.
// мы можем использовать это значение, но в данном случае мы не будем этого делать.
func incrementLikes(ctx context.Context, client *redis.Client, photoID int) error {
	key := "photo:" + string(photoID) + ":likes"
	_, err := client.Incr(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}

// getLikes возвращает количество лайков для фотографии.
// заметьте, что мы возвращаем значение и ошибку.
// это потому что в Redis операция получения значения возвращает значение и ошибку.
func getLikes(ctx context.Context, client *redis.Client, photoID int) (int64, error) {
	key := "photo:" + string(photoID) + ":likes"
	likes, err := client.Get(ctx, key).Int64()
	if err != nil {
		return 0, err
	}

	return likes, nil
}

func main() {
	// Создаем клиент Redis.
	// Создаем несколько горутин
	for i := 0; i < 1000; i++ {
		go func() {
			// Здесь может быть ваш код
		}()
	}

	// Позволяем горутинам выполниться
	runtime.Gosched()

	// Выводим количество горутин
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())
}
