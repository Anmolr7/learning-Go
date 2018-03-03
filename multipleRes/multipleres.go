package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a := "hello" //"a :=" is shorthand for "var a ="
	b := "world"

	fmt.Println("a =", a, "b =", b)

	a, b = swap(a, b)

	fmt.Println("after swap a=", a, "b=", b)

}
