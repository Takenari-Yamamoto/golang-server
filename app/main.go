package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	id int
	name string
}

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Hello World from Go.")
}

func main() {

	db, err := sqlx.Connect("postgres", "user=username password=password dbname=dbname sslmode=disable")
	if err != nil {
			fmt.Println("データベースの接続に失敗しました")
			log.Fatal(err)
	}

	 // set the connection for boil
    boil.SetDB(db)

    // http.HandleFunc("/hello", handler)

    // 8080ポートで起動
    http.ListenAndServe(":8000", nil)

	// if err != nil {
	// 	fmt.Println("HTTPサーバーが起動しました")
	// 	http.HandleFunc("/", handler)
	// } else {
	// 	fmt.Println("エラーです")
	// }

}
