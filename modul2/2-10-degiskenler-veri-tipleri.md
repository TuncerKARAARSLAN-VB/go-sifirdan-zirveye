# Değişkenler ve Veri Tipleri

Go dilinde değişkenler ve veri tipleri, programın temel yapı taşlarındandır. 

## 1. Değişken Tanımlama

Go dilinde değişken tanımlamanın birkaç yolu vardır:

### a. `var` Anahtar Kelimesi ile Değişken Tanımlama

`var` anahtar kelimesi kullanılarak değişken tanımlanabilir. Bu yöntemle değişkenin tipi açıkça belirtilmelidir.

```go
var x int // x adında bir tamsayı değişkeni tanımlanıyor
var y float64 = 3.14 // y adında bir float64 değişkeni tanımlanıyor ve başlangıç değeri atanıyor
var name string = "Go Dili" // name adında bir string değişkeni tanımlanıyor
```

### b. Kısa Tanımlama (:=)

Kısa tanımlama operatörü (`:=`) kullanılarak değişkenler tanımlanabilir. Bu yöntem, yalnızca yerel değişkenler için geçerlidir ve tip çıkarımı yapılır.

```go
x := 10 // x adında bir tamsayı değişkeni tanımlanıyor ve 10 değeri atanıyor
y := 3.14 // y adında bir float64 değişkeni tanımlanıyor
name := "Go Dili" // name adında bir string değişkeni tanımlanıyor
```

## 2. Veri Tipleri

Go, çeşitli veri tipleri sunar. Bunlar arasında temel veri tipleri ve yapısal veri tipleri bulunmaktadır.

### a. Temel Veri Tipleri

1. **Tamsayılar (Integer Types)**:
   - `int` (platforma bağlı olarak 32 veya 64 bit)
   - `int8` (8 bit tamsayı)
   - `int16` (16 bit tamsayı)
   - `int32` (32 bit tamsayı)
   - `int64` (64 bit tamsayı)
   - `uint` (platforma bağlı olarak 32 veya 64 bit işaretsiz tamsayı)
   - `uint8`, `uint16`, `uint32`, `uint64` (işaretsiz tamsayılar)

2. **Ondalık Sayılar (Floating Point Types)**:
   - `float32` (32 bit ondalıklı sayı)
   - `float64` (64 bit ondalıklı sayı)

3. **Boole (Boolean)**:
   - `bool` (true veya false değerlerini alabilir)

4. **Karakter Dizileri (String)**:
   - `string` (metinleri temsil eder)

### b. Yapısal Veri Tipleri

1. **Diziler (Arrays)**:
   - Sabit boyutlu veri yapılarıdır. Örneğin: `var a [5]int` (5 elemanlı bir tamsayı dizisi).

2. **Kesitler (Slices)**:
   - Dinamik boyutlu dizilerdir. Örneğin: `var b []int` (tamsayı kesiti).

3. **Haritalar (Maps)**:
   - Anahtar-değer çiftlerini depolamak için kullanılır. Örneğin: `var m map[string]int` (string anahtarları ve tamsayı değerleri içeren bir harita).

4. **Yapılar (Structs)**:
   - Farklı veri tiplerini bir araya getiren özelleştirilmiş veri tipleridir. Örneğin:

   ```go
   type Person struct {
       Name string
       Age  int
   }
   ```

5. **Arayüzler (Interfaces)**:
   - Bir nesnenin bir veya daha fazla yöntemini tanımlayan türlerdir.

## Temel Veri Tipleri Örnek Kullanımları

Aşağıda, değişkenlerin ve veri tiplerinin kullanıldığı basit bir örnek verilmiştir:

[Değişkenler Code](codes/degiskenler/degiskenler.go)

```go
package main

import "fmt"

func main() {
    // Değişken tanımlama
    var age int = 30
    var height float64 = 1.75
    var isStudent bool = false
    name := "Ali"

    // Değerleri yazdırma
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
}
```

![Değişkenler kod ve çıktısı](images/degiskenler.png)

Bu örnekte, farklı veri tiplerinde değişkenler tanımlanmış ve bu değişkenlerin değerleri `fmt.Println` ile yazdırılmıştır. Go dilinde değişkenler ve veri tipleri, uygulama geliştirme sürecinde önemli bir rol oynar ve doğru kullanımları yazılımın başarısını etkileyebilir.

## Yapısal Kullanım Örnekleri

Go dilinde yapısal veri türleri, birden fazla veriyi bir araya getirmek ve birlikte kullanmak için yapılandırılmış veri tipleridir. En yaygın kullanılan yapısal veri türlerinden biri **struct**'lardır. Struct, birden fazla farklı veri türünü bir araya getirip tek bir veri tipi olarak kullanmanızı sağlar.

Aşağıda yapısal veri türü olan `struct` kullanılarak verilen örneğin geliştirilmiş bir versiyonu bulunmaktadır:

### Örnek: Struct ile Yapısal Veri Türü Kullanımı

