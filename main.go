package main

import (
	"fmt"
	"net/http"
)

var Users []User

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Joł, mój pierwszy request w GO!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To jest inny route")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Tomasz", Email: "tmkryb@gmail.com"},
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8081", nil)
}
