package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	myproject_sqlc "projectlibrary/sqlcout"
	"strconv"
	"strings"
)

// genres представляет собой предопределенный список жанров для использования в приложении.
var genres = []string{
	"Фантастика", "Детектив", "Роман", "Приключения", "Научпоп", "Детская литература", "Поэзия", "Классика",
}

type Book struct {
	ID    int64  `json:"id"     form:"id"`
	Title string `json:"title"  form:"title"`
	Author string `json:"author"  form:"author"`
	Year  int    `json:"year"   form:"year"`
	Genre string `json:"genre"  form:"genre"`
	Content string `json:"content"  form:"content"`
}

type BookHandler struct {
	DB *myproject_sqlc.Queries
}

// NewController создает новый экземпляр BookHandler с внедренными зависимостями SQLC.
func NewController(db *myproject_sqlc.Queries) BookHandler {
		return BookHandler{DB: db}
}

// GetBooks отображает главную страницу со списком всех доступных книг.
func (b BookHandler) GetBooks(ctx *gin.Context) {
	books, err := b.GetBookSlice(ctx)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Библиотека книг", "books": books})
}

// GetBookSlice извлекает полный список книг напрямую из базы данных.
func (b BookHandler) GetBookSlice(ctx *gin.Context) ([]myproject_sqlc.Books, error) {

	rows, err := b.DB.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	return rows, err
}

// PageGreateBooks отображает форму добавления новой книги.
func (b BookHandler) PageGreateBooks(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add.html", gin.H{
		"TitleForm": "Добавление книги",
		"Genres":    genres,
	})
}

// AddBook обрабатывает отправку формы и создает новую запись о книге в базе данных.
func (b BookHandler) AddBook(ctx *gin.Context) {
	var book Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.HTML(http.StatusBadRequest, "add.html", gin.H{
			"TitleForm": "Добавление книги",
			"error":     "Ошибка при добавлении книги: " + err.Error(),
			"book":      book,
		})
		return
	}

		// Параметры для sqlc - запроса (поля в БД имеют ограничение NOT NULL)
	params := myproject_sqlc.InsertBookParams{

		Title:  book.Title,
		Author: book.Author,
		Year:   int32(book.Year),
		Genre:  book.Genre,
		Content: book.Content,
	}
		// Вставка в БД
		_, err := b.DB.InsertBook(ctx, params)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "add.html", gin.H{
				"TitleForm": "Добавление книги",
				"error": "Ошибка базы данных: " + err.Error(),
				"book":  book,
			})
			return
		}

		ctx.Redirect(http.StatusFound, "/")
}

// DeleteBooks удаляет книгу по её идентификатору, если редактирование разрешено.
func (b BookHandler) DeleteBooks(ctx *gin.Context){
	// Проверяем флаг из энвов
	if os.Getenv("ENABLE_BOOK_EDITING") != "true" {
		// Если в .env написано что-то кроме "true", отдаем ошибку 403 (Запрещено)
		ctx.HTML(http.StatusForbidden, "error.html", gin.H{
			"error": "Удаление книг временно отключено администратором",
		})
		return // Прерываем выполнение, в базу даже не смотрим!
	}
	bookID := ctx.Param("id")  //Когда мы получаем ID из параметров URL (ctx.Param("id")), он всегда возвращается как строка.

	// Преобразование bookID в int
	id, err := strconv.Atoi(bookID)//Используем strconv.Atoi, если уверены, что bookID всегда будет в пределах диапазона int
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Некорректный ID книги"})
		return
	}

	err = b.DB.DeleteBook(ctx, int32(id))  // Преобразование int в int32
		if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", nil )
		return
	}
	ctx.Redirect(http.StatusFound, "/")
	}

// PageUpdateForm отображает страницу редактирования существующей книги.
func (b BookHandler) PageUpdateForm(ctx *gin.Context) {
	if os.Getenv("ENABLE_BOOK_EDITING") != "true" {
		ctx.HTML(http.StatusForbidden, "error.html", gin.H{
			"error": "Редактирование книг отключено администратором",
		})
		return
	}
	bookID := ctx.Param("id")
	id, _ := strconv.Atoi(bookID)
	var book Book
	book, err := b.GetBookByID(ctx, int32(id))
		if err != nil{
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "update.html", gin.H{
		"BookID":    bookID,
		"Title":     book.Title,
		"Paragraphs": strings.Split(book.Content, "\n"),
		"Author":    book.Author,
		"Year":      book.Year,
		"Genre":     book.Genre,
		"Content":   book.Content,
		"TitleForm": "Редактирование книги",
	})

}

// GetBookByID запрашивает книгу по ID и преобразует результат в структуру Book.
func (b BookHandler) GetBookByID(ctx *gin.Context, id int32) (Book, error) {

	row, err := b.DB.GetBookAllData(ctx, id)
	if err != nil {
		return Book{}, err // Возвращаем пустую Book и ошибку
	}

	// Преобразуем результат из типа, сгенерированного SQLc, в тип Book
	result := Book{
		ID:    int64(row.ID),
		Title: row.Title,
		Author: row.Author,
		Year:  int(row.Year),
		Genre: row.Genre,
		Content: row.Content,

	}
	return result, nil
}

