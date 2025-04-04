ToDo API на Golang + Gin + GORM + PostgreSQL

Описание
Этот проект представляет собой REST API для управления задачами (ToDo) с использованием:
- Golang (Gin для веб-фреймворка, GORM для ORM)
- PostgreSQL (в качестве базы данных)
- Docker (для контейнеризации)

Запуск проекта

1. Клонирование репозитория

git clone https://github.com/DiasSatymbek05/go_lang.git
cd go_lang


2. Настройка переменных окружения
Создайте `.env` файл и укажите параметры базы данных:

DB_HOST=localhost
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydatabase
DB_PORT=5432


3. Запуск с помощью Docker

docker-compose up --build


4. Запуск без Docker
Убедитесь, что PostgreSQL запущен, затем:

go mod tidy
go run main.go


API Эндпоинты

--Получить все задачи--

GET http://localhost:8080/todos

--Получить задачу по ID--

GET http://localhost:8080/todos/{id}

--Создать новую задачу--

POST http://localhost:8080/todos
Content-Type: application/json
{
  "title": "Купить молоко",
  "completed": false
}

--Обновить задачу--

PUT http://localhost:8080/todos/{id}
Content-Type: application/json
{
  "title": "Купить хлеб",
  "completed": true
}

--Удалить задачу--

DELETE http://localhost:8080/todos/{id}

