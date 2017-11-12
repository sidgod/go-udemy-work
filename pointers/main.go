package main

import "fmt"

func main() {
	a := 9

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	var b *int = &a

	fmt.Println("value of b = ", b)
	fmt.Println("value of *b = ", *b)

	var c **int = &b // Pointer to pointer :-)

	fmt.Println("value of c = ", c)
	fmt.Println("value of *c = ", *c)
}