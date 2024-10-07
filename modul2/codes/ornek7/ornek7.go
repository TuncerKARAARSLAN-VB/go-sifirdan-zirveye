package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println("Eşit:", a == b)            // false
	fmt.Println("Eşit Değil:", a != b)      // true
	fmt.Println("Büyüktür:", a > b)         // false
	fmt.Println("Küçüktür:", a < b)         // true
	fmt.Println("Büyük veya Eşit:", a >= b) // false
	fmt.Println("Küçük veya Eşit:", a <= b) // true
}
