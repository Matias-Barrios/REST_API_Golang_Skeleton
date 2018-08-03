package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Matias-Barrios/RESTAPIGolangSkeleton/employee"
	"github.com/julienschmidt/httprouter"
)

type Health_Response struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// PORT defines what port the app listens on
const PORT int = 8080

// employees is the slice we will be getting employees, deleting and adding them from
var employees = &[]employee.Employee{

	employee.Employee{
		Id:       1,
		Name:     "Matias",
		Lastname: "Barrios",
		Age:      31,
		Phones:   []string{"5552323", "5554232"},
		Email:    "soymatiasbarrios@gmail.com",
	},
}

// Health is the handler for getting app information. For instance, it might be useful to inquiery which server IP is the app replying from
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	myresponse := Health_Response{
		Name:   "Golang REST API",
		Status: "healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myresponse)

}

func getEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", Health)
	router.GET("/employee", getEmployee)
	return router
}

func main() {

	fmt.Println("Server up and running. Listening in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), getRouter()))
}
