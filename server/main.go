package main

import (
	"fmt"
	"log"
	"net/http"
)

func predict(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET method.")
		params := r.URL.Query()
		for key, element := range params {
			fmt.Println("Key:", key, "=>", "Element:", element[0])
		}

	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		r.ParseForm() // Parses the request body
		params := r.Form
		for key, element := range params {
			fmt.Println("Key:", key, "=>", "Element:", element[0])
		}
		fmt.Fprintf(w, "POST method.")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/hello", predict)

	log.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
