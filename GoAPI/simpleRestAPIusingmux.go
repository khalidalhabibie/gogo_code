package main

import (
	"encoding/json"
	"log"

	//"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//book struct(model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:title`
	Author *Author `json:author`
}

//author struct(model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:lastname`
}

//Init books var a slices book structure

var books []Book

// get books()
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(books)
}

//get a book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	//get params
	params := mux.Vars(r)

	//loop for find a book
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(&book)
}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}

	}
}

//delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	//get params
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init Router
	r := mux.NewRouter()

	//simple json data

	books = append(books, Book{ID: "1", Isbn: "12345", Title: "kisah seorang pelaut", Author: &Author{Firstname: "Anip", Lastname: "anipan"}})
	books = append(books, Book{ID: "2", Isbn: "5678", Title: "buku baca tulis cerdas", Author: &Author{Firstname: "Budi", Lastname: "saja"}})
	//router Handlers/ endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//turn  of http listen and serve
	log.Fatal(http.ListenAndServe(":8000", r))
}

/*
GET
get all book
http://localhost:8000/api/books
*/

/*
GET
get a book
http://localhost:8000/api/books/{id}
*/

/*
POST
create a book
http://localhost:8000/api/book
body:
	{
  "isbn":"33332",
   "title":"tiga teman",
   "author":{"firstname":"Untung",  "lastname":"Saja"}
 }

*/

/*
PUT
update book
http://localhost:8000/api/books/{id}
body
	{
  "isbn":"33332",
   "title":"empat teman",
   "author":{"firstname":"Untung",  "lastname":"Saja"}
 }

*/

/*
DELETE
http://localhost:8000/api/books/{id}

*/
