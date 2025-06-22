package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID      int
	Name    string
	Balance float32
}

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	return conn, err
}

func performQuery(conn *pgx.Conn, query string) {
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(conn *pgx.Conn, name string, balance float32) {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (name, balance) VALUES ($1, $2)", name, balance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s created with balance %f\n\n", name, balance)
}

func getUserById(conn *pgx.Conn, id int) User {
	var user User
	err := conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User ID: %d, Name: %s, Balance: %f\n", user.ID, user.Name, user.Balance)
	return user
}

func updateBalance(conn *pgx.Conn, name string, balance float32) {
	_, err := conn.Exec(context.Background(), "UPDATE users SET balance = $1 WHERE name = $2", balance, name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s balance updated to %f\n\n", name, balance)
}

func deleteUser(conn *pgx.Conn, name string) {
	_, err := conn.Exec(context.Background(), "DELETE FROM users WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s deleted\n\n", name)
}

func queryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ID int
		var Name string
		var Balance float32
		if err := rows.Scan(&ID, &Name, &Balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s, Balance: %f\n", ID, Name, Balance)
	}
}

func main() {
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	// performQuery(conn, "CREATE TABLE IF NOT EXISTS users (id serial primary key, name text, balance real)")
	// performQuery(conn, "INSERT INTO users (name, balance) VALUES ('John Doe', 1000)")
	// performQuery(conn, "INSERT INTO users (name, balance) VALUES ('Jane Doe', 2000)")
	// performQuery(conn, "INSERT INTO users (name, balance) VALUES ('Andrew Black', 3000)")
	createUser(conn, "Amy White", 7000)
	defer conn.Close(context.Background())
	queryData(conn)
}
