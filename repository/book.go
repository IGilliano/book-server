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
	GetBook(id string) []*Book
	PostBook(book Book)
	DeleteBook(id string)
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

func (b *BookRepository) GetBook(id string) []*Book {
	var book []*Book
	strid := id
	intid := 0
	_, err := fmt.Sscan(strid, &intid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intid)
	rows, err := b.db.Query("SELECT * FROM books WHERE id=$1", intid)
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

func (b *BookRepository) DeleteBook(id string) {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func (b *BookRepository) PostBook(bk Book) {
	fmt.Println(bk.Name, bk.Price)
	res, err := b.db.Exec("INSERT INTO books (name, author, publisher, count, is_new, price, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", bk.Name, bk.Author, bk.Publisher, bk.Count, bk.IsNew, bk.Price, bk.CreatedAt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(res)
}
