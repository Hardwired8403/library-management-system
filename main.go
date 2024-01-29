package main

import (
	"fmt"
	"library-management-system/database"
	"library-management-system/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Ошибка при инициализации базы данных:", err)
		return
	}
	defer db.Close()

	// Создание маршрутизатора
	r := mux.NewRouter()

	// Регистрация обработчиков
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Настройка обработки корневого URL
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать в систему управления библиотекой!")
	})

	// Обработка запросов через маршрутизатор
	http.Handle("/", r)

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}
