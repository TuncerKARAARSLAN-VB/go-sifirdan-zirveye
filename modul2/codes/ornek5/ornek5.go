package main

import "fmt"

func main() {
	// Bir kesit oluşturma
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Kesit:", numbers)

	// Kesit elemanlarına erişim
	for i, v := range numbers {
		fmt.Println("Eleman", i, ":", v)
	}
}
