package main

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Println(int(r.URL.Query().Get("a")[0]))
}

func subtraction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Subtraction")
	fmt.Println("Endpoint hit: subtraction")
}

func multiplication(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Multiplication")
	fmt.Println("Endpoint hit: multiplication")
}

func division(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Division")
	fmt.Println("Endpoint hit: division")
}
func main() {
	handleRequests()
}
