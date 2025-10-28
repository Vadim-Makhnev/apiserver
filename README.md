# User Management REST API

RESTful API для управления пользователями, написанное на Go. Сервис предоставляет endpoints для регистрации, аутентификации и управления пользователями с использованием сессий для авторизации.

## 🚀 Функциональность

- **Регистрация пользователя** - Создание нового аккаунта с валидацией данных
- **Аутентификация** - Вход в систему с созданием сессии
- **Авторизация** - Защищённые endpoints, требующие активной сессии
- **Профиль пользователя** - Получение информации о текущем пользователе
- **Автоматическая документация** - Swagger UI для тестирования API

## 🛠 Технологии

- **Язык**: Go (Golang)
- **Маршрутизатор**: Gorilla Mux
- **База данных**: PostgreSQL
- **Миграции БД**: golang-migrate
- **Аутентификация**: Gorilla Sessions (cookie-based сессии)
- **Валидация**: ozzo-validation
- **Хеширование паролей**: bcrypt
- **Логирование**: Logrus
- **Тестирование**: Testify
- **Документация**: Swagger/OpenAPI
- **CORS**: Gorilla Handlers

## Swagger
<img width="2494" height="1122" alt="Screenshot 2025-10-28 at 13-24-31 Swagger UI" src="https://github.com/user-attachments/assets/f5bbcc4e-3399-4546-8159-afd270bbd473" />

## Быстрый старт
1. **Создание и настройка базы данных**
```sql
CREATE DATABASE IF NOT EXISTS restapi;
CREATE USER restapi WITH PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE restapi TO restapi;
```
```bash
migrate -path migrations -database "postgres://restapi:admin@localhost:5432/restapi?sslmode=disable" up
```
2. **Запуск приложения**
```go
go run ./cmd/apiserver
```
