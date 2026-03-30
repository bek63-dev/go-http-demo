package main

import (
	"log"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8080"

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	log.Printf("Запуск клиента (timeout: %v). Целевой путь к серверу: %s\n",
		client.Timeout, baseURL)
	executeWork(client, baseURL)
}
