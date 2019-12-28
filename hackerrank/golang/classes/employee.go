package classes

import "fmt"

type Employee struct {
	firstName string
	lastName string
	empId int
}

func New() Employee{
	return Employee{}
}

func NewValues(firstName string, lastName string, empId int) Employee{
	return Employee{firstName, lastName, empId}
}

func (e Employee)SetFirstName(firstName string)  {
	e.firstName = firstName
	fmt.Println(e)
}

func (e Employee)SetLastName(lastName string)  {
	e.lastName = lastName
	fmt.Println(e)
}

func (e Employee)SetEmpId(empId int)  {
	e.empId = empId
	fmt.Println(e)
}

func (e Employee)GetFirstName()  string{
	return e.firstName
}

func (e Employee)GetLastName()  string{
	return e.lastName
}

func (e Employee)GetEmpId()  int{
	return e.empId
}