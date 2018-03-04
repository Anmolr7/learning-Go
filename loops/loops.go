package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 50; i += 5 { //Normal For loop
		sum += i
	}

	fmt.Printf("sum = %d\n", sum)

	sum = 0

	for sum < 50 { //While loop
		sum += 10
	}

	fmt.Println("sum = ", sum)

	// for sum<50 {							//Forever
	// 	sum+=10;
	// }
}
