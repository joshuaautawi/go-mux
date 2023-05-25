package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaautawi/warehouse/cmd/controllers"
	"github.com/joshuaautawi/warehouse/db"
)


func main(){
	db.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/products",controllers.Find).Methods("GET")
	r.HandleFunc("/products/{id}",controllers.FindOne).Methods("GET")
	r.HandleFunc("/products",controllers.Create).Methods("POST")
	r.HandleFunc("/products/{id}",controllers.Update).Methods("PUT")
	r.HandleFunc("/products",controllers.Delete).Methods("DELETE")
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000",r)
}