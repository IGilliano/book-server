package service

import (
	"book-server/repository"
	"fmt"
	"time"
)

type BookForService struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Count     int    `json:"count"`
	IsNew     bool   `json:"is_new"`
	Price     int    `json:"price"`
}

type IBookService interface {
	GetBooks() ([]*repository.Book, error)
	GetBook(int) ([]*repository.Book, error)
	PostBook(BookForService) error
	DeleteBook(int) error
	UpdateBook(BookForService) error
}

type BookService struct {
	bookRepository repository.IBookRepository
}

func NewBookService(r repository.IBookRepository) IBookService {
	return &BookService{bookRepository: r}
}

func (s *BookService) GetBooks() ([]*repository.Book, error) {
	fmt.Printf("got GET request\n")
	books, err := s.bookRepository.GetBooks()
	return books, err
}

func (s *BookService) GetBook(id int) ([]*repository.Book, error) {
	fmt.Println("got GET ID request\n")
	book, err := s.bookRepository.GetBook(id)
	return book, err
}

func (s *BookService) DeleteBook(id int) error {
	fmt.Println("got DELETE request\n")
	err := s.bookRepository.DeleteBook(id)
	return err
}

func FromRequestToBook(r BookForService) repository.Book {
	book := repository.Book{
		Id:        r.ID,
		Name:      r.Name,
		Author:    r.Author,
		Publisher: r.Publisher,
		Count:     r.Count,
		IsNew:     r.IsNew,
		Price:     r.Price,
		CreatedAt: time.Now(),
	}
	return book
}

func (s *BookService) PostBook(book BookForService) error {
	fmt.Println("got POST request\n")
	err := s.bookRepository.PostBook(FromRequestToBook(book))
	return err
}

func (s *BookService) UpdateBook(book BookForService) error {
	fmt.Println("got PUT request\n")
	fmt.Println(book)
	err := s.bookRepository.UpdateBook(FromRequestToBook(book))
	return err
}
