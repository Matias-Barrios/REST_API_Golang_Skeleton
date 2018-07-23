package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type healthResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

const PORT int = 8080

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	myresponse := healthResponse{
		Name:   "Golang REST API",
		Status: "healthy",
	}
	fmt.Println(myresponse.Status)
	js, err := json.Marshal(myresponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", Health)
	router.GET("/hello/:name", Hello)
	return router
}

func main() {

	fmt.Println("Server up and running. Listening in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), getRouter()))
}
