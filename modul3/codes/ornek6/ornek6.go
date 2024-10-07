package main

import "fmt"

func main() {
	ageMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 27,
	}

	// "Bob" anahtarını map'ten silme
	delete(ageMap, "Bob")
	fmt.Println("Güncellenmiş yaşlar:", ageMap)
}
