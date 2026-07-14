# ProjectLibrary

A simple web application for managing a personal book library. Built with Go, Gin, and PostgreSQL.

## Features

*   View a list of all books.
*   Add new books with title, author, year, genre, and content.
*   Edit book details and content.
*   Delete books (this feature can be disabled via configuration).
*   Read book content in a user-friendly reader view.
*   Reader view includes font size adjustment and a dark/light theme switcher.
*   Filter books by genre.
*   Simple and clean UI using [Pico.css](https://picocss.com/).

## Tech Stack

*   **Backend:** Go, Gin
*   **Database:** PostgreSQL
*   **Database Interaction:** sqlc
*   **Migrations:** golang-migrate
*   **Frontend:** HTML Templates, Pico.css
*   **Containerization:** Docker

## Prerequisites

*   Go (version 1.18 or higher)
*   Docker
*   `make` (optional, for using Makefile commands)
*   `migrate` CLI (optional, for manual migrations)

## Getting Started

### 1. Configure Environment

Create a `.env` file in the root directory. This file stores the database connection details and application settings.

```env
# PostgreSQL Connection
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secret_password
DB_NAME=library_db

# Feature Flag
# Set to "true" to allow editing, deleting, and adding content to books.
ENABLE_BOOK_EDITING=true
```

### 2. Running the Application

#### Option A: Using Docker (Recommended)

This is the easiest way to get the application and a compatible database running.

1.  **Build the Docker image:**
    ```sh
    docker build -t project-library .
    ```
2.  **Run the application container** (linking it to your running PostgreSQL container):
    ```sh
    docker run -p 8080:8080 --env-file .env --name library-app project-library
    ```
    *Note: This assumes you have a PostgreSQL container running and accessible.*

#### Option B: Running Locally

1.  **Ensure you have a running PostgreSQL instance** that matches the configuration in your `.env` file.

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Run database migrations:**
    The application runs migrations automatically on startup. You can also run them manually using the `migrate` CLI.

4.  **Run the application:**
    ```sh
    go run main.go
    ```

The application will be available at [http://localhost:8080](http://localhost:8080).

## Project Structure

```
/controller   - Gin handlers for HTTP requests.
/sql          - SQL files for migrations and sqlc queries.
/sqlcout      - Generated Go code from sqlc.
/templates    - HTML templates for the frontend.
main.go       - Application entry point, router setup.
Dockerfile    - Docker configuration for the application.
Makefile      - Helper commands for building, running, and migrating.
README.md     - This file.
```

---

# ProjectLibrary

Простое веб-приложение для управления личной библиотекой книг. Создано с использованием Go, Gin и PostgreSQL.

## Возможности

*   Просмотр списка всех книг.
*   Добавление новых книг с указанием названия, автора, года, жанра и содержания.
*   Редактирование информации о книге и ее содержания.
*   Удаление книг (эту функцию можно отключить в конфигурации).
*   Чтение содержимого книги в удобном интерфейсе.
*   Интерфейс для чтения включает настройку размера шрифта и переключатель темной/светлой темы.
*   Фильтрация книг по жанру.
*   Простой и чистый UI с использованием [Pico.css](https://picocss.com/).

## Технологический стек

*   **Бэкенд:** Go, Gin
*   **База данных:** PostgreSQL
*   **Взаимодействие с БД:** sqlc
*   **Миграции:** golang-migrate
*   **Фронтенд:** HTML-шаблоны, Pico.css
*   **Контейнеризация:** Docker

## Необходимые компоненты

*   Go (версия 1.18 или выше)
*   Docker
*   `make` (необязательно, для использования команд из Makefile)
*   `migrate` CLI (необязательно, для ручного выполнения миграций)

## Начало работы

### 1. Настройка окружения

Создайте файл `.env` в корневой директории проекта. Этот файл хранит данные для подключения к базе данных и настройки приложения.

```env
# Подключение к PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secret_password
DB_NAME=library_db

# Флаг функциональности
# Установите "true", чтобы разрешить редактирование, удаление и добавление контента в книги.
ENABLE_BOOK_EDITING=true
```

### 2. Запуск приложения

#### Вариант А: Использование Docker (Рекомендуется)

Это самый простой способ запустить приложение и совместимую с ним базу данных.

1.  **Соберите Docker-образ:**
    ```sh
    docker build -t project-library .
    ```
2.  **Запустите контейнер приложения** (связав его с вашим запущенным контейнером PostgreSQL):
    ```sh
    docker run -p 8080:8080 --env-file .env --name library-app project-library
    ```
    *Примечание: Предполагается, что у вас запущен и доступен контейнер PostgreSQL.*

#### Вариант Б: Локальный запуск

1.  **Убедитесь, что у вас запущен экземпляр PostgreSQL**, соответствующий настройкам в вашем файле `.env`.

2.  **Установите зависимости:**
    ```sh
    go mod tidy
    ```

3.  **Выполните миграции базы данных:**
    Приложение выполняет миграции автоматически при запуске. Вы также можете выполнить их вручную с помощью `migrate` CLI.

4.  **Запустите приложение:**
    ```sh
    go run main.go
    ```

Приложение будет доступно по адресу [http://localhost:8080](http://localhost:8080).

## Структура проекта

```
/controller   - Обработчики Gin для HTTP-запросов.
/sql          - SQL-файлы для миграций и запросов sqlc.
/sqlcout      - Сгенерированный Go-код от sqlc.
/templates    - HTML-шаблоны для фронтенда.
main.go       - Точка входа в приложение, настройка роутера.
Dockerfile    - Конфигурация Docker для приложения.
Makefile      - Вспомогательные команды для сборки, запуска и миграций.
README.md     - Этот файл.
```
