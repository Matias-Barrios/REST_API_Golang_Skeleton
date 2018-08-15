package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Matias-Barrios/RESTAPIGolangSkeleton/router"
)

// PORT defines what port the app listens on
const PORT int = 8080

func main() {

	fmt.Println("Server up and running. Listening in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), router.GetRouter()))

}
