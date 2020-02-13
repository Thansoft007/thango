package main

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", sayHai).Methods("GET")
	router.HandleFunc("/getEmployees", getEmployees).Methods("GET")
	router.HandleFunc("/getEmployeesv2", getEmployees1).Methods("GET")
	router.HandleFunc("/createEmployee", createEmployee).Methods("POST")
	log.Println("server up and running on port " + port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
