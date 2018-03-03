package main

import "fmt"

func anyfunc(sum int) (x, y int) {
	x = sum * 3 / 5
	y = sum - x
	return
}

func main() {
	fmt.Println(anyfunc(12))
}
