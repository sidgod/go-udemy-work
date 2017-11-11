package main

import "fmt"

func main()  {
	a := 10
	b := 3.142
	c := true
	d := "Voltron"

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)
	fmt.Printf("%v is of type %T\n", d, d)

	var p int
	var q float64
	var r bool
	var s string

	fmt.Printf("%v is of type %T\n", p, p)
	fmt.Printf("%v is of type %T\n", q, q)
	fmt.Printf("%v is of type %T\n", r, r)
	fmt.Printf("%v is of type %T\n", s, s)
}