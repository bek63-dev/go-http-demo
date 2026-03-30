package main

import (
	"log"
	"net/http"
)

const serverURL = "localhost:8080"

func main() {
	log.Println("\nСервер запущен на:", serverURL)
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/time", handleTime)
	mux.HandleFunc("/date", handleDate)
	log.Fatal(http.ListenAndServe(serverURL, mux))
}
