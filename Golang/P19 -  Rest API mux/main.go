package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
)

// Book struct represents a book
type Book struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Author string  `json:"author"`
}

// books slice holds the collection of books
var books []Book

func main() {
        // Initialize gorilla mux router
        router := mux.NewRouter()

        // Define routes for CRUD operations
        router.HandleFunc("/books", getBooks).Methods(http.MethodGet)
        router.HandleFunc("/books/{id}", getBookByID).Methods(http.MethodGet)
        router.HandleFunc("/books", createBook).Methods(http.MethodPost)
        router.HandleFunc("/books/{id}", updateBook).Methods(http.MethodPut)
        router.HandleFunc("/books/{id}", deleteBook).Methods(http.MethodDelete)

        fmt.Println("Server listening on port 8080")
        log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id := params["id"]

        for _, book := range books {
                if book.ID == id {
                        w.Header().Set("Content-Type", "application/json")
                        json.NewEncoder(w).Encode(book)
                        return
                }
        }

        w.WriteHeader(http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
        var newBook Book
        err := json.NewDecoder(r.Body).Decode(&newBook)
        if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, "Error decoding book: %v", err)
                return
        }

        books = append(books, newBook)
        w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id := params["id"]

        var updatedBook Book
        err := json.NewDecoder(r.Body).Decode(&updatedBook)
        if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, "Error decoding book: %v", err)
                return
        }

        for i, book := range books {
                if book.ID == id {
                        books[i] = updatedBook
                        w.WriteHeader(http.StatusOK)
                        return
                }
        }

        w.WriteHeader(http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id := params["id"]

        for i, book := range books {
                if book.ID == id {
                        books = append(books[:i], books[i+1:]...)
                        w.WriteHeader(http.StatusNoContent)
                        return
                }
        }

        w.WriteHeader(http.StatusNotFound)
}