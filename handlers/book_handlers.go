package handlers

import (
	"encoding/json"
	"library-management-system/database"
	"library-management-system/models"
	"net/http"

	"github.com/gorilla/mux"
)

// GetBooks обрабатывает запрос на получение списка всех книг.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Извлекаем список книг из базы данных
	books, err := database.GetAllBooks()
	if err != nil {
		http.Error(w, "Ошибка при получения списка книг", http.StatusInternalServerError)
		return
	}

	// Преобразуем список книг в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook обрабатывает запрос на получение информации о конкретной книге по ее ID.
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID книги из параметров запроса
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Извлекаем информацию о книге из базы данных
	book, err := database.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Ошибка при получения информации о книге", http.StatusInternalServerError)
		return
	}

	// Если книга не найдена, возвращаем ошибку 404
	if book == nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	// Преобразуем информацию о книге в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook обрабатывает запрос на создание новой книги.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Распарсим данные из тела запроса в структуру Book
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Вызовем функцию для создания новой книги
	createdBook, err := database.CreateNewBook(newBook)
	if err != nil {
		http.Error(w, "Ошибка при создании новой книги: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправим клиенту информацию о созданной книге
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBook)
}

// UpdateBook обрабатывает запрос на обновление информации о книге по ее ID.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID книги из параметров запроса
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Распарсим данные из тела запроса в структуру Book
	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Вызовем функцию для обновления информации о книге
	err = database.UpdateBookByID(bookID, updatedBook)
	if err != nil {
		http.Error(w, "Ошибка при обновлении информации о книге", http.StatusInternalServerError)
		return
	}

	// Отправим клиенту подтверждение об успешном обновлении
	w.WriteHeader(http.StatusOK)
}

// DeleteBook обрабатывает запрос на удаление книги по ее ID.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID книги из параметров запроса
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Вызовем функцию для удаления книги из базы данных
	err := database.DeleteBookByID(bookID)
	if err != nil {
		http.Error(w, "Ошибка при удалении книги", http.StatusInternalServerError)
		return
	}

	// Отправим клиенту подтверждение об успешном удалении
	w.WriteHeader(http.StatusOK)
}
