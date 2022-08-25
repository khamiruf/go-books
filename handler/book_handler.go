package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/khamiruf/go-toilets/model"
	"github.com/khamiruf/go-toilets/service"
)

type BookStore struct {
	bookList []string
	authors  []string
}

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.getBook(w)
	case http.MethodPost:
		b.postBook(w, r)
	case http.MethodPut:
		b.updateBook(w, r)
	case http.MethodDelete:
		b.deleteBook(w, r)
	}

}

func (b *BookHandler) getBook(w http.ResponseWriter) {
	books, err := b.service.GetBooks()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Println(err)
	}
}

func (b *BookHandler) postBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rowsAffected, err := b.service.AddBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Book %s added successfully (%d rows affected)\n", book.Title, rowsAffected)
}

func (b *BookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("error converting id to int: %s\n", err)
	}

	var book model.Book
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rowsAffected, err := b.service.UpdateBook(&book, intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Book %d updated successfully (%d rows affected)\n", book.ID, rowsAffected)
}

func (b *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("error converting id to int: %s\n", err)
	}
	rowsAffected, err := b.service.DeleteBook(int64(intID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Book %s deleted succesfully (%d rows afected)", id, rowsAffected)
}
