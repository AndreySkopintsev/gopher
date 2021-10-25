package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/add", calculation)
	http.HandleFunc("/api/sub", calculation)
	http.HandleFunc("/api/mul", calculation)
	http.HandleFunc("/api/div", calculation)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func calculation(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Path
	var first int
	var second int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = a
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = b
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	switch u {
	case "/api/add":
		fmt.Fprintf(w, "The sum is %d", first+second)
	case "/api/sub":
		fmt.Fprintf(w, "Subtraction result is %d", first-second)
	case "/api/mul":
		fmt.Fprintf(w, "Multiplication result is %d", first*second)
	case "/api/div":
		fmt.Fprintf(w, "Division result is %d", first/second)
	}

}

func main() {
	handleRequests()
}
