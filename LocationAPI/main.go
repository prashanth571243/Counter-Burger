package main

import (
    _"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    
)

func GetStars(w http.ResponseWriter, r *http.Request)  {
    
}
func GetStar(w http.ResponseWriter, r *http.Request)  {
    
}
func CreateStar(w http.ResponseWriter, r *http.Request)  {
    
}
func DeleteStar(w http.ResponseWriter, r *http.Request)  {
    
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/stars", GetStars).Methods("GET")
    router.HandleFunc("/stars/{id}", GetStar).Methods("GET")
    router.HandleFunc("/stars/{id}", CreateStar).Methods("POST")
    router.HandleFunc("/stars/{id}", DeleteStar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}