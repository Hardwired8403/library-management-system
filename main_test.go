package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"library-management-system/database"
	"library-management-system/handlers"
	"library-management-system/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	// Инициализация тестовой базы данных
	db, err := database.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	// Создание тестового маршрутизатора
	r := setupTestRouter(db)

	// Создание тестового запроса с данными о книге
	newBook := models.Book{
		Title:    "Тестовая книга",
		Author:   "Тестовый автор",
		Quantity: 5,
	}
	newBookJSON, _ := json.Marshal(newBook)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
	assert.NoError(t, err)

	// Запуск теста
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	// Проверка кода ответа
	assert.Equal(t, http.StatusOK, resp.Code)

	// Проверка того, что создана новая книга
	var createdBook models.Book
	err = json.NewDecoder(resp.Body).Decode(&createdBook)
	assert.NoError(t, err)
	assert.NotZero(t, createdBook.ID)
	assert.Equal(t, newBook.Title, createdBook.Title)
	assert.Equal(t, newBook.Author, createdBook.Author)
	assert.Equal(t, newBook.Quantity, createdBook.Quantity)
}

func TestGetBooks(t *testing.T) {
	// Инициализация тестовой базы данных
	db, err := database.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	// Создание тестового маршрутизатора
	r := setupTestRouter(db)

	// Создание тестового запроса на получение списка книг
	req, err := http.NewRequest("GET", "/books", nil)
	assert.NoError(t, err)

	// Запуск теста
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	// Проверка кода ответа
	assert.Equal(t, http.StatusOK, resp.Code)

	// Проверка формата ответа (ожидаемый формат - JSON)
	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))

	// Проверка наличия данных о книгах в ответе
	var books []models.Book
	err = json.NewDecoder(resp.Body).Decode(&books)
	assert.NoError(t, err)
	assert.NotEmpty(t, books)
}

// Тесты для GetBook, UpdateBook, и DeleteBook могут быть написаны аналогичным образом.

func TestGetBook(t *testing.T) {
	// Инициализация тестовой базы данных
	db, err := database.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	// Создание тестового маршрутизатора
	r := setupTestRouter(db)

	// Создание тестовой книги для последующего получения информации
	newBook := models.Book{
		Title:    "Тестовая книга",
		Author:   "Тестовый автор",
		Quantity: 5,
	}
	newBookJSON, _ := json.Marshal(newBook)

	// Создание тестового запроса на создание новой книги
	createReq, err := http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
	assert.NoError(t, err)

	// Запуск теста создания книги
	createResp := httptest.NewRecorder()
	r.ServeHTTP(createResp, createReq)

	// Проверка кода ответа создания
	assert.Equal(t, http.StatusOK, createResp.Code)

	// Парсинг ответа создания книги для получения ID
	var createdBook models.Book
	err = json.NewDecoder(createResp.Body).Decode(&createdBook)
	assert.NoError(t, err)

	// Создание тестового запроса на получение информации о книге по ID
	getReq, err := http.NewRequest("GET", "/books/"+fmt.Sprint(createdBook.ID), nil)
	assert.NoError(t, err)

	// Запуск теста получения информации о книге
	getResp := httptest.NewRecorder()
	r.ServeHTTP(getResp, getReq)

	// Проверка кода ответа получения
	assert.Equal(t, http.StatusOK, getResp.Code)

	// Проверка данных в ответе
	var retrievedBook models.Book
	err = json.NewDecoder(getResp.Body).Decode(&retrievedBook)
	assert.NoError(t, err)

	// Проверка совпадения данных с созданной книгой
	assert.Equal(t, createdBook.ID, retrievedBook.ID)
	assert.Equal(t, createdBook.Title, retrievedBook.Title)
	assert.Equal(t, createdBook.Author, retrievedBook.Author)
	assert.Equal(t, createdBook.Quantity, retrievedBook.Quantity)
}

func TestDeleteBook(t *testing.T) {
	// Инициализация тестовой базы данных
	db, err := database.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	// Создание тестового маршрутизатора
	r := setupTestRouter(db)

	// Создание тестовой книги для последующего удаления
	newBook := models.Book{
		Title:    "Тестовая книга",
		Author:   "Тестовый автор",
		Quantity: 5,
	}
	newBookJSON, _ := json.Marshal(newBook)

	// Создание тестового запроса на создание новой книги
	createReq, err := http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
	assert.NoError(t, err)

	// Запуск теста создания книги
	createResp := httptest.NewRecorder()
	r.ServeHTTP(createResp, createReq)

	// Проверка кода ответа создания
	assert.Equal(t, http.StatusOK, createResp.Code)

	// Парсинг ответа создания книги для получения ID
	var createdBook models.Book
	err = json.NewDecoder(createResp.Body).Decode(&createdBook)
	assert.NoError(t, err)

	// Создание тестового запроса на удаление книги
	deleteReq, err := http.NewRequest("DELETE", "/books/"+fmt.Sprint(createdBook.ID), nil)
	assert.NoError(t, err)

	// Запуск теста удаления книги
	deleteResp := httptest.NewRecorder()
	r.ServeHTTP(deleteResp, deleteReq)

	// Проверка кода ответа удаления
	assert.Equal(t, http.StatusOK, deleteResp.Code)

	// Попытка получения информации о удаленной книге
	getReq, err := http.NewRequest("GET", "/books/"+fmt.Sprint(createdBook.ID), nil)
	assert.NoError(t, err)

	// Запуск теста получения информации о удаленной книге
	getResp := httptest.NewRecorder()
	r.ServeHTTP(getResp, getReq)

	// Проверка кода ответа получения (ожидаем 404, так как книга удалена)
	assert.Equal(t, http.StatusNotFound, getResp.Code)
}

// setupTestRouter создает тестовый маршрутизатор с переданной базой данных
func setupTestRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Регистрация обработчиков с переданной базой данных
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Настройка обработки корневого URL
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать в систему управления библиотекой!")
	})

	return r
}
