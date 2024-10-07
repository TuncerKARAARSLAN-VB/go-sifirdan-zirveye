package main

import "fmt"

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Dizi:", numbers)

	// Dizinin elemanlarına erişim
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Eleman", i, ":", numbers[i])
	}
}
