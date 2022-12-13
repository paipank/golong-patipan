package main

import "fmt"

func main() {
	num := 34

	if num == 34 && (num > 30 || num < 39) {
		fmt.Println("Yes!!, it's Thirty four.")
	}

	ratings := 8.4

	if ratings < 5.0 {
		fmt.Println("Disappointed 😒")
	} else if ratings >= 5.0 && ratings < 7.0 {
		fmt.Println("Normal 😊")
	} else if ratings >= 7.0 && ratings < 10.0 {
		fmt.Println("Good 😁")
	} else {
		fmt.Println("🙌🙌🙌🙌")
	}
}