// UpdateBooks обрабатывает запрос на обновление данных книги.
func (b BookHandler) UpdateBooks(ctx *gin.Context){
	if os.Getenv("ENABLE_BOOK_EDITING") != "true" {
		ctx.HTML(http.StatusForbidden, "error.html", gin.H{
			"error": "Редактирование книг отключено администратором",
		})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Некорректный ID книги"})
		return
	}
	// Получаем  данные книги, если они есть
	_, err = b.GetBookByID(ctx, int32(id))
	if err != nil {
		// Обработка ошибки
		return
	}

	var updatedBook Book
	if err := ctx.ShouldBind(&updatedBook); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Ошибка привязки данных формы"})
		return
	}

	params := myproject_sqlc.UpdateBookParams{
		ID:      int32(id),
		Title:   updatedBook.Title,
		Author:   updatedBook.Author,
		Year:    int32(updatedBook.Year),
		Genre:   updatedBook.Genre,
		Content:  updatedBook.Content,
	}

	_, err = b.DB.UpdateBook(ctx, params)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"title": "Error", "error": "Ошибка при обновлении книги в базе данных"})
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/read_book/"+idStr)
}

// ReadBookPage - функция-обработчик для чтения книги
func (b BookHandler) ReadBookPage(ctx *gin.Context) {
	bookIDStr := ctx.Param("id")
	id, _ := strconv.Atoi(bookIDStr)

	book, err := b.GetBookByID(ctx, int32(id))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	paragraphs := strings.Split(book.Content, "\n")
	allowEditing := os.Getenv("ENABLE_BOOK_EDITING") == "true"

	ctx.HTML(http.StatusOK, "reading.html", gin.H{
		"Title": book.Title,
		"Paragraphs": paragraphs,
		"BookID": bookIDStr,
	"AllowEditing": allowEditing,
	})
}

// GetBooksByGenre выводит список книг, отфильтрованных по определенному жанру.
func (b BookHandler) GetBooksByGenre(ctx *gin.Context) {
	bookGenre := ctx.Param("genre")
	books, err := b.GetListBooksByGenre(ctx, bookGenre)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Библиотека книг", "books": books})
}

// GetListBooksByGenre запрашивает книги по жанру и приводит их к общему типу []Book.
func (b BookHandler) GetListBooksByGenre(ctx *gin.Context, bookGenre string) ([]Book, error) {

	rows, err := b.DB.ListBooksGenre(ctx, bookGenre)
		if err != nil {
		return nil, err
	}

	// Преобразование []myproject_sqlc.Book в []Book
	books := make([]Book, len(rows))
	for i, row := range rows {
		books[i] = Book{
			ID:    int64(row.ID),
			Title: row.Title,
			Author: row.Author,
			Year:  int(row.Year),
			Genre: row.Genre,
		}
	}

	return books, nil
}

// GetBookContentByID возвращает текстовое содержимое книги по её идентификатору.
func (b BookHandler) GetBookContentByID(ctx *gin.Context, bookID int32 ) (string, error) {

	book, err := b.GetBookByID(ctx, bookID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "reading.html", gin.H{"error": err.Error()})
		return "", err
	}
	return book.Content, nil

}

// AddBookContext дополняет существующий текст книги новыми абзацами из формы.
func (b BookHandler) AddBookContext(ctx *gin.Context) {

	if os.Getenv("ENABLE_BOOK_EDITING") != "true" {
		ctx.HTML(http.StatusForbidden, "read_book.html", gin.H{
			"error": "Редактирование содержимого книг временно отключено администратором",
		})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Обработка ошибки
		return
	}
	bookID := int32(id)

	// Получаем существующее содержимое книги.
	existingContent, err := b.GetBookContentByID(ctx, bookID)
	if err != nil {
		// Обработка ошибки
		return
	}

	// Получаем данные из формы.
	var updatedBookData struct {
		AddText string `form:"add_text"` // Поле для добавления нового текста
	}
	if err := ctx.ShouldBind(&updatedBookData); err != nil {
		ctx.HTML(http.StatusBadRequest, "reading.html", gin.H{"error": "Ошибка привязки данных формы"})
		return
	}
	// Проверяем, был ли введен новый текст.
	if updatedBookData.AddText != "" {
		// Объединяем старый и новый текст.
		newContent := existingContent + "\n\n" + updatedBookData.AddText

		// Создаем параметры для обновления.
		params := myproject_sqlc.UpdateBookContentParams{
			ID:      bookID,
			Content:  newContent,
		}
		_, err = b.DB.UpdateBookContent(ctx, params)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "reading.html", gin.H{"error": "Ошибка обновления содержимого книги"})
			return
		}
	}

	ctx.Redirect(http.StatusFound, "/read_book/"+idStr)
}
