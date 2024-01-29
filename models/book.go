package models

// Book представляет собой структуру данных для книги в библиотеке.
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// NewBook создает новый экземпляр структуры Book с заданными параметрами
func NewBook(title, author string, quantity int) *Book {
	return &Book{
		Title:    title,
		Author:   author,
		Quantity: quantity,
	}
}
