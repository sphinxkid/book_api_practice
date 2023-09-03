CREATE DATABASE IF NOT EXISTS books;

USE books;
DROP TABLE IF EXISTS books;
CREATE TABLE books(
    book_id int NOT NULL AUTO_INCREMENT,
    book_name varchar(255),
    genre varchar(255),
    count int,
    PRIMARY KEY(book_id)
);
