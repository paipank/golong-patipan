package main

import "fmt"

func main() {

	switch today := "Saturday"; today {
	case "Saturday":
		fmt.Println("todat is Saturday")
		fallthrough
	case "Monday", "Tueday":
		fmt.Println("today is weekdays")
	default:
		fmt.Printf("สวัสดี %#v อยู่ดีมีแฮงเดย์\n", today)
	}

	num := 1

	switch {
	case num < 0:
		fmt.Println("nagative number.")
	case num == 0:
		fmt.Println("zero.")
	case num > 0:
		fmt.Println("positive number.")
	default:
		fmt.Println("👌👌👌👌")
	}

	ratings := 8.4

	switch {
	case ratings < 5.0:
		fmt.Println("Disappointed 😒")
	case ratings >= 5.0 && ratings < 7.0:
		fmt.Println("Normal 😊")
	case  ratings >= 7.0 && ratings < 10.0:
		fmt.Println("Good 😁")
	default:
		fmt.Println("🙌🙌🙌🙌")
	}

}