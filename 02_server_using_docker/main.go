package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	PORT := 4000

	// handle the root link "localhost:4000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 'w' (response which will be sent back to the user), 'r' (request which user sends to the go lang server) it a pointer.
		fmt.Fprintf(w, "Hello, you are at path: %q\n", html.EscapeString(r.URL.Path)) // URL is "localhost:4000" and Path is '/'
	}) // it is used to handle routes.

	// handle the link "localhost:4000/hi"
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Welcome\n")
	})

	// starting the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil) // it will start the server at port no. 4000
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
