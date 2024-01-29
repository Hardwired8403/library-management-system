# Library Management System

Проект "Library Management System" представляет собой систему управления библиотекой, написанную на языке программирования Go. Система предоставляет API для работы с книгами, их создания, чтения, обновления и удаления.

## Установка и Запуск

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/yourusername/library-management-system.git
   ```
2. Перейдите в каталог проект:
    ```bash
    cd library-management-system
   ```
3. Установите зависимости:
    ```bash
    go get -u ./...
   ```
4. Запустите проект:
    ```bash
    go run main.go
   ```
Сервер будет запущен на порту 8080.

## Использование API

### Получение списка книг
```bash
curl http://localhost:8080/books 
```

### Получение информации о книге по ID
```bash
curl http://localhost:8080/books/{id}
```

### Создание новой книги
```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Новая книга", "author": "Автор", "quantity": 10}' http://localhost:8080/books
```

### Обновление информации о книге по ID
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Обновленная книга", "author": "Обновленный автор", "quantity": 15}' http://localhost:8080/books/{id}
```

### Удаление книги по ID
```bash
curl -X DELETE http://localhost:8080/books/{id}
```

## Тестирование
Для запуска тестов используйте команду:
```bash
go test
```

## Зависимости
1. [Gorilla Mux](https://github.com/gorilla/mux) - маршрутизатор HTTP для Go.
2. [Mattn Go-SQLite3](https://github.com/mattn/go-sqlite3) - SQLite3 драйвер для базы данных.

## Автор
Карен

## Лицензия
Этот проект лицензирован в соответствии с лицензией MIT - подробности смотрите в файле LICENSE.