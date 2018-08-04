package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/Matias-Barrios/REST_API_Golang_Skeleton/methods"
)

// PORT defines what port the app listens on
const PORT int = 8080

func main() {

	fmt.Println("Server up and running. Listening in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), methods.GetRouter()))
}
