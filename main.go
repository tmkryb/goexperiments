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
	users, err := services.GetAllUsers()
	if err != nil {
		returnServerError(w, err, "Cannot read users: %v")
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	getJsonBody(r, user)
	log.Printf("new user to add: %v", user)
	user, err := services.AddNewUser(user)
	if err != nil {
		returnServerError(w, err, "Error occured on accessing database: %v")
	} else {
		fmt.Fprintf(w, "User sucessfully added to database: %v", user.ID)
	}

}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
	http.HandleFunc("/slice", printSlice)
	http.HandleFunc("/array", printArray)
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8081", nil)
}

//some helper functions for handling http
func getElemStringVal(r *http.Request) (int, error) {
	elemStringVal := r.URL.Query().Get("elemNum")
	log.Printf(elemStringVal)
	elemNum, err := strconv.Atoi(elemStringVal)
	return elemNum, err
}

//get any query value as string value
func getQueryValue(r *http.Request, keyName string) string {
	return r.URL.Query().Get(keyName)
}

//Get json body, and decodes it to struct
func getJsonBody(r *http.Request, out *models.User) error {
	log.Printf("Json body is %v", r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(out)
	return err
}

func returnServerError(w http.ResponseWriter, err error, message string) {
	log.Printf(message, err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, message, err.Error())
}
