package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println("Ve:", a && b)   // false
	fmt.Println("Veya:", a || b) // true
	fmt.Println("Değil:", !a)    // false
}
