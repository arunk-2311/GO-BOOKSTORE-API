package routes

import (
	"github.com/arunk-2311/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Call the function create book if the http url is /book/ and the method used is post
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
}
