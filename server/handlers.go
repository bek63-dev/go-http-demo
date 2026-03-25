package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	var out string
	var headers string
	var form string

	if req.URL.Path == "/time" || req.URL.Path == "/time/" {
		out = time.Now().Format("02.01.2006 15:04:05")
	}

	switch req.Method {
	case http.MethodGet:
		for k, values := range req.Header {
			for _, v := range values {
				headers += fmt.Sprintf("%s: %s\n", k, v)
			}
		}
		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s\nHeader: %s\n",
			req.Host, req.URL.Path, req.Method, headers)
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			http.Error(res, "Ошибка парсинга формы", http.StatusBadRequest)
			return
		}
		for k, values := range req.Form {
			for _, v := range values {
				form += fmt.Sprintf("%s: %s\n", k, v)
			}
		}
		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s\nForm Data: %s\n",
			req.Host, req.URL.Path, req.Method, form)
	default:
		http.Error(res, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	res.Write([]byte(out))
}
