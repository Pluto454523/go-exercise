package main

import "fmt"

func main() {
	// Workshop: switch case
	// กำหนด: ให้แก้โจทย์นี้ด้วยการใช้ switch case เท่านั้น

	// Output:
	// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
	// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
	// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
	// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"

	// TODO: write code below.
	ratings := 8.4
	switch {
	case ratings < 5.0:
		fmt.Println("Disappointed 😞")
	case ratings >= 5.0 && ratings < 7.0:
		fmt.Println("Normal 😐")
	case ratings >= 7.0 && ratings < 10.0:
		fmt.Println("Good 🥰")
	default:
		fmt.Println("🤷🤷🤷🤷")
	}
}
