package main

import (
	"fmt"
)

type account struct {
	account_id int
	account_holder string
}

func type_print(input interface{}) {
	switch input.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case account:
		fmt.Println("account type")
	default:
		fmt.Println("default type")
	}
}

func main() {
	type_print(10)
	type_print("SidGod")
	type_print(account{})
	type_print(3.142)

	fmt.Printf("%T\n", 10)
	fmt.Printf("%T\n", "SidGod")
	fmt.Printf("%T\n", account{})
	fmt.Printf("%T\n", 3.142)
}