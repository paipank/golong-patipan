package main

import (
	"fmt"
	"math"
)

func add(x float64, y float64) (float64, string) {
	fmt.Println("Hello, World!", x, y)
	v := x + y
	return v, "Hello!!"
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// first class variables
var add_2 func(float64, float64) float64 = func(x float64, y float64) float64 {
	return x + y
}

func compute(fn func(float64, float64) float64) float64{
	v := fn(42, 13)
	return v
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func adder() (func () int, func() int) {
	sum := 0
	return func() int {
		sum += 1
		return sum
	}, func() int {
		return 2
	}
}

func main() {
	a, b := add(42, 13)
	fmt.Println(a, b)

	c, d := swap("Hello", "World")
	fmt.Println(c, d)

	e, f := split(24)
	fmt.Println(e, f)

	fmt.Printf("add_2 --Signature: %T\n", add_2)

	r := compute(add_2)
	fmt.Println("r:", r)

	x := compute(hypot)
	fmt.Println("r:", x)

	z, y := adder()
	fmt.Println("z:", z())
	// reference
	fmt.Println("z:", z())
	fmt.Println("y:", y())
}

