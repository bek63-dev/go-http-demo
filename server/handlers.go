package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// dumpHTTPRequest сериализует HTTP-запрос в "сырой" текстовый формат (Wire Format).
// Она возвращает срез байтов, содержащий стартовую строку, заголовки и тело запроса.
// Если в процессе дампа происходит ошибка, функция логирует её и возвращает пустой срез.
func dumpHTTPRequest(res http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println("Ошибка при дампе:", err)
		http.Error(res, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Fprintf(res, "\n%s", dump)
}

func dumpHTTPParams(res http.ResponseWriter, req *http.Request, values url.Values) {
	fmt.Fprintf(res, "%s-параметры (Body):\n", req.Method)

	if len(values) == 0 {
		fmt.Fprintf(res, "Параметры отсутствуют\n\n")
		return
	}

	for k, v := range values {
		fmt.Fprintf(res, "%s: %v\n", k, v)
	}
}

func handleHomeGet(res http.ResponseWriter, req *http.Request) {
	dumpHTTPRequest(res, req)
	dumpHTTPParams(res, req, req.URL.Query())
}

func handleHomePost(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}
	dumpHTTPRequest(res, req)
	dumpHTTPParams(res, req, req.PostForm)
}

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch req.Method {
	case http.MethodGet:
		handleHomeGet(res, req)
	case http.MethodPost:
		handleHomePost(res, req)
	default:
		http.Error(res, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
