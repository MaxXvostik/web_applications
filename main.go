package main

import (
	"html/template"
	"net/http"
	"os"
)

// Подключаем HTML файл
var tpl = template.Must(template.ParseFiles("index.html"))

//Функцию-обработчик для корневого пути /
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Для подключения css стилей
	fs := http.FileServer(http.Dir("assets"))

	// Создаем новый мультиплексор HTTP-запросов(Сопоставляет URL-адрес входящих запросов)
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// Этот тип реализует интерфейс Handler путем реализации функции ServeHTTP

	mux.HandleFunc("/", indexHandler)
	//Запускает сервер на порту 3000
	http.ListenAndServe(":"+port, mux)
}
