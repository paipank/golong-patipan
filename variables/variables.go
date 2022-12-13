package main

import "fmt"

func main() {
	// Workshop: Println
	// Output:
	// เรื่อง: Avengers: Endgame
	// ปี: 2019
	// เรดติ้ง: 8.4
	// ชุปเปอร์ฮีโร่: true

	// TODO: write code below.

	var movie string = "Avengers Endgame."
	var year int = 2019
	var rating float32 = 8.4
	var genre string = "Sci-Fi"
	var superhero bool = true
	
	fmt.Println("เรื่อง:", movie)
	fmt.Println("ปี:", year)
	fmt.Println("เรดติ้ง:", rating)
	fmt.Println("ประเภท:", genre)
	fmt.Println("ชุปเปอร์ฮีโร่:", superhero)


	// Row string
	row_str := `เรื่อง: Avengers: Endgame
ปี: 2019
เรดติ้ง: 8.4
ชุปเปอร์ฮีโร่: true` 

	fmt.Println(row_str)
	
}
