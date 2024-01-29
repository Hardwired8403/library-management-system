package database

import (
	"database/sql"
	"fmt"
	"library-management-system/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB инициализирует базу данных, создает необходимые таблицы и возвращает экземпляр DB.
func InitDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./library.db")
	if err != nil {
		return nil, err
	}

	// Создайте таблицы для книг и читателей
	_, err = db.Exec(`
  CREATE TABLE IF NOT EXISTS books (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title TEXT,
   author TEXT,
   quantity INTEGER
  );

  CREATE TABLE IF NOT EXISTS readers (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT
  );
 `)
	if err != nil {
		return nil, err
	}

	fmt.Println("База данных инициализирована успешно.")
	return db, nil
}

// GetAllBooks извлекает все книги из база данных.
func GetAllBooks() ([]models.Book, error) {
	rows, err := db.Query("SELECT id, title, author, quantity FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func GetBookByID(bookID string) (*models.Book, error) {
	var book models.Book
	err := db.QueryRow("SELECT id, title, author, quantity FROM books WHERE id = ?", bookID).
		Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		// Обрабатываем слуйчай, когда книга не найдена
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func CreateNewBook(newBook models.Book) (*models.Book, error) {
	// Выполним SQL-запрос для вставки новой книги в базу данных
	result, err := db.Exec("INSERT INTO books (title, author, quantity) VALUES (?, ?, ?)",
		newBook.Title, newBook.Author, newBook.Quantity)
	if err != nil {
		return nil, err
	}

	// Получим ID только что созданной книги
	newBookID, _ := result.LastInsertId()

	// Установим ID книги и вернем информацию о созданной книге
	newBook.ID = int(newBookID)
	return &newBook, nil
}

// UpdateBookByID обновляет информацию о книге в базе данных по ее ID.
func UpdateBookByID(bookID string, updatedBook models.Book) error {
	// Выполним SQL-запрос для обновления информации о книге
	_, err := db.Exec("UPDATE books SET title = ?, author = ?, quantity = ? WHERE id = ?",
		updatedBook.Title, updatedBook.Author, updatedBook.Quantity, bookID)
	return err
}

// DeleteBookByID удаляет книгу из базы данных по ее ID.
func DeleteBookByID(bookID string) error {
	// Выполним SQL-запрос для удаления книги из базы данных
	_, err := db.Exec("DELETE FROM books WHERE id = ?", bookID)
	return err
}
