package main

import "fmt"

func main() {
	// Boyutu sabit olan bir int dizisi tanımlama
	var numbers [5]int

	// Dizinin elemanlarına değer atama
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Dizi:", numbers)

	// Diziyi doğrudan tanımlama ve başlatma
	colors := [3]string{"Kırmızı", "Yeşil", "Mavi"}
	fmt.Println("Renkler:", colors)
}
