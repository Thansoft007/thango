package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

var employees []employee

func init() {
	log.Println("init started ...")
	employees = []employee{employee{Age: 27, Name: "Thannasi", Salary: 20000}, employee{Name: "Annadurai", Age: 25, Salary: 30000}}
}

func getEmployees(resp http.ResponseWriter, req *http.Request) {
	result, error := json.Marshal(employees)
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("error "))
		return
	}
	resp.Write(result)
	resp.WriteHeader(http.StatusOK)
}

func getEmployees1(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(employees)
}

func sayHai(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hi,am alive")
}

func createEmployee(res http.ResponseWriter, req *http.Request) {
	var emp employee
	json.NewDecoder(req.Body).Decode(&emp)
	fmt.Println("emp " + emp.Name)
	employees = append(employees, emp)
	json.NewEncoder(res).Encode(employees)
}
