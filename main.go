package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	pgx2 "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/file"
	myproject_sqlc "projectlibrary/sqlcout"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	"projectlibrary/controller"
)

// RunMigrate executes database schema migrations.
func runMigrate(connString string) {
	// Path to the SQL migration files.
	sourceURL := "file://./sql/migrations"

	// Migration driver.
	p := &pgx2.Postgres{}
	dbDriver, err := p.Open(connString)
	if err != nil {
		log.Fatalf("Could not connect to db files: %s", err)
	}

	// Open migration files from the source URL.
	s, err := (&file.File{}).Open(sourceURL)
	if err != nil {
		log.Fatalf("Could not open migration files: %s", err)
	}

	// Create a new migrate instance.
	m, err := migrate.NewWithInstance(sourceURL, s, "postgres", dbDriver)
	if err != nil {
		log.Fatalf("Could not create migration instance: %s", err)
	}

	// Apply all up migrations. ErrNoChange is ignored as it means the database is already up-to-date.
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %s", err)
	}
}

func main() {
	var err error
	// Read database connection configuration from environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Format the DSN string for the pgx driver.
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	// Format the connection string for the golang-migrate library.
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass,
		dbHost, dbPort, dbName)

	// Initialize a background context.
	ctx := context.Background()

	// Establish a connection to PostgreSQL using pgx.
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: `%v`", dsn)
	}
	defer conn.Close(ctx) // Defer closing the connection to ensure it's closed on application exit.

	// Run database migrations.
	runMigrate(connString)

	// Initialize the SQLC query generator with the active connection.
	queries := myproject_sqlc.New(conn)

	// Create the HTTP request controller and inject the SQLC queries.
	b := controller.NewController(queries)

	// Set Gin to debug mode for detailed logging.
	gin.SetMode(gin.DebugMode)

	// Initialize the default Gin router (with logger and recovery middleware).
	r := gin.Default()

	// Load HTML templates for rendering frontend pages.
	r.LoadHTMLGlob("templates/*.html")

	// Application routes.
	r.GET("/", b.GetBooks)                      // Main page: list of books.
	r.GET("/add_books", b.PageGreateBooks)      // Page with the form to add a new book.
	r.POST("/add_books", b.AddBook)             // Handle new book form submission.
	r.GET("/delete_books/:id", b.DeleteBooks)   // Delete a book.
	r.GET("/update_books/:id", b.PageUpdateForm) // Page with the form to edit book data.
	r.POST("/update_books/:id", b.UpdateBooks)  // Handle book update form submission.
	r.GET("/read_book/:id", b.ReadBookPage)     // Book reading page.
	r.GET("/genre/:genre", b.GetBooksByGenre)   // Filter book list by genre.
	r.POST("/add_text/:id", b.AddBookContext)   // Add/update book content.

	// Start the web server on port 8080.
	r.Run(":8080")
}
