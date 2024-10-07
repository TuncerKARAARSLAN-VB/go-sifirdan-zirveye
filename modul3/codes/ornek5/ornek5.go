package main

import "fmt"

func main() {
	ageMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// Yeni anahtar-değer çifti ekleme
	ageMap["Carol"] = 27
	fmt.Println("Güncellenmiş yaşlar:", ageMap)
}
