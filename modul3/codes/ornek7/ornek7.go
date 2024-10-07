package main

import "fmt"

func main() {
	ageMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// "Carol" anahtarını arama
	value, exists := ageMap["Carol"]
	if exists {
		fmt.Println("Carol'ın yaşı:", value)
	} else {
		fmt.Println("Carol map'te yok.")
	}

	// "Carol" anahtarını arama
	value2, exists2 := ageMap["Bob"]
	if exists2 {
		fmt.Println("Bob'un yaşı:", value2)
	} else {
		fmt.Println("Bob map'te yok.")
	}
}
