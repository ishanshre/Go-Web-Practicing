package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // w http.ResponseWrite is what we pass to front end and r is requst and *http.Request is the request pointer
	if r.URL.Path != "/hello" {
		fmt.Fprintf(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Fprintf(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello! welcome to go universe")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static")) // specify the static files directory
	http.Handle("/", fileserver)                        // refers to the index.html page
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting the server at 8080\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
