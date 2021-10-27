package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const AddOperator int = 0
const SubOperator int = 1
const MulOperator int = 2
const DivOperator int = 3

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/add", addNumbers)
	http.HandleFunc("/api/sub", subNumbers)
	http.HandleFunc("/api/mul", mulNumbers)
	http.HandleFunc("/api/div", divNumbers)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

// Handlers for operations
// Addition

func addNumbers(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int
	var result int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = int(a)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = int(b)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	result = calculation(first, second, AddOperator)

	fmt.Fprintf(w, "Addition result is %d", result)

}

// Subtraction

func subNumbers(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int
	var result int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = int(a)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = int(b)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	result = calculation(first, second, SubOperator)

	fmt.Fprintf(w, "Addition result is %d", result)

}

// Multiplication

func mulNumbers(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int
	var result int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = int(a)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = int(b)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	result = calculation(first, second, MulOperator)

	fmt.Fprintf(w, "Multiplication result is %d", result)

}

// Division

func divNumbers(w http.ResponseWriter, r *http.Request) {
	var first int
	var second int
	var result int

	if a, err := strconv.Atoi(r.URL.Query().Get("a")); err == nil {
		first = int(a)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if b, err := strconv.Atoi(r.URL.Query().Get("b")); err == nil {
		second = int(b)
	} else {
		fmt.Fprintf(w, "There was an error: %d", err)
	}

	if second != 0 {
		result = calculation(first, second, DivOperator)
		fmt.Fprintf(w, "Division result is %d", result)
	} else {
		fmt.Fprintf(w, "You can't divide by zero")
	}
}

// Calculation fucntion which handles all the operation based on values queried from URL
//  and an operator constant that's being passed
func calculation(a int, b int, oper int) int {
	var result int

	switch oper {
	case 0:
		result = a + b
	case 1:
		result = a - b
	case 2:
		result = a * b
	case 3:
		result = a / b
	}

	return result
}

func main() {
	handleRequests()
}
