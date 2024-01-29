package models

// Reader представляет собой структуру данных для читателя в библиотеке.
type Reader struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewReader создает новый экземпляр структуры Reader с заданным именем.
func NewReader(name string) *Reader {
	return &Reader{
		Name: name,
	}
}
