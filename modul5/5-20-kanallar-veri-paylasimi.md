# 1. Kanallar (Channels)

## 1.1 Kanal Tanımı ve Kullanımı
Kanallar, Go dilinde goroutine'ler arasında veri iletimi sağlamak için kullanılan bir yapıdır. Bir kanal, bir türdeki verileri (örneğin, `int`, `string`) güvenli bir şekilde iletmek için tasarlanmıştır. Kanallar, `make` fonksiyonu kullanılarak oluşturulur ve `chan` anahtar kelimesi ile tanımlanır.

**Kanal Oluşturma:**
```go
c := make(chan int) // int türünde bir kanal oluştur
```

**Kanal Kullanımı:**
Kanal üzerinden veri göndermek için `<-` operatörü kullanılır:
```go
c <- 42 // Kanala 42 değerini gönder
```

Veri almak için de aynı operatör kullanılır:
```go
value := <-c // Kanaldan bir değer al
```

**Örnek:**
```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int) // int türünde bir kanal oluştur

    // Yeni bir goroutine başlatıyoruz
    go func() {
        c <- 42 // Kanala 42 değerini gönderiyoruz
    }()

    value := <-c // Kanaldan veri alıyoruz
    fmt.Println(value) // 42 yazdırıyoruz
}
```

## 1.2 Eş Zamanlı İşlemler Arası Veri Paylaşımı
Kanallar, eş zamanlı işlemler arasında veri paylaşımını güvenli bir şekilde yönetir. Goroutine'ler, kanalları kullanarak verileri paylaşabilir ve bu sayede veri tutarlılığını sağlamak mümkündür.

**Örnek:**
```go
package main

import (
    "fmt"
)

// Üretici fonksiyonu, kanala verileri gönderir
func üretici(c chan int) {
    for i := 0; i < 5; i++ {
        c <- i // Kanal üzerinden veriyi gönderiyoruz
    }
    close(c) // Kanalı kapatıyoruz
}

// Tüketici fonksiyonu, kanaldan gelen verileri okur
func tüketici(c chan int) {
    for num := range c { // Kanaldan gelen verileri okuyoruz
        fmt.Println("Tüketici:", num) // Tüketilen veriyi yazdırıyoruz
    }
}

func main() {
    c := make(chan int) // int türünde bir kanal oluşturuyoruz

    go üretici(c)  // Üretici goroutine'i başlatıyoruz
    tüketici(c)    // Tüketici goroutine'i çalıştırıyoruz
}
```

Bu örnekte, bir üretici goroutine kanala 0'dan 4'e kadar sayıları gönderirken, tüketici goroutine bu sayıları alır ve ekrana yazdırır.

## 1.3 `select` Yapısı ile Kanalları Yönetme
`select` yapısı, birden fazla kanaldan veri almayı yönetmek için kullanılır. `select`, bir veya daha fazla kanaldan gelen verilerin durumuna göre işlem yapar. Eğer bir kanal hazırsa, ilgili case çalışır.

**Örnek:**
```go
package main

import (
    "fmt"
    "time"
)

// kanaliGonder fonksiyonu, belirtilen kanala bir mesaj gönderir
func kanaliGonder(c chan string, mesaj string) {
    time.Sleep(1 * time.Second) // Simülasyon için bekleme
    c <- mesaj // Kanala mesaj gönderiyoruz
}

func main() {
    c1 := make(chan string) // İlk kanal
    c2 := make(chan string) // İkinci kanal

    // İki goroutine başlatıyoruz
    go kanaliGonder(c1, "Mesaj 1")
    go kanaliGonder(c2, "Mesaj 2")

    // select ile kanalları yönetiyoruz
    for i := 0; i < 2; i++ {
        select {
        case mesaj1 := <-c1: // c1 kanalından mesaj al
            fmt.Println("Alınan:", mesaj1) // Mesajı yazdır
        case mesaj2 := <-c2: // c2 kanalından mesaj al
            fmt.Println("Alınan:", mesaj2) // Mesajı yazdır
        }
    }
}
```

Bu örnekte, `kanaliGonder` fonksiyonu iki kanala mesaj gönderir. Ana fonksiyonda `select` yapısı ile her iki kanaldan gelen mesajlar dinlenir ve alınan mesaj ekrana yazdırılır.

## 1.4 Kanallarda Bloklama ve Asenkron İletişim
Kanallar, veri gönderildiğinde veya alındığında bloklama (blocking) yapar. Yani, bir goroutine, kanal üzerinden veri gönderdiğinde, alıcı goroutine bu veriyi alana kadar durur. Bu durum, senkronizasyon sağlamak için yararlıdır.

**Bloklama Örneği:**
```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int) // int türünde bir kanal oluştur

    // Goroutine
    go func() {
        c <- 10 // Kanal üzerinden veri gönderir, alıcı yoksa bloklanır
    }()

    value := <-c // Kanaldan veri alır
    fmt.Println("Alınan Değer:", value) // Alınan değeri yazdır
}
```

Bu örnekte, goroutine kanala veri gönderdiğinde, ana goroutine bu veriyi alana kadar bekler.

**Asenkron İletişim:**
Asenkron iletişim, bir goroutine’in diğerine bağımlı olmadan çalışabilmesini sağlar. Kanal kapatılmadan önce her iki tarafın da veri gönderip alabilmesi için yeterli kontrol sağlanmalıdır.

**Asenkron İletişim Örneği:**
```go
package main

import (
    "fmt"
    "time"
)

// asenkronGorev fonksiyonu, bir süre bekledikten sonra mesaj gönderir
func asenkronGorev(c chan string) {
    time.Sleep(2 * time.Second) // 2 saniye bekle
    c <- "Görev tamamlandı!" // Kanala mesaj gönder
}

func main() {
    c := make(chan string) // string türünde bir kanal oluştur

    go asenkronGorev(c) // Asenkron görev başlat

    fmt.Println("Ana görev devam ediyor...") // Ana görev devam ediyor mesajını yazdır
    mesaj := <-c // Kanaldan mesaj al
    fmt.Println(mesaj) // "Görev tamamlandı!" yazdır
}
```

Bu örnekte, ana görev, asenkron görevin tamamlanmasını beklerken devam eder. Kanal üzerinden mesaj geldiğinde, görev tamamlandığını belirtir.

## Sonuç
Kanallar, Go'da eş zamanlı programlama için güçlü bir araçtır. Verilerin güvenli bir şekilde paylaşımını sağlarken, `select` yapısı ile çoklu kanalları yönetmek mümkündür. Bloklama ve asenkron iletişim özellikleri ile goroutine'ler arasında etkili bir senkronizasyon sağlanır. Bu özellikler, yüksek performanslı ve verimli uygulamalar geliştirmek için büyük önem taşır.
