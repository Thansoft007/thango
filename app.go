package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8010"
	router := mux.NewRouter()
	router.HandleFunc("/", sayHai).Methods("GET")
	router.HandleFunc("/getEmployees", getEmployees).Methods("GET")
	router.HandleFunc("/getEmployeesv2", getEmployees1).Methods("GET")
	router.HandleFunc("/createEmployee", createEmployee).Methods("POST")
	log.Println("server up and running on port " + port)
	log.Fatalln(http.ListenAndServe(port, router))
}
