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
	elemNum, err := getElemStringVal(r)
	if err != nil {
		log.Printf("The value equals: %v", elemNum)
		slices := services.CreateSlice(elemNum)
		json.NewEncoder(w).Encode(slices)
	} else {
		returnServerError(w, err, "Error on creating slice: %v")
	}

}

func printArray(w http.ResponseWriter, r *http.Request) {
	elemNum, err := getElemStringVal(r)
	log.Printf("The value equals: %v", elemNum)
	array, err := services.CreateArray(elemNum)
	if err != nil {
		returnServerError(w, err, "Error on creating array: %v")
	} else {
		json.NewEncoder(w).Encode(array)
	}
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
	http.HandleFunc("/array", printArray)
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8081", nil)
}

//some helper functions for handling http
func getElemStringVal(r *http.Request) (int, error) {
	elemStringVal := r.URL.Query().Get("elemNum")
	log.Printf(elemStringVal)
	elemNum, err := strconv.Atoi(elemStringVal)
	return elemNum, err

}

func returnServerError(w http.ResponseWriter, err error, message string) {
	log.Printf(message, err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, message, err.Error())
}
