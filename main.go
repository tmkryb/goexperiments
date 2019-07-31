package main

import (
	"encoding/json"
	"fmt"
	"goexperiments/models"
	"net/http"
)

var users []models.User

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
	users := []models.User{
		models.User{Name: "Tomasz", Email: "tmkryb@gmail.com"},
	}
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8081", nil)
}
