# Arayüz - Interface

Arayüzler, Go dilinde belirli bir davranışı tanımlayan metot imzalarını barındıran bir yapıdır. Bir arayüz, bir grup metotun imzalarını tanımlayarak, bu metotları uygulayacak olan türlerin (struct) belirli bir davranışı gerçekleştirmesini sağlar. Go'da arayüzler, esnekliği ve yeniden kullanılabilirliği artırmak için kullanılır.

## Arayüzlerin Go Dilindeki Yeri

- **Soyutlama**: Arayüzler, farklı türlerin benzer davranışlarını soyutlamaya yardımcı olur. Bu sayede, bir arayüzü uygulayan her tür, arayüzde tanımlı metotları gerçekleştirmek zorundadır.
- **Polimorfizm**: Arayüzler, farklı türlerin aynı arayüz üzerinden kullanılmasını sağlar. Bu, yazılımın daha esnek ve modüler olmasına yardımcı olur.

## Arayüz Tanımlama ve Uygulama

Bir arayüz tanımlamak için `type` anahtar kelimesi kullanılır. Aşağıda basit bir arayüz tanımı ve bu arayüzü uygulayan bir tür örneği verilmiştir:

## Örnek 1: Temel Arayüz Kullanımı

```go
package main

import (
	"fmt"
)

// Arayüz tanımı
type Animal interface {
	Speak() string
}

// Arayüzü uygulayan bir tür
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var animal Animal

	// Dog türünü kullanma
	animal = Dog{}
	fmt.Println(animal.Speak()) // Çıktı: Woof!

	// Cat türünü kullanma
	animal = Cat{}
	fmt.Println(animal.Speak()) // Çıktı: Meow!
}
```

## Açıklamalar:
- **Animal Arayüzü**: `Speak()` metodu tanımlar. Bu metodu uygulayan her tür, bu metodu içermek zorundadır.
- **Dog ve Cat**: Bu türler, `Animal` arayüzünü uygulayarak kendi `Speak` metotlarını sağlar.
- **main Fonksiyonu**: `Animal` tipinde bir değişken ile `Dog` ve `Cat` türleri üzerinde işlem yapar.

