package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	Id int
	Name string
}
// psql -h 127.0.0.1 -p 5433 -U postgres test_db
func main() {

	log.Fatal(http.ListenAndServe(":8000", nil))

	db, err := sql.Open("postgres", "user=myuser password=mypassword dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("データベースの接続に成功しました")
		log.Fatal(err)
	} else {
		fmt.Println("データベースの接続に成功しました")
	}
	defer db.Close()

	// db.Query()
}
