package main

import (
	"fmt"
)

func main() {
	switch x := 2; x {
	case 1:
		fmt.Println("Bir")
	case 2:
		fmt.Println("İki")
		fallthrough // bu durumda "Üç" de yazdırılır
	case 3:
		fmt.Println("Üç")
	default:
		fmt.Println("Geçersiz")
	}
}