![Örnek 1 Çıktı](https://via.placeholder.com/300x100?text=Output%3A+Woof%21+Meow%21)

## Dinamik Tipler ve Type Assertion

Go dilinde dinamik tipler, bir arayüz değişkeni ile çalıştığınızda ortaya çıkar. Arayüzler, farklı türlerdeki verileri saklamak için kullanılabilir. **Type assertion**, bir arayüz değişkeninin belirli bir türde olup olmadığını kontrol etmenizi sağlar.

## Örnek 2: Type Assertion Kullanımı

```go
package main

import (
	"fmt"
)

// Arayüz tanımı
type Animal interface {
	Speak() string
}

// Arayüzü uygulayan bir tür
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	var animal Animal = Dog{}

	// Type assertion
	dog, ok := animal.(Dog)
	if ok {
		fmt.Println("Animal is a Dog:", dog.Speak())
	} else {
		fmt.Println("Animal is not a Dog")
	}
}
```

## Açıklamalar

- **Type Assertion**: `animal.(Dog)` ifadesi ile `animal` değişkeninin `Dog` türünde olup olmadığını kontrol ederiz. Eğer `animal` bir `Dog` ise, `ok` değişkeni `true` değerini alır ve ilgili metot çağrılır.

![Örnek 2 Çıktı](https://via.placeholder.com/300x100?text=Output%3A+Animal+is+a+Dog%3A+Woof%21)

# Çoklu Arayüzler ile Çoklu Kalıtım Davranışı

Go dilinde çoklu kalıtım doğrudan desteklenmez, ancak arayüzler (interfaces) kullanarak benzer bir davranış elde edilebilir. Aşağıda, bu kavramı daha iyi anlamak için birkaç örnek sunuyorum.

## Örnek 3: Çoklu Arayüz Uygulama

Bir hayvan sınıfının farklı davranışlarını temsil eden birkaç arayüz tanımlayalım. Bu arayüzler, hayvanların konuşma ve hareket etme yeteneklerini tanımlasın.

```go
package main

import (
	"fmt"
)

// Konuşma yeteneği
type Speaker interface {
	Speak() string
}

// Hareket etme yeteneği
type Mover interface {
	Move() string
}

// Hayvan arayüzü: Hem konuşma hem hareket etme yeteneğini içerir
type Animal interface {
	Speaker
	Mover
}

// Dog türü
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "The dog runs!"
}

// Cat türü
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "The cat jumps!"
}

func main() {
	var animal Animal

	// Dog türünü kullanma
	animal = Dog{}
	fmt.Println(animal.Speak()) // Çıktı: Woof!
	fmt.Println(animal.Move())  // Çıktı: The dog runs!

	// Cat türünü kullanma
	animal = Cat{}
	fmt.Println(animal.Speak()) // Çıktı: Meow!
	fmt.Println(animal.Move())  // Çıktı: The cat jumps!
}
```

## Açıklamalar

- **Animal Arayüzü**: `Speaker` ve `Mover` arayüzlerini birleştirir. Bu sayede `Dog` ve `Cat` türleri hem konuşabilir hem de hareket edebilir.
- **main Fonksiyonu**: `Animal` tipi kullanılarak her iki tür üzerinde de işlemler yapılır.

### Örnek 4: Çoklu Arayüzleri Birleştirme

Farklı arayüzleri bir araya getirerek daha karmaşık bir yapı oluşturalım.

```go
package main

import (
	"fmt"
)

// Uçabilen arayüz
type Flyer interface {
	Fly() string
}

// Su üzerinde hareket edebilen arayüz
type Swimmer interface {
	Swim() string
}

// Kuş ve balık arayüzü: Uçma ve yüzme yeteneklerini içerir
type BirdFish interface {
	Flyer
	Swimmer
}

// BirdFish türü
type Duck struct{}

func (d Duck) Fly() string {
	return "The duck flies!"
}

func (d Duck) Swim() string {
	return "The duck swims!"
}

func main() {
	var birdFish BirdFish

	// Duck türünü kullanma
	birdFish = Duck{}
	fmt.Println(birdFish.Fly())  // Çıktı: The duck flies!
	fmt.Println(birdFish.Swim()) // Çıktı: The duck swims!
}
```

## Açıklamalar

- **BirdFish Arayüzü**: Hem `Flyer` hem de `Swimmer` arayüzlerini içerir. `Duck` türü bu arayüzü uygulayarak hem uçma hem de yüzme yeteneğini kazanır.

### Örnek 5: Type Assertion ile Çoklu Kalıtım

Type assertion kullanarak bir arayüzün belirli bir türde olup olmadığını kontrol edebiliriz. Bu, farklı türlerin aynı işlevselliği nasıl sunduğunu görmek için yararlıdır.

```go
package main

import (
	"fmt"
)

// Arayüzler
type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Animal interface {
	Speaker
	Mover
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "The dog runs!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "The cat jumps!"
}

func main() {
	var animal Animal

	animal = Dog{}
	speakAndMove(animal)

	animal = Cat{}
	speakAndMove(animal)
}

// Type assertion ile farklı türleri kontrol etme
func speakAndMove(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.Move())

	// Type assertion ile Dog kontrolü
	if dog, ok := a.(Dog); ok {
		fmt.Println("This is a dog:", dog.Speak())
	}

	// Type assertion ile Cat kontrolü
	if cat, ok := a.(Cat); ok {
		fmt.Println("This is a cat:", cat.Speak())
	}
}
```

## Açıklamalar

- **speakAndMove Fonksiyonu**: `Animal` türünde bir argüman alır. `Type assertion