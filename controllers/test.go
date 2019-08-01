package controllers

import (
	"encoding/json"
	"fmt"
	"goexperiments/services"
	"goexperiments/tools"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Joł, mój pierwszy request w GO!")
}

func PrintSlice(w http.ResponseWriter, r *http.Request) {
	elemNum, err := tools.GetElemStringVal(r)
	if err != nil {
		tools.ReturnServerError(w, err, "Error on creating slice: %v")
	} else {
		log.Printf("The value equals: %v", elemNum)
		slices := services.CreateSlice(elemNum)
		json.NewEncoder(w).Encode(slices)
	}
}

func PrintArray(w http.ResponseWriter, r *http.Request) {
	elemNum, err := tools.GetElemStringVal(r)
	log.Printf("The value equals: %v", elemNum)
	array, err := services.CreateArray(elemNum)
	if err != nil {
		tools.ReturnServerError(w, err, "Error on creating array: %v")
	} else {
		json.NewEncoder(w).Encode(array)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To jest inny route")
}
