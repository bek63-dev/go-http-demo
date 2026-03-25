package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	const serverURL = "http://localhost:8080"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	log.Printf("Запуск клиента (timeout : %v). Целевой сервер: %s\n",
		client.Timeout, serverURL)
	ExecuteWork(client, serverURL)
}

func ExecuteWork(client *http.Client, serverURL string) {
	// Проверка соединения с сервером
	log.Printf("Проверка соединения с %s\n", serverURL)

	if err := PingServer(client, serverURL); err != nil {
		log.Printf("ошибка соединения к %s: %v", serverURL, err)
		return
	}
	log.Printf("Связь с сервером %s установлена\n", serverURL)

	// GET запрос
	log.Printf("\nЗапрос данных (GET) к %s\n", serverURL)
	getResponse, err := GetServerInfo(client, serverURL)
	if err != nil {
		log.Printf("GET ошибка: %v", err)
		return
	}
	fmt.Printf("\nРезультат GET запроса: \n%s\n", getResponse)

	// POST запрос
	log.Printf("\nОтправка формы (POST) к %s\n", serverURL)
	form := url.Values{}
	form.Set("nickname", "Student")
	form.Set("feedback", "Всё отлично!")

	postResponse, err := PostServerData(client, serverURL, form)
	if err != nil {
		log.Printf("POST ошибка: %v", err)
		return
	}
	fmt.Printf("\nРезультат POST запроса: \n%s\n", postResponse)

	log.Println("Цикл работы клиента завершен успешно")
}
