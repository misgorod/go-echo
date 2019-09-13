package main

import (
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("VERSION 2"))
}

func main() {
	http.HandleFunc("/", EchoHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
