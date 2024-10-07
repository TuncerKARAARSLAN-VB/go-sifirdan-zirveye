package main

import (
	"fmt"
)

func main() {
	if y := 20; y < 15 {
		fmt.Println("Y değeri 15'ten küçük")
	} else {
		fmt.Println("Y değeri 15 veya daha büyük")
	}
}
