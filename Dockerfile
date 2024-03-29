# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем переменную окружения для работы вне модуля Go
ENV GO111MODULE=on

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы проекта в текущую директорию
COPY . .

# Этап для сборки и запуска тестов
RUN go test ./...

# Этап для сборки приложения
RUN go build -o main .

# Экспортируем порт 8080
EXPOSE 8080

# Команда для запуска приложения при старте контейнера
CMD ["./main"]
