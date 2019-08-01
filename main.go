package main

import (
	"goexperiments/controllers"
	"goexperiments/models"
	"net/http"
)

var users []models.User

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/about", controllers.About)
	http.HandleFunc("/slice", controllers.PrintSlice)
	http.HandleFunc("/array", controllers.PrintArray)
	http.HandleFunc("/users", controllers.HandleUsers)
	http.ListenAndServe(":8081", nil)
}
