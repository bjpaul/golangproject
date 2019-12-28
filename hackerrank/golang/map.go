package main

import (
	"fmt"
	"hackerrank/golang/classes"
)

func main() {
	m := make(map[int]classes.Employee)
	for i:=1;i<=3 ;i++  {
		m[i] = classes.NewValues("Bijoy", "Paul", i)
	}

	for i:=1;i<=3 ;i++  {
		v, ok := m[i]
		fmt.Println(i, "Value Exist?", ok, " | ", v)
	}

	println("************************************")
	delete(m, 2)

	for i:=1;i<=3 ;i++  {
		v, ok := m[i]
		fmt.Println(i, "Value Exist?", ok, " | ", v)
	}
}
