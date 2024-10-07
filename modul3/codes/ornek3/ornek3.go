package main

import "fmt"

func main() {
	// Bir slice tanımlama ve başlatma
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Dilim:", numbers)
	fmt.Println("Uzunluk:", len(numbers))  // 5
	fmt.Println("Kapasite:", cap(numbers)) // 5

	// Dilimi genişletme (append ile eleman ekleme)
	numbers = append(numbers, 6, 7)
	fmt.Println("Yeni dilim:", numbers)
	fmt.Println("Yeni uzunluk:", len(numbers))  // 7
	fmt.Println("Yeni kapasite:", cap(numbers)) // Genişleyebilir
}
