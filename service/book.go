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
	Quantity  int    `json:"quantity"`
	IsNew     bool   `json:"is_new"`
	Price     int    `json:"price"`
}

type IBookService interface {
	GetBooks() []*repository.Book
	GetBook(id int) []*repository.Book
	PostBook(book BookForService)
	DeleteBook(id int)
	UpdateBook(book BookForService)
}

type BookService struct {
	bookRepository repository.IBookRepository
}

func NewBookService(r repository.IBookRepository) IBookService {
	return &BookService{bookRepository: r}
}

func (s *BookService) GetBooks() []*repository.Book {
	fmt.Printf("got GET request\n")
	books := s.bookRepository.GetBooks()
	return books
}

func (s *BookService) GetBook(id int) []*repository.Book {
	fmt.Println("got GET ID request\n")
	book := s.bookRepository.GetBook(id)
	return book
}

func (s *BookService) DeleteBook(id int) {
	fmt.Println("got DELETE request\n")
	s.bookRepository.DeleteBook(id)
}

func FromRequestToBook(r BookForService) repository.Book {
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

func (s *BookService) PostBook(book BookForService) {
	fmt.Println("got POST request\n")
	s.bookRepository.PostBook(FromRequestToBook(book))
}

func (s *BookService) UpdateBook(book BookForService) {
	fmt.Println("got PUT request\n")
	s.bookRepository.UpdateBook(FromRequestToBook(book))
}
