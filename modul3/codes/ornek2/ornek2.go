package main

import "fmt"

func main() {
	// Bir dilim (slice) tanımlama ve başlatma
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Asal sayılar:", primes)

	// Bir diziden dilim oluşturma
	var numbers = [6]int{10, 20, 30, 40, 50, 60}
	var sliceNumbers []int = numbers[1:4] // 1. indisten 4. indise kadar (4 hariç)
	fmt.Println("Diziden alınan slice:", sliceNumbers)
}
