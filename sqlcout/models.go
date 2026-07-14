package myproject_sqlc

type Books struct {
	ID      int32  `db:"id"`
	Title   string `db:"title"`
	Author   string `db:"author"`
	Year    int32  `db:"year"`
	Genre   string `db:"genre"`
	Content string `db:"content"`
}
