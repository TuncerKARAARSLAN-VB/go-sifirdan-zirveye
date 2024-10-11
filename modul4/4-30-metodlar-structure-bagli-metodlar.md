# 1. Metotlar

Metotlar, bir nesnenin (struct’ın) davranışlarını tanımlayan işlevlerdir. Go dilinde metotlar, bir türün (struct veya arayüz) üzerine tanımlanabilir. Bu, o türden bir değerin metotlarını çağırabilmemizi sağlar. Metotlar, veri üzerinde işlem yapma yeteneği sağlar ve kodun yeniden kullanılabilirliğini artırır.

## 1.1 Structlara Bağlı Metotlar Tanımlama

Go dilinde struct'lar, bir veya daha fazla değeri bir arada tutan veri yapılarıdır. Struct'lar, bir nesne gibi davranabilen fakat değer tipi olan yapılardır. Struct'lar için metot tanımlamak mümkündür ve bu metotlar struct'ın davranışlarını belirler.

**Örnek: Struct ile Metot Tanımlama**

```go
package main

import (
	"fmt"
)

// Rectangle yapısı
type Rectangle struct {
	Width  float64
	Height float64
}

// Alan hesaplama metodu
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Çevre hesaplama metodu
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// Rectangle nesnesi oluşturma
	rect := Rectangle{Width: 5, Height: 10}
	
	// Alan ve çevre hesaplama
	fmt.Printf("Alan: %.2f\n", rect.Area())        // Çıktı: Alan: 50.00
	fmt.Printf("Çevre: %.2f\n", rect.Perimeter()) // Çıktı: Çevre: 30.00
}
```

## Açıklamalar

- **Rectangle Yapısı**: `Width` ve `Height` adında iki alan içerir.
- **Metot Tanımları**: `Area()` ve `Perimeter()` metotları, `Rectangle` yapısının alan ve çevresini hesaplar. Metotlar, `Rectangle` yapısının bir alıcı (`receiver`) olarak kullanılır.
- **Kullanım**: `main` fonksiyonu içinde bir `Rectangle` nesnesi oluşturulmuş ve alan ile çevre hesaplamaları yapılmıştır.

## 1.2 Pointer ve Değer Üzerinden Metotlar

Go’da metotlar, değer tipi ve referans tipi değişkenlerle çalışabilir. Değer tipi değişkenler (struct gibi) değerlerini kopyalayarak geçerken, referans tipi değişkenler (pointer gibi) bellekteki adresi geçirir.

### Değer Üzerinden Metotlar

Değer tipleri doğrudan kopyalanır. Bu, metot çağrısında orijinal değişkenin etkilenmemesi anlamına gelir.

**Örnek: Değer Üzerinden Metot**

```go
package main

import (
	"fmt"
)

type Counter struct {
	count int
}

// Değer üzerinden metot tanımlama
func (c Counter) Increment() {
	c.count++
}

func main() {
	counter := Counter{count: 0}
	counter.Increment()
	fmt.Println("Count after increment:", counter.count) // Çıktı: Count after increment: 0
}
```

### Açıklamalar

- **Counter Yapısı**: `count` adında bir alan içerir.
- **Değer Üzerinden Metot**: `Increment()` metodu, `count` değerini artırmaya çalışır. Ancak bu işlem yalnızca kopyalanmış bir değer üzerinde yapılır. Orijinal `counter` nesnesi etkilenmez.

### Pointer Üzerinden Metotlar

Pointer kullanıldığında, orijinal değişken üzerinde değişiklik yapılabilir.

**Örnek: Pointer Üzerinden Metot**

```go
package main

import (
	"fmt"
)

type Counter struct {
	count int
}

// Pointer üzerinden metot tanımlama
func (c *Counter) Increment() {
	c.count++ // Orijinal nesne üzerinde değişiklik yapılıyor
}

func main() {
	counter := Counter{count: 0}
	counter.Increment() // Pointer üzerinden çağrılıyor
	fmt.Println("Count after increment:", counter.count) // Çıktı: Count after increment: 1
}
```

### Açıklamalar

- **Pointer Alıcı**: `Increment()` metodu, `*Counter` türünde bir alıcı alır. Bu, orijinal `Counter` nesnesinin değiştirilebileceği anlamına gelir.
- **Kullanım**: `Increment()` metodu çağrıldığında, orijinal `counter` nesnesinin `count` değeri 1 artar.

### Özet

1. **Struct'lar**: Verileri gruplamak için kullanılır ve metotlar tanımlamak mümkündür.
2. **Değer Üzerinden Metotlar**: Değer tipleri metotlara geçerken kopyalanır; bu, orijinal değişkenin etkilenmediği anlamına gelir.
3. **Pointer Üzerinden Metotlar**: Pointer kullanıldığında, orijinal değişken üzerinde değişiklik yapılabilir.

## Arayüzler ve Metot Kullanımı

Arayüzler, Go dilinde polimorfizmi sağlamak için kullanılır. Bir arayüz, bir veya daha fazla metodu tanımlar ve bu metotları uygulayan herhangi bir tür, bu arayüzü kullanabilir. Bu sayede, farklı türlerdeki nesneleri tek bir arayüz altında toplayabiliriz.

**Örnek: Arayüz ile Metot Kullanımı**

```go
package main

import (
	"fmt"
)

// Shape arayüzü
type Shape interface {
	Area() float64
}

// Rectangle yapısı
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle yapısı
type Circle struct {
	Radius float64
}

// Rectangle için Area metodu
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle için Area metodu
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Alanı hesaplayan fonksiyon
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	// Farklı türdeki şekiller
	rect := Rectangle{Width: 5, Height: 10}
	circ := Circle{Radius: 3}

	// Alanları yazdırma
	printArea(rect) // Çıktı: Area: 50.00
	printArea(circ) // Çıktı: Area: 28.26
}
```

### Açıklamalar:

- **Shape Arayüzü**: `Area()` metodunu tanımlar.
- **Rectangle ve Circle Yapıları**: Bu yapılar, `Shape` arayüzünü uygular ve kendi `Area` metotlarını sağlar.
- **printArea Fonksiyonu**: `Shape` arayüzünü alır ve ilgili alanı yazdırır.

## Sonuç

Bu modül, Go dilinde metotlar, yapılar (struct) ve pointer’lar hakkında temel bilgileri kapsamaktadır. 

- Metotlar, yapılarla birlikte kullanılarak belirli davranışları tanımlar.
- Değer ve pointer üzerinden metotlar kullanarak, orijinal değişkenin üzerinde değişiklik yapma yeteneği sağlanır.
- Arayüzler, farklı türleri bir araya getirerek polimorfizmi destekler.

Bu konular, Go dilinde etkili ve esnek kod yazmanıza yardımcı olacak temel kavramlardır. Eğer daha fazla detay veya belirli bir konuda örnek isterseniz, lütfen belirtin!