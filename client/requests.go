package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrRequestCreationFailed = errors.New("не удалось создать запрос")
	ErrRequestExecution      = errors.New("ошибка выполнения запроса")
	ErrResponseBodyRead      = errors.New("ошибка чтения тела")
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
