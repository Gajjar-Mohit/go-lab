package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successfull")
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	fmt.Fprintln(w, "Name: %s", name)
	fmt.Fprintln(w, "Email: %s", email)
	fmt.Fprintln(w, "Message: %s", message)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "GET method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello there!!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	fmt.Printf("Server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
