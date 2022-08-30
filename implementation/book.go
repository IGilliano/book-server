package implementation

import (
	"book-server/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

type IBookImplementation interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	PostBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

type BookImplementation struct {
	bookService service.IBookService
}

func NewBookImplementation(s service.IBookService) IBookImplementation {
	return &BookImplementation{bookService: s}
}

func (i *BookImplementation) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := i.bookService.GetBooks()
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
	vars := mux.Vars(r)
	id := vars["id"]
	book := i.bookService.GetBook(id)
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
	vars := mux.Vars(r)
	id := vars["id"]
	i.bookService.DeleteBook(id)
}

func (i *BookImplementation) PostBook(w http.ResponseWriter, r *http.Request) {

	var book service.BookRequest
	jsonbookstring := `{"name": "Idiot", "author": "Dostoevskiy", "publisher": "AST", "quantity": 1, "is_new": true, "price": 399}`
	fmt.Println(jsonbookstring)

	err := json.Unmarshal([]byte(jsonbookstring), &book)
	if err != nil {
		fmt.Println("Panic is right here in implementation")
		panic(err)
	}
	fmt.Println(book)

	i.bookService.PostBook(book)

	//resp, err := http.Post("http://localhost:8000", "application/json", bytes.NewBufferString(jsonbookstring))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp)
}
