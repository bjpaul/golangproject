package main

import "fmt"

func adder() func(a int, b int) int{
	return func(a int, b int) int{
		return a+b
	}
}

func main() {

	fn := adder()
	print(2, 3, fn)
	
}

func print(a int, b int, fn func(a int, b int) int)  {
	fmt.Println(fn(a, b))
}