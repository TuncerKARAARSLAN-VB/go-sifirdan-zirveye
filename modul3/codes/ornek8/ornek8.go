package main

import "fmt"

// Person adında bir struct tanımlama
type Person struct {
	Name string
	Age  int
}

func main() {
	// Struct kullanarak bir nesne oluşturma
	person1 := Person{Name: "Alice", Age: 30}

	// Struct içindeki verilere erişim
	fmt.Println("İsim:", person1.Name)
	fmt.Println("Yaş:", person1.Age)
}
