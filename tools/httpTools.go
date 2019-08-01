package tools

import (
	"encoding/json"
	"fmt"
	"goexperiments/models"
	"log"
	"net/http"
	"strconv"
)

//some helper functions for handling http
func GetElemStringVal(r *http.Request) (int, error) {
	elemStringVal := r.URL.Query().Get("elemNum")
	log.Printf(elemStringVal)
	elemNum, err := strconv.Atoi(elemStringVal)
	return elemNum, err
}

//get any query value as string value
func GetQueryValue(r *http.Request, keyName string) string {
	return r.URL.Query().Get(keyName)
}

//Get json body, and decodes it to struct
func GetJsonBody(r *http.Request, out *models.User) error {
	log.Printf("Json body is %v", r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(out)
	return err
}

func ReturnServerError(w http.ResponseWriter, err error, message string) {
	log.Printf(message, err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, message, err.Error())
}
