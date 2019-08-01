package controllers

import (
	"encoding/json"
	"fmt"
	"goexperiments/models"
	"goexperiments/services"
	"goexperiments/tools"
	"log"
	"net/http"
	"strconv"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		tools.ReturnServerError(w, err, "Cannot read users: %v")
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	tools.GetJsonBody(r, user)
	log.Printf("new user to add: %v", user)
	user, err := services.AddNewUser(user)
	if err != nil {
		tools.ReturnServerError(w, err, "Error occured on accessing database: %v")
	} else {
		fmt.Fprintf(w, "User sucessfully added to database: %v", user.ID)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := tools.GetQueryValue(r, "id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		tools.ReturnServerError(w, err, "Error on Atoi function: %v")
	} else {
		err = services.DeleteUser(userId)
		//I don't like it, if in if, but I will leave it for now
		if err != nil {
			tools.ReturnServerError(w, err, "Error on deleting user: %v")
		} else {
			fmt.Fprintf(w, "User %v deleted sucessfully from database", userId)
		}
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	}
}
