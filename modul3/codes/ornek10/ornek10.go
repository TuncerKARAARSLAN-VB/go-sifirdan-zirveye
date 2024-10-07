package main

import "fmt"

// Person adında bir struct tanımlama
type Person struct {
	Name string
	Age  int
}

// Metot: Person struct'ı için bir metot tanımlama
func (p Person) Greet() {
	fmt.Printf("Merhaba, benim adım %s ve yaşım %d.\n", p.Name, p.Age)
}

func main() {
	// Struct kullanarak bir nesne oluşturma
	person1 := Person{Name: "Alice", Age: 30}

	// Metodu çağırma
	person1.Greet()
}
