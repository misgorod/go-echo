package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	http.HandleFunc("/", EchoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
