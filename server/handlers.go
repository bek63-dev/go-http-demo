package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch {
	case req.URL.Path == "/time", req.URL.Path == "/time/":
		currentTime := time.Now().Format("02.01.2006 15:04:05")
		fmt.Fprintf(res, "Текущее время: %s", currentTime)

	case req.Method == http.MethodGet:
		fmt.Fprintf(res, "--- HTTP Заголовки ---\n")
		for k, values := range req.Header {
			for _, v := range values {
				fmt.Fprintf(res, "%s: %s\n", k, v)
			}
		}

		fmt.Fprintf(res, "--- GET Параметры ---\n")
		query := req.URL.Query()
		if len(query) == 0 {
			fmt.Fprintf(res, "Параметры отсутствуют\n")
		} else {
			for k, v := range query {
				fmt.Fprintf(res, "%s: %v\n", k, v)
			}
		}
	case req.Method == http.MethodPost:
		if err := req.ParseForm(); err != nil {
			http.Error(res, "Ошибка парсинга формы", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(res, "Host: %s\nPath: %s\nMethod: %s\n", req.Host, req.URL.Path, req.Method)
		fmt.Fprintf(res, "\n--- Данные формы (POST) ---\n")
		for k, values := range req.Form {
			for _, v := range values {
				fmt.Fprintf(res, "%s: %s\n", k, v)
			}
		}
	default:
		http.Error(res, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
