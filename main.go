package main

import (
	"encoding/json"
	"fmt"
	"goexperiments/models"
	"goexperiments/services"
	"log"
	"net/http"
	"strconv"
)

var users []models.User

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Joł, mój pierwszy request w GO!")
}

func printSlice(w http.ResponseWriter, r *http.Request) {
	elemStringVal := r.URL.Query().Get("elemNum")
	log.Printf(elemStringVal)
	elemNum, _ := strconv.Atoi(elemStringVal)

	log.Printf("The value equals: %v", elemNum)
	slices := services.CreateSlice(elemNum)
	json.NewEncoder(w).Encode(slices)
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
	http.HandleFunc("/slice", printSlice)
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8081", nil)
}
