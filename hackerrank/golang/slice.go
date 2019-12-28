package main

import "fmt"

func main() {
	arr := [3]int{2, 3 , 5}
	slc := arr[:2]
	fmt.Println(arr)
	fmt.Println(slc)
	slc[1] = 23
	fmt.Println(arr)
	fmt.Println(slc)
	fmt.Println(len(slc), cap(slc))
	fmt.Println(append(slc, 10))

}
