package main

import (
	"fmt"
)

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Pazartesi")
	case 2:
		fmt.Println("Salı")
	case 3:
		fmt.Println("Çarşamba")
	default:
		fmt.Println("Geçersiz gün")
	}
}
