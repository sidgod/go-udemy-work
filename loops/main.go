package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("i = %d\n", i)
	}

	j := 0
	for ; j < 100; j++ {
		fmt.Printf("j = %d\n", j)
	}

	var k int = 0
	for ;; k++ {
		if k >= 100 {
			break
		}
		fmt.Printf("k = %d\n", k)
	}

	l := 0
	for l < 10 {
		fmt.Printf("l = %d\n", l)
		l++
	}

	m := 0
	for {
		if m >= 10 {
			break
		}
		fmt.Printf("m = %d\n", m)
		m++
	}

	n := 0
	for {
		if n >= 10 {
			break
		}
		n++
		if n % 2 == 0 {
			continue
		}
		fmt.Printf("n = %d\n", n)
	}

}