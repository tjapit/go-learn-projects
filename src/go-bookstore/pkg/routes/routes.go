package routes

import (
	"github.com/gorilla/mux"
	"github.com/tjapit/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
