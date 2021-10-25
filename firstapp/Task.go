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
	http.HandleFunc("/api/add", addition)
	http.HandleFunc("/api/sub", subtraction)
	http.HandleFunc("/api/mul", multiplication)
	http.HandleFunc("/api/div", division)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func addition(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = a
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = b
		fmt.Fprintf(w, "The sum is %d", first+second)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}
}

func subtraction(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = a
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = b
		fmt.Fprintf(w, "The sum is %d", first-second)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}
}

func multiplication(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = a
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = b
		fmt.Fprintf(w, "The sum is %d", first*second)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}
}

func division(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = a
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = b
		fmt.Fprintf(w, "The sum is %d", first/second)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}
}
func main() {
	handleRequests()
}
