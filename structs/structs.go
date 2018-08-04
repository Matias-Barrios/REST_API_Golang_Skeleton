package structs

import
(
	"github.com/Matias-Barrios/REST_API_Golang_Skeleton/employee"
)

type Health_Response struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// employees is the slice we will be getting employees, deleting and adding them from
var Employees = &[]employee.Employee{

	employee.Employee{
		Id:       1,
		Name:     "Matias",
		Lastname: "Barrios",
		Age:      31,
		Phones:   []string{"5552323", "5554232"},
		Email:    "soymatiasbarrios@gmail.com",
	},
	employee.Employee{
		Id:       2,
		Name:     "Peperino",
		Lastname: "Pomoro",
		Age:      45,
		Phones:   []string{"099111222", "25249867"},
		Email:    "elpepepo@hotmail.com",
	},
	employee.Employee{
		Id:       3,
		Name:     "Carlitos",
		Lastname: "Bala",
		Age:      65,
		Phones:   []string{"3571113", "6660000"},
		Email:    "carlitos_elmas@gmail.com",
	},
}
