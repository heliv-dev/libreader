-- name: ListBooks :many

SELECT *
FROM books;

-- name: InsertBook :one

INSERT INTO books (title, author, year, genre, content)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteBook :exec

DELETE
FROM books
WHERE id = $1;

-- name: UpdateBook :one

UPDATE books
SET title = $2,
    author = $3,
    year = $4,
           genre = $5,
           content = $6
WHERE id = $1 RETURNING *;

-- name: GetBookAllData :one

SELECT *
FROM books
WHERE id = $1; -- Retrieve all fields including the book content.

-- name: ListBooksGenre :many

SELECT *
FROM books
WHERE genre = $1;

-- name: UpdateBookContent :one

UPDATE books
SET content = $2
WHERE id = $1 RETURNING *;