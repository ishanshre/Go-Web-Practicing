package main

import (
	"fmt"
	"log"
	"net/http" // for http request , response and many more
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "404 Not Found, Status Code: %v", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Fprintf(w, "Method Not Supported, Status Code: %v", http.StatusNotAcceptable)
		return
	} else {
		fmt.Fprintf(w, "Hello, You requested a get method on this base url: %s, Status Code : %v", r.URL.Path, http.StatusAccepted)
	}
}

func main() {
	http.HandleFunc("/", helloHandle)
	fmt.Println("Starting the server at 127.0.0.1:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
