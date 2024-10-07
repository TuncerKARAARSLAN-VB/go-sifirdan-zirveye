package main

import "fmt"

func main() {
	// Bir string-int map tanımlama ve başlatma
	ageMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 27,
	}

	// Map'in içeriğini yazdırma
	fmt.Println("Yaşlar:", ageMap)
}
