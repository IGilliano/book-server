package main

import (
	"book-server/implementation"
	"book-server/repository"
	"book-server/service"

	"github.com/gorilla/mux"

	//"bytes"
	//"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository)
	bookImplemetation := implementation.NewBookImplementation(bookService)
	r.HandleFunc("/store", bookImplemetation.GetBooks)
	r.HandleFunc("/store/{id:[0-9]+}", bookImplemetation.GetBook)
	r.HandleFunc("/delete/{id:[0-9]+}", bookImplemetation.DeleteBook)
	r.HandleFunc("/post", bookImplemetation.PostBook)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
