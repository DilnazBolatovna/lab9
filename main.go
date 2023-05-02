package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Определяем директорию, в которой хранятся статические файлы HTML
	fs := http.FileServer(http.Dir("static"))

	// Определяем функцию-обработчик для корневого пути "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Отправляем файл index.html в качестве ответа
		http.ServeFile(w, r, "static/index.html")
	})

	// Определяем функцию-обработчик для пути "/submit"
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Разбираем значения формы из запроса
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Ошибка при разборе значений формы:", err)
			return
		}

		// Получаем значение поля "name" из формы
		name := r.Form.Get("name")

		// Отправляем значение name в качестве ответа
		fmt.Fprintf(w, "Привет %s dwd!", name)
	})

	// Запускаем сервер и слушаем входящие запросы
	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", fs)
}
