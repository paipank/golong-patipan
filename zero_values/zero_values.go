package main

import "fmt"

func main() {
	var msg string
	var age int
	var price float64
	var check bool
	// rune ใช้สำหรับเช็คบิต type int
	var r rune

	fmt.Printf("msg: %s -- type: %T\n", msg, msg)

	// อยากเห็น str
	fmt.Printf("msg: %q -- type: %T\n", msg, msg)

	fmt.Printf("age: %d -- type: %T\n", age, age)
	fmt.Printf("price: %.2f -- type: %T\n", price, price)
	fmt.Printf("check: %t -- type: %T\n", check, check)
	fmt.Printf("r: %d -- type: %T\n", r, r)

}
	