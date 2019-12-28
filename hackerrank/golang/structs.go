package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) show()  {
	fmt.Println(v.X, v.Y)
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v.show()
}
