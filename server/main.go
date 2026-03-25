package main

import (
	"log"
	"net/http"
)

const serverURL = "localhost:8080"

func main() {
	log.Println("\nСервер запущен на:", serverURL)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(serverURL, nil))
	log.Println("Ошибка в запуске сервера")
}
