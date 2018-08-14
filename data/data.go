package data

import (
	"github.com/Matias-Barrios/RESTAPIGolangSkeleton/employee"
)

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
}
