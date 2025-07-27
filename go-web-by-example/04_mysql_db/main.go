package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:00000000@(127.0.0.1:3306)/go_test?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("open db error: %v", err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Printf("ping error: %v", err)
	}

	{
		// create a new table
		query := `
			CREATE TABLE IF NOT EXISTS users (
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			);
		`
		if _, err := db.Exec(query); err != nil {
			log.Printf("create table error: %v", err)
			return
		}
	}

	{
		// insert a new user
		username := "johndoe"
		password := "secret"
		createAt := time.Now()

		result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", username, password, createAt)
		if err != nil {
			log.Printf("insert sql error: %v", err)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("insert error: %v", err)
			return
		}
		fmt.Println(id)
	}

	{
		// query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)
		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Printf("query error: %v", err)
			return
		}

		fmt.Println(id, username, password, createdAt)
	}

	{
		// query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query("SELECT id, username, password, created_at FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("users: %+v\n", users)
	}

	{
		_, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
