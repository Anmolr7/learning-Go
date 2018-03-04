package main

import "fmt"

//Vertex is a demo struct (commend added to adhere to Go coding style)
type Vertex struct {
	x, y int
}

func main() {
	var v1 = Vertex{1, 2}
	v3 := Vertex{}
	var v4 = Vertex{x: 1, y: 5}
	var p = &Vertex{x: 3}

	fmt.Println(v1, v3, v4, p)
}
