CREATE TABLE books (
                      id SERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      author VARCHAR(255) NOT NULL,
                      year int NOT NULL,
                      genre VARCHAR(255) NOT NULL,
content TEXT NOT NULL DEFAULT ''
);


INSERT INTO books (title, author, year, genre, content) VALUES ('Война и мир', 'Лев Толстой', 1869, 'Исторический роман', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('Джейн Эйр', 'Шарлотта Бронте', 1847, 'Роман', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('Гарри Поттер и философский камень', 'Дж. К. Роулинг', 1997, 'Фэнтези', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('Дюна', 'Фрэнк Герберт', 1965, 'Научная фантастика', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('1984', 'Джордж Оруэлл', 1949, 'Антиутопия', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('Мастер и Маргарита', 'Михаил Булгаков', 1967, 'Классика', '');
INSERT INTO books (title, author, year, genre, content) VALUES ('Великий Гэтсби', 'Ф. Скотт Фицджеральд', 1925, 'Классика', '');