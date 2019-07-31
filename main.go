package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Joł, mój pierwszy request w GO!")
}

func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To jest inny route")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", help)
	http.ListenAndServe(":8081", nil)
}
