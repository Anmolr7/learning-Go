package main

import "fmt"

func main() {
	var a int = 4
	if b := a * a; b < 16 {
		fmt.Printf("square of a is less than 16\n")
	} else if b == 16 {
		fmt.Printf("square of a is equal to 16\n")
	} else {
		fmt.Printf("square of a is greater than 16\n")
	}

	switch b := a * a; b {
	case 16:
		fmt.Println("square of a is 16")
	default:
		fmt.Println("square of a is not 16")
	}

	switch b := a * a; {
	case b < 16:
		fmt.Printf("square of a is less than 16\n")
	case (b == 16):
		fmt.Printf("square of a is equal to 16\n")
	default:
		fmt.Printf("square of a is greater than 16\n")
	}
}
