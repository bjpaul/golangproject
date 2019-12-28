package main

import "fmt"

func main() {
	fmt.Println(GetNthFibRecursive(6	))
}

func GetNthFib(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	a, b := 0, 1
	temp := b
	for i := 1; i <= n - 2  ; i ++{
		temp = b
		b = a + b
		a = temp
	}
	return b
}

func GetNthFibRecursive(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return GetNthFibRecursive(n-1) + GetNthFibRecursive(n-2)
}


