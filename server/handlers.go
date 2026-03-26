package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if req.URL.Path == "/time" || req.URL.Path == "/time/" {
		currentTime := time.Now().Format("02.01.2006 15:04:05")
		fmt.Fprintf(res, "Текущее время: %s", currentTime)
	}

	switch req.Method {
	case http.MethodGet:
		handleHomeGet(res, req)
	case http.MethodPost:
		handleHomePost(res, req)
	default:
		http.Error(res, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// dumpHTTPRequest сериализует HTTP-запрос в "сырой" текстовый формат (Wire Format).
// Она возвращает срез байтов, содержащий стартовую строку, заголовки и тело запроса.
// Если в процессе дампа происходит ошибка, функция логирует её и возвращает пустой срез.
func dumpHTTPRequest(req *http.Request) []byte {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println("Ошибка при дампе:", err)
		return []byte{}
	}
	return dump
}

func handleHomeGet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s", dumpHTTPRequest(req))

	fmt.Fprintf(res, "%s-параметры (Body):\n", req.Method)
	query := req.URL.Query()
	if len(query) == 0 {
		fmt.Fprintf(res, "параметры отсутствуют\n")
		return
	}
	for k, v := range query {
		fmt.Fprintf(res, "%s: %v\n", k, v)
	}
}

func handleHomePost(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, "%s", dumpHTTPRequest(req))

	fmt.Fprintf(res, "%s-параметры (Body):\n", req.Method)

	if len(req.PostForm) == 0 {
		fmt.Fprintln(res, "тело запроса пустое")
		return
	}

	for k, values := range req.PostForm {
		for _, v := range values {
			fmt.Fprintf(res, "%s: %s\n", k, v)
		}
	}
}
