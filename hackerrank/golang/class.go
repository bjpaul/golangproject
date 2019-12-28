package main

import (
	"fmt"
	"hackerrank/golang/classes"
)

func main() {
	e := classes.New()
	e.SetEmpId(1)
	e.SetFirstName("Bijoy")
	e.SetLastName("Paul")
	fmt.Println(e)
	fmt.Println(e.GetEmpId())
	fmt.Println(e.GetFirstName())
	fmt.Println(e.GetLastName())
}