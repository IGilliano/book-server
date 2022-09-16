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
	r.HandleFunc("/book", bookImplemetation.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", bookImplemetation.GetBook).Methods("GET")
	r.HandleFunc("/book/{id:[0-9]+}", bookImplemetation.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book", bookImplemetation.PostBook).Methods("POST")
	r.HandleFunc("/book/{id:[0-9]+}", bookImplemetation.UpdateBook).Methods("PUT")
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
