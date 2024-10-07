Go dilindeki kontrol yapıları, program akışını yönetmek ve koşullara göre farklı yollar izlemek için kullanılır. 

### 1. `if` ve `else` Yapıları
- **Temel Kullanım**: `if` yapısı, belirli bir koşulun doğru olup olmadığını kontrol eder. Eğer koşul doğruysa, ilgili kod bloğu çalıştırılır.
- **`else` Kullanımı**: `else`, `if` koşulu yanlışsa alternatif bir yol sağlar.

**Örnek**:
```go
package main

import (
    "fmt"
)

func main() {
    x := 10

    if x < 0 {
        fmt.Println("Negatif")
    } else {
        fmt.Println("Pozitif veya sıfır")
    }
}
```

### 2. `if` ile Kısa Değişken Tanımlama
Go’da `if` ifadesinde koşul tanımlarken kısa değişken tanımlaması yapılabilir. Bu, koşul ifadesinin başlangıcında tanımlanan bir değişkeni içerir.

**Örnek**:
```go
if y := 20; y < 15 {
    fmt.Println("Y değeri 15'ten küçük")
} else {
    fmt.Println("Y değeri 15 veya daha büyük")
}
```

### 3. `switch` Yapısı
`switch`, bir değişkenin değerine göre farklı yollar izlemeyi sağlar. Birden fazla `case` ifadesi bulunabilir ve `switch` içindeki herhangi bir case koşulu doğru olduğunda, o case’e ait kod bloğu çalıştırılır.

**Örnek**:
```go
package main

import (
    "fmt"
)

func main() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Pazartesi")
    case 2:
        fmt.Println("Salı")
    case 3:
        fmt.Println("Çarşamba")
    default:
        fmt.Println("Geçersiz gün")
    }
}
```

### 4. `switch`-`case` Yapısındaki Farklılıklar
Go'daki `switch` yapısı, bazı özellikleriyle diğer dillerden ayrılır:

- **Varsayılan Durum**: `default` case, hiçbir case koşulu doğru değilse çalıştırılır.
- **Düşey (Fall-through)**: Go'da, `case` ifadeleri arasında varsayılan bir geçiş yoktur. Eğer bir `case` doğruysa, o `case`'in altındaki kod çalıştırılır ve diğer `case`'lere geçilmez. Ancak, `fallthrough` anahtar kelimesi ile bir `case`'den sonra gelen `case`'in de çalıştırılması sağlanabilir.

**Örnek**:
```go
switch x := 2; x {
case 1:
    fmt.Println("Bir")
case 2:
    fmt.Println("İki")
    fallthrough // bu durumda "Üç" de yazdırılır
case 3:
    fmt.Println("Üç")
default:
    fmt.Println("Geçersiz")
}
```

### 5. `for` Döngüsü
Go, döngü yapıları için sadece `for` döngüsünü kullanır. `for`, üç bileşenle tanımlanabilir: başlangıç ifadesi, koşul ifadesi ve artış ifadesi.

**Örnek**:
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

### 6. Sonsuz Döngü
Eğer koşul ifadesi belirtilmezse, `for` döngüsü sonsuz bir döngü haline gelir.

**Örnek**:
```go
package main

import "fmt"

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        if i > 5 {
            break
        }
    }
}
```

Bu kontrol yapıları, Go dilinde program akışını yönetmek ve koşullara göre farklı yollar izlemek için kullanılır. Programcılar bu yapıları kullanarak, belirli koşullara dayalı karar verme süreçlerini etkin bir şekilde gerçekleştirebilirler.