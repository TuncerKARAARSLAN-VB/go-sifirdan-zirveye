# Fonksiyon Geliştirme

Go dilinde fonksiyonlar basit ve etkili bir şekilde tanımlanıp kullanılabilir. İşte fonksiyonlarla ilgili temel konular:

## 1. Fonksiyon Tanımlama ve Kullanma

Go'da bir fonksiyon `func` anahtar kelimesi ile tanımlanır ve genellikle fonksiyon adını, parametre listesini ve dönüş değerini içerir.

```go
package main

import "fmt"

// Fonksiyon tanımlama
func sayHello() {
    fmt.Println("Hello, World!")
}

func main() {
    // Fonksiyon çağırma
    sayHello()
}
```

Bu örnekte, `sayHello` isimli bir fonksiyon tanımlanmıştır ve `main` fonksiyonundan çağrılmıştır. `fmt.Println` ile ekrana bir metin yazdırılmıştır.

## 2. Parametreler ve Dönüş Değerleri

Go'da fonksiyonlar parametreler alabilir ve bir veya daha fazla değer döndürebilir. Parametreler ve dönüş tipi fonksiyonun imzasında belirtilir.

**Tek parametre ve dönüş değeri:**

```go
package main
import "fmt"

// Parametre alan ve değer döndüren bir fonksiyon
func square(x int) int {
    return x * x
}

func main() {
    result := square(4)
    fmt.Println("4'ün karesi:", result)
}
```

Bu örnekte `square` fonksiyonu bir `int` parametre alır ve bir `int` değeri döner.

**Birden fazla parametre ile çalışmak:**

```go
package main
import "fmt"

// İki parametre alan ve toplam döndüren bir fonksiyon
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println("Toplam:", add(3, 7))
}
```

Bu örnekte `add` fonksiyonu iki `int` parametre alır ve toplama işlemi yapar.

## 3. Çoklu Dönüş Değeri Kullanımı

Go'da fonksiyonlar birden fazla değer döndürebilir. Bu, özellikle bir sonuçla birlikte hata veya ek bilgi döndürmek için çok kullanışlıdır.

```go
package main
import "fmt"

// İki değer döndüren bir fonksiyon
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("bölme hatası: sıfıra bölünemez")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Hata:", err)
    } else {
        fmt.Println("Sonuç:", result)
    }
}
```

Bu örnekte `divide` fonksiyonu iki değer döndürür: bölüm sonucu ve bir `error` (hata) nesnesi. Eğer `b == 0` olursa hata döndürülür, aksi takdirde bölüm sonucu döner.

## Özet

- Go'da fonksiyonlar `func` anahtar kelimesi ile tanımlanır.
- Fonksiyonlar parametre alabilir ve bir veya birden fazla değer döndürebilir.
- Çoklu dönüş değeri Go'nun güçlü özelliklerinden biridir ve hata işleme gibi durumlar için sıklıkla kullanılır.
