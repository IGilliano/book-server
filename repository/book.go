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
	GetBooks() ([]*Book, error)
	GetBook(id int) ([]*Book, error)
	PostBook(book Book) error
	DeleteBook(id int) error
	UpdateBook(book Book) error
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

func (b *BookRepository) GetBooks() ([]*Book, error) {
	var books []*Book
	rows, err := b.db.Query("SELECT * FROM books")

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Count, &book.IsNew, &book.Price, &book.CreatedAt)
		if err == nil {
			books = append(books, &book)
		}
	}
	return books, err
}

func (b *BookRepository) GetBook(id int) ([]*Book, error) {
	var book []*Book

	rows, err := b.db.Query("SELECT * FROM books WHERE id=$1", id)

	for rows.Next() {
		var bookscan Book
		err = rows.Scan(&bookscan.Id, &bookscan.Name, &bookscan.Author, &bookscan.Publisher, &bookscan.Count, &bookscan.IsNew, &bookscan.Price, &bookscan.CreatedAt)
		if err == nil {
			book = append(book, &bookscan)
		}
	}
	return book, err
}

func (b *BookRepository) DeleteBook(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}

func (b *BookRepository) PostBook(bk Book) error {
	var id int
	err := b.db.QueryRow("INSERT INTO books (name, author, publisher, count, is_new, price, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id", bk.Name, bk.Author, bk.Publisher, bk.Count, bk.IsNew, bk.Price, bk.CreatedAt).Scan(&id)
	return err
}

func (b *BookRepository) UpdateBook(bk Book) error {
	fmt.Println(bk)
	_, err := b.db.Exec("UPDATE books SET name = $1, author = $2, publisher = $3, count = $4, is_new = $5, price = $6, created_at = $7 WHERE id = $8", bk.Name, bk.Author, bk.Publisher, bk.Count, bk.IsNew, bk.Price, bk.CreatedAt, bk.Id)
	return err
}
