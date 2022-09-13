package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Book struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	Count     int       `json:"count"`
	IsNew     bool      `json:"is_new"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type IBookRepository interface {
	GetBooks() []*Book
	GetBook(id int) []*Book
	PostBook(book Book)
	DeleteBook(id int)
	UpdateBook(book Book)
}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() IBookRepository {
	db, err := sql.Open("pgx", "postgres://postgres:456123789@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &BookRepository{db: db}
}

func (b *BookRepository) GetBooks() []*Book {
	var books []*Book
	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Count, &book.IsNew, &book.Price, &book.CreatedAt)
		if err != nil {
			panic(err)
		}
		books = append(books, &book)
	}
	return books
}

func (b *BookRepository) GetBook(id int) []*Book {
	var book []*Book

	rows, err := b.db.Query("SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var bookscan Book
		err = rows.Scan(&bookscan.Id, &bookscan.Name, &bookscan.Author, &bookscan.Publisher, &bookscan.Count, &bookscan.IsNew, &bookscan.Price, &bookscan.CreatedAt)
		if err != nil {
			panic(err)
		}
		book = append(book, &bookscan)
	}
	return book
}

func (b *BookRepository) DeleteBook(id int) {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func (b *BookRepository) PostBook(bk Book) {
	var id int
	if err := b.db.QueryRow("INSERT INTO books (name, author, publisher, count, is_new, price, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id", bk.Name, bk.Author, bk.Publisher, bk.Count, bk.IsNew, bk.Price, bk.CreatedAt).Scan(&id); err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func (b *BookRepository) UpdateBook(bk Book) {
	_, err := b.db.Exec("UPDATE books SET name = $1, author = $2, publisher = $3, count = $4, is_new = $5, price = $6, created_at = $7 WHERE id = $8", bk.Name, bk.Author, bk.Publisher, bk.Count, bk.IsNew, bk.Price, bk.CreatedAt, bk.Id)
	if err != nil {
		panic(err)
	}
}
