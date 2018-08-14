package router

import (
	"encoding/json"
	"net/http"

	"github.com/Matias-Barrios/RESTAPIGolangSkeleton/data"
	"github.com/Matias-Barrios/RESTAPIGolangSkeleton/helpers"
	"github.com/julienschmidt/httprouter"
)

// Health is the handler for getting app information. For instance, it might be useful to inquiery which server IP is the app replying from
func health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	myresponse := helpers.Health_Response{
		Name:   "Golang REST API",
		Status: "healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myresponse)

}

func getEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Employees)
}

func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", health)
	router.GET("/employee", getEmployee)
	return router
}
