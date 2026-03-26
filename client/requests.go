package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func pingServer(client *http.Client, serverURL string) {
	log.Printf("Проверка соединения с %s\n", serverURL)
	res, err := client.Head(serverURL)
	if err != nil {
		log.Printf("ошибка соединения: сервер недоступен: %v", err)
		return
	}
	defer res.Body.Close()
	log.Printf("Связь с сервером установлена\n")
}

func getRequest(client *http.Client, serverURL string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, serverURL, http.NoBody)
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
	log.Printf("%s-запрос. Запрос данных к %s\n", res.Request.Method, serverURL)
	return res
}

func postRequest(client *http.Client, serverURL string) *http.Response {
	// Собираем данные формы в удобную структуру
	form := url.Values{}
	form.Set("nickname", "Student")
	form.Set("feedback", "Всё отлично!")

	res, err := client.PostForm(serverURL, form) // Отправляем POST-запрос с данными формы
	if err != nil {
		log.Printf("Ошибка выполнения %s-запроса: %v", http.MethodPost, err)
		return nil
	}
	log.Printf("%s-запрос. Отправка формы к %s\n", res.Request.Method, serverURL)
	return res
}

func printResponse(res *http.Response) {
	// Проверка на пустой ответ от сервера
	if res == nil {
		return
	}

	// Безопасно планируем закрытие тела ответа. Это главное правило для предотвращения утечек!
	defer res.Body.Close()

	// Формируем и выводим на экран статус и заголовок ответа сервера
	fmt.Printf("[Статус и Заголовок ответа от сервера на %s-запрос:]", res.Request.Method)
	dump, _ := httputil.DumpResponse(res, false)
	fmt.Printf("\n%s", dump)

	// Читаем все данные, которые прислал сервер
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Ошибка чтения тела %s-запроса: %v", res.Request.Method, err)
		return
	}
	// Выводим все данные, которые прислал сервер на экран
	fmt.Printf("[Вывод на экран ответ сервера на %s-запрос:] %s", res.Request.Method, string(body))
}

func executeWork(client *http.Client, serverURL string) {
	// Проверка соединения с сервером
	pingServer(client, serverURL)

	getResponse := getRequest(client, serverURL) // формирование и отправка GET-запроса на сервер
	printResponse(getResponse)                   // получение и вывод на экран ответ сервера на GET-запрос

	postResponse := postRequest(client, serverURL) // формирование и отправка POST-запроса на сервер
	printResponse(postResponse)                    // получение и вывод на экран ответ сервера на POST-запрос

	log.Println("Цикл работы клиента завершен успешно\n")
}
