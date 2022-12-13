package main

import "fmt"

func main() {
	
	movie := "Avengers Endgame."
	year := 2019
	rating := 8.4
	genre := "Sci-Fi"
	superhero := true
	fav := '😊'
	
	fmt.Printf("เรื่อง: %s\n", movie)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรดติ้ง: %.1f\n", rating)
	fmt.Printf("ประเภท: %s\n", genre)
	fmt.Printf("ชุปเปอร์ฮีโร่: %t\n", superhero)
	fmt.Printf("เรื่องโปรด: %c\n", fav)

	// Row string
	row_str := `เรื่อง: Avengers: Endgame
ปี: 2019
เรดติ้ง: 8.4
ชุปเปอร์ฮีโร่: true` 

	fmt.Printf("สติ่งดิบ: %s", row_str)

}