```go
package main

import "fmt"

// Bir struct tanımlama
type Person struct {
    Name      string
    Age       int
    Height    float64
    IsStudent bool
}

func main() {
    // Person türünden bir değişken oluşturma ve değer atama
    person := Person{
        Name:      "Ali",
        Age:       30,
        Height:    1.75,
        IsStudent: false,
    }

    // Değerleri yazdırma
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("Height:", person.Height)
    fmt.Println("Is Student:", person.IsStudent)
}
```

## Açıklama

- **`Person` struct'ı**: `Name`, `Age`, `Height` ve `IsStudent` gibi farklı veri türlerine sahip alanları bir araya getirir. Bu alanlar bir kişinin bilgilerini temsil eder.
- **Struct örneği oluşturma**: `person` adında bir değişken `Person` struct'ı ile oluşturulmuş ve tüm alanlarına değer atanmıştır.
- **Değerleri yazdırma**: `person` değişkeninin her alanına (`Name`, `Age`, `Height`, `IsStudent`) erişilerek ekrana yazdırılmıştır.

Bu yapıda, bir kişinin bilgilerini tek bir yapı içinde gruplamış olduk ve veriler arasında daha anlamlı bir ilişki kurduk. `struct` kullanarak kodun daha düzenli ve anlaşılır olmasını sağladık.

Go dilinde, temel veri tiplerini kullanarak ayrı ayrı örnekler aşağıda verilmiştir. Bu örneklerde **integer (int)**, **float (float64)**, **boolean (bool)** ve **string** veri tipleri ayrı ayrı tanımlanıp kullanılmıştır.

Go dilinde diziler, kesitler (slices) ve haritalar (maps) temel veri yapılarıdır. Aşağıda bu veri yapılarını ayrı ayrı açıklayan örnekler verilmiştir:

## 1. Dizi (Array) Örneği

Diziler, sabit boyutlu veri yapılarıdır ve tanımlandıktan sonra boyutları değiştirilemez.

```go
package main

import "fmt"

func main() {
    // 5 elemanlı bir integer dizisi tanımlama
    var numbers [5]int = [5]int{10, 20, 30, 40, 50}
    
    // Dizinin elemanlarına erişim
    fmt.Println("Numbers array:", numbers)
    fmt.Println("First element:", numbers[0]) // İlk elemanı yazdırma
    fmt.Println("Length of array:", len(numbers)) // Dizinin uzunluğunu yazdırma
}
```

## 2. Kesit (Slice) Örneği

Kesitler (slices), dizilerden farklı olarak dinamik boyutlu veri yapılarıdır. Boyutları değiştirilebilir ve dizilerden kesit almak için kullanılır.

```go
package main

import "fmt"

func main() {
    // Kesit tanımlama
    var numbers []int = []int{10, 20, 30, 40, 50}
    
    // Kesitin elemanlarına erişim
    fmt.Println("Numbers slice:", numbers)
    fmt.Println("First element:", numbers[0]) // İlk elemanı yazdırma
    fmt.Println("Length of slice:", len(numbers)) // Kesitin uzunluğu

    // Kesite yeni eleman ekleme
    numbers = append(numbers, 60, 70)
    fmt.Println("Slice after append:", numbers) // Yeni elemanlar eklendikten sonra
}
```

## 3. Harita (Map) Örneği

Haritalar, anahtar-değer (key-value) çiftleriyle çalışır. Anahtarlar benzersizdir ve her anahtar bir değere karşılık gelir.

```go
package main

import "fmt"

func main() {
    // Harita tanımlama: string anahtarlar ve integer değerler
    var studentGrades map[string]int = map[string]int{
        "Ali":  85,
        "Ayşe": 90,
        "Mehmet": 78,
    }

    // Haritanın elemanlarına erişim
    fmt.Println("Student Grades:", studentGrades)
    fmt.Println("Ali's Grade:", studentGrades["Ali"]) // Ali'nin notunu yazdırma
    
    // Yeni bir anahtar-değer çifti ekleme
    studentGrades["Fatma"] = 88
    fmt.Println("Updated Grades:", studentGrades)

    // Bir anahtarın var olup olmadığını kontrol etme
    grade, exists := studentGrades["Veli"]
    if exists {
        fmt.Println("Veli's Grade:", grade)
    } else {
        fmt.Println("Veli not listesinde bulunamadı.")
    }
}
```

## Açıklamalar

- **Diziler (Arrays)**: Sabit uzunlukta olup, her elemanın aynı veri tipinde olması gerekir. Boyutları sabittir.
- **Kesitler (Slices)**: Dinamik boyutlu diziler gibi çalışırlar ve boyutları esnek bir şekilde değiştirilebilir. `append` fonksiyonu ile yeni elemanlar eklenebilir.
- **Haritalar (Maps)**: Anahtar-değer çifti olarak veri saklar. Anahtarlar benzersizdir ve bir değere karşılık gelir. Anahtar ve değer tipleri farklı olabilir.
