package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gajjar-Mohit/bookstore-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
