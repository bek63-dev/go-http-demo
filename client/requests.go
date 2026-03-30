package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func pingServer(client *http.Client, baseURL string) {
	log.Printf("Проверка соединения с %s\n", baseURL)
	res, err := client.Head(baseURL)
	if err != nil {
		log.Printf("ошибка соединения: сервер недоступен: %v", err)
		return
	}
	defer res.Body.Close()
	log.Printf("Связь с сервером установлена\n")
}

func getRequest(client *http.Client, baseURL string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, baseURL, http.NoBody)
	if err != nil {
		log.Printf("Не удалось создать %s-запрос: %v", http.MethodGet, err)
		return nil
	}

	// Добавляем дополнительные кастомные заголовки
	req.Header.Set("Custom-Header", "John Doe")
	req.Header.Set("Accept-Language", "en")

	res, err := client.Do(req) // Отправляем запрос через клиент
	if err != nil {
		log.Printf("Ошибка выполнения %s-запроса: %v", http.MethodGet, err)
		return nil
	}
	log.Printf("%s: Запрос данных к %s\n", res.Request.Method, baseURL)
	return res
}

func postRequest(client *http.Client, baseURL string) *http.Response {
	// Собираем данные формы в удобную структуру
	form := url.Values{}
	form.Set("email", "my@my-best-site.ru")
	form.Set("name", "Василий")

	res, err := client.PostForm(baseURL, form) // Отправляем POST-запрос с данными формы
	if err != nil {
		log.Printf("Ошибка выполнения %s-запроса: %v", http.MethodPost, err)
		return nil
	}
	log.Printf("%s: Отправка формы к %s\n", res.Request.Method, baseURL)
	return res
}

func printResponse(res *http.Response, baseURL string) {
	// Проверка на пустой ответ от сервера
	if res == nil {
		return
	}

	// Безопасно планируем закрытие тела ответа. Это главное правило для предотвращения утечек!
	defer res.Body.Close()

	// Формируем и выводим на экран статус и заголовок ответа сервера
	fmt.Printf("[Статус и Заголовок ответа сервера (%s: %s]", res.Request.Method, baseURL)
	dump, _ := httputil.DumpResponse(res, false)
	fmt.Printf("\n%s", dump)

	// Читаем все данные, которые прислал сервер
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Ошибка чтения тела (%s %s: %v", res.Request.Method, baseURL, err)
		return
	}
	// Выводим все данные, которые прислал сервер на экран
	fmt.Printf("[Вывод на экран ответ сервера (%s: %s] %s", res.Request.Method, baseURL, string(body))
}

func executeWork(client *http.Client, baseURL string) {
	// Проверка соединения с сервером
	pingServer(client, baseURL)

	getResponse := getRequest(client, baseURL) // формирование и отправка GET-запроса на сервер по конкретному URL
	printResponse(getResponse, baseURL)        // получение и вывод на экран ответ сервера на GET-запрос по конкретному URL

	postResponse := postRequest(client, baseURL) // формирование и отправка POST-запроса на сервер по конкретному URL
	printResponse(postResponse, baseURL)         // получение и вывод на экран ответ сервера на POST-запрос по конкретному URL

	log.Println("Цикл работы клиента завершен успешно\n")
}
