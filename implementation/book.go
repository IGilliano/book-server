package implementation

import (
	"book-server/service"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BookRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Count     int    `json:"count"`
	IsNew     bool   `json:"is_new"`
	Price     int    `json:"price"`
}

func ValidateBookRequest(r BookRequest) error {
	if r.Name == "" {
		return errors.New("Name is missing")
	}
	if r.Author == "" {
		return errors.New("Author is missing")
	}
	if r.Publisher == "" {
		return errors.New("Publisher is missing")
	}
	if r.Count == 0 {
		return errors.New("Count is missing")
	}
	if r.Price == 0 {
		return errors.New("Price is missing")
	}
	return nil
}

type IBookImplementation interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	PostBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type BookImplementation struct {
	bookService service.IBookService
}

func NewBookImplementation(s service.IBookService) IBookImplementation {
	return &BookImplementation{bookService: s}
}

func GetID(r *http.Request) int {
	vars := mux.Vars(r)
	id := vars["id"]
	strid := id
	intid := 0

	_, err := fmt.Sscan(strid, &intid)
	if err != nil {
		fmt.Println(err)
	}
	return intid
}

func (i *BookImplementation) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := i.bookService.GetBooks()
	booksByte, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(booksByte)
	if errWrite != nil {
		log.Println(errWrite)
	}
}

func (i *BookImplementation) GetBook(w http.ResponseWriter, r *http.Request) {
	id := GetID(r)
	book, err := i.bookService.GetBook(id)
	bookByte, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}

	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(bookByte)
	if errWrite != nil {
		log.Println(errWrite)
	}
}

func (i *BookImplementation) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := GetID(r)
	i.bookService.DeleteBook(id)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Book with id", id, "deleted\n")
}

func (i *BookImplementation) PostBook(w http.ResponseWriter, r *http.Request) {

	var book BookRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}

	err = json.Unmarshal(reqBody, &book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}

	err = ValidateBookRequest(book)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	i.bookService.PostBook(FromBookRequestToBookForService(book, 0))
}

func FromBookRequestToBookForService(r BookRequest, id int) service.BookForService {
	book := service.BookForService{
		ID:        id,
		Name:      r.Name,
		Author:    r.Author,
		Publisher: r.Publisher,
		Count:     r.Count,
		IsNew:     r.IsNew,
		Price:     r.Price,
	}
	return book
}

func (i *BookImplementation) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var b BookRequest
	id := GetID(r)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}

	err = json.Unmarshal(reqBody, &b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, newErr := w.Write([]byte(err.Error()))
		if newErr != nil {
			log.Println(newErr)
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(b)
	i.bookService.UpdateBook(FromBookRequestToBookForService(b, id))

}
