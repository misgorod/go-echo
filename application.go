package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("VERSION 2"))
}

func main() {
	connStr := "user=postgres password=Yungleansoset!1488 dbname=pqgotest host=echogodb.cgpygcvzbwp1.eu-central-1.rds.amazonaws.com port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	http.HandleFunc("/", EchoHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
