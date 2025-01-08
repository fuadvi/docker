package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

var ctx = context.Background()

func Handler(w http.ResponseWriter, r *http.Request) {
	val, err := client.Get(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Count: %s", val)
}

func AddCounter(w http.ResponseWriter, r *http.Request) {
	_, err := client.Incr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
}

func ResetCounter(w http.ResponseWriter, r *http.Request) {
	err := client.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
}

func SetupRedistClient() {
	fmt.Println("Executing redis client")
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})
}

func SetKey(key string, value int) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	SetupRedistClient()
	SetKey("counter", 1)

	http.HandleFunc("/", Handler)
	http.HandleFunc("/add", AddCounter)
	http.HandleFunc("/reset", ResetCounter)
	http.ListenAndServe(":8080", nil)
}
