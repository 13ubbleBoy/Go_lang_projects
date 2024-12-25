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
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

// ----------------------  response			 request
func helloHandler(w http.ResponseWriter, r *http.Request) { // every route has a response or request
	if r.URL.Path != "/hello" { // checking if requested path is correct or not
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" { // checking if the method the user requested
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // go lang know thats it has to look for index.html file.

	http.Handle("/", fileServer) // this is to for index.html

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 4000\n")
	if err := http.ListenAndServe(":4000", nil); err != nil { // 1 of the 2 values will be assigned to err
		log.Fatal(err)
	}
}
