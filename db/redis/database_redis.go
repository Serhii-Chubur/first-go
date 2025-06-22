package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID      int
	Name    string
	Balance float64
}

func main() {
	log.SetFlags(log.Lshortfile)

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// 	Protocol: 2,
	// })

	conn, err := redis.ParseURL("redis://redis:@localhost:6379/0")
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(conn)

	ctx := context.Background()
	// err = client.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// val, err := client.Get(ctx, "key").Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("key", val)

	var user User = User{ID: 1, Name: "John Doe", Balance: 1000}

	// data, err := json.Marshal(user)
	// fmt.Println(string(data))

	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]interface{}{
		"ID":      strconv.Itoa(user.ID),
		"Name":    user.Name,
		"Balance": fmt.Sprintf("%f", user.Balance), // strconv.FormatFloat(user.Balance),
	}

	err = client.HMSet(ctx, "h_user", fields).Err()
	if err != nil {
		log.Fatal(err)
	}

	// err = client.Set(ctx, "user", data, 0).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	val, err := client.HGet(ctx, "h_user", "Balance").Result()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("User ID: %d, Name: %s, Balance: %f\n", user.ID, user.Name, user.Balance)

	fmt.Println(val)
	fmt.Println("")

	var result User
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User ID: %d, Name: %s, Balance: %f\n", result.ID, result.Name, result.Balance)
}
