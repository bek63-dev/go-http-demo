package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var (
	ErrRequestCreationFailed = errors.New("не удалось создать запрос: ")
	ErrRequestExecution      = errors.New("ошибка выполнения запроса: ")
	ErrResponseBodyRead      = errors.New("ошибка чтения тела: ")
	ErrServerUnavailable     = errors.New("сервер недоступен")
)

func PingServer(client *http.Client, serverURL string) error {
	resp, err := client.Head(serverURL)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrServerUnavailable, err)
	}
	defer resp.Body.Close()
	return nil
}

func GetServerInfo(client *http.Client, serverURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, serverURL, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrRequestCreationFailed, err)
	}

	req.Header.Set("Custom-Header", "John Doe")
	req.Header.Set("Accept-Language", "en")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrResponseBodyRead, err)
	}

	return string(body), nil
}

func PostServerData(client *http.Client, serverURL string, form url.Values) (string, error) {
	resp, err := client.PostForm(serverURL, form)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrRequestExecution, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrResponseBodyRead, err)
	}

	return string(body), nil
}

func ExecuteWork(client *http.Client, serverURL string) {
	// Проверка соединения с сервером
	log.Printf("Проверка соединения с %s\n", serverURL)

	if err := PingServer(client, serverURL); err != nil {
		log.Printf("Ошибка соединения: %v", err)
		return
	}
	log.Printf("Связь с сервером установлена\n")

	// GET запрос
	log.Printf("\nGET-запрос. Запрос данных к %s\n", serverURL)
	getResponse, err := GetServerInfo(client, serverURL)
	if err != nil {
		log.Printf("GET ошибка: %v", err)
		return
	}
	fmt.Printf("\nРезультат GET запроса: \n%s\n", getResponse)

	// POST запрос
	log.Printf("\nPOST-запрос. Отправка формы к %s\n", serverURL)
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
