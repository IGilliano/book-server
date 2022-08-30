package service

import (
	"book-server/repository"
	"fmt"
	"time"
)

type BookRequest struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Count     int    `json:"count"`
	Quantity  int    `json:"quantity"`
	IsNew     bool   `json:"is_new"`
	Price     int    `json:"price"`
}

type IBookService interface {
	GetBooks() []*repository.Book
	GetBook(id string) []*repository.Book
	PostBook(book BookRequest)
	DeleteBook(id string)
}

type BookService struct {
	bookRepository repository.IBookRepository
}

func NewBookService(r repository.IBookRepository) IBookService {
	return &BookService{bookRepository: r}
}

func (s *BookService) GetBooks() []*repository.Book {
	fmt.Printf("got /store request\n")
	books := s.bookRepository.GetBooks()
	return books
}

func (s *BookService) GetBook(id string) []*repository.Book {
	fmt.Println("got /store/<id> request\n")
	book := s.bookRepository.GetBook(id)
	return book
}

func (s *BookService) DeleteBook(id string) {
	fmt.Println("got /delete request\n")
	s.bookRepository.DeleteBook(id)
}

func FromRequestToBook(r BookRequest) repository.Book {
	book := repository.Book{
		Id:        0,
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

func (s *BookService) PostBook(book BookRequest) {
	fmt.Println("got /post request\n")
	s.bookRepository.PostBook(FromRequestToBook(book))
}
