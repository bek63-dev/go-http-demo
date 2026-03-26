package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const serverURL = "http://localhost:8080"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	log.Printf("Запуск клиента (timeout: %v). Целевой сервер: %s\n",
		client.Timeout, serverURL)
	executeWork(client, serverURL)
}
