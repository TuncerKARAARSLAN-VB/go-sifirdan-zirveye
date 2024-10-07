### 1. Goroutine'ler

#### 1.1 Goroutine Kavramı ve Go’da Eş Zamanlı İşlemler
Goroutine, Go dilinde eş zamanlı programlamanın temel yapı taşıdır. Goroutine'ler, hafif iş parçacıkları (lightweight threads) olarak düşünülebilir. Go dilinde, goroutine'ler sayesinde programlar, aynı anda birden fazla işlemi yürütme yeteneğine sahiptir. Bu, programların daha verimli çalışmasını ve daha hızlı yanıt vermesini sağlar.

**Goroutine'lerin Avantajları:**
- **Hafiflik:** Geleneksel iş parçacıklarına göre daha az bellek kullanırlar.
- **Kolay Başlatma:** Başlatmak ve yönetmek oldukça basittir.
- **Görev Dağıtımı:** Görevler arasında daha iyi bir dağıtım sağlar.

#### 1.2 Goroutine’lerin Başlatılması ve Yönetimi
Goroutine'leri başlatmak için `go` anahtar kelimesi kullanılır. Bir fonksiyonun önüne `go` ekleyerek, o fonksiyonu yeni bir goroutine içinde çalıştırabilirsiniz.

**Örnek:**
```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Merhaba!")
}

func main() {
    go sayHello() // sayHello fonksiyonunu bir goroutine olarak başlat
    time.Sleep(time.Second) // Ana fonksiyonun hemen sonlanmasını önlemek için bekliyoruz
}
```

Bu örnekte, `sayHello` fonksiyonu bir goroutine olarak çalıştırılır. Ana fonksiyon, goroutine'in tamamlanmasını beklemek için bir saniye boyunca bekler.

#### 1.3 Goroutine Yönetimi
Goroutine'lerin yönetimi, Go'nun eş zamanlılık modeline dayanır. Go, goroutine'leri otomatik olarak yöneten bir çalışma zamanına (runtime) sahiptir. Bu çalışma zamanı, goroutine'lerin yönetiminde ve zamanlamasında etkilidir.

**Goroutine'lerin Durumu:**
- **Çalışma (Running):** Goroutine aktif olarak çalışıyor.
- **Bekleme (Waiting):** Goroutine, başka bir goroutine'in tamamlanmasını bekliyor.
- **Hazır (Ready):** Goroutine çalışmaya hazır, fakat CPU'da çalışmıyor.

**Goroutine'lerin İletişimi:**
Goroutine'ler arasında iletişim sağlamak için `channels` kullanılır. Channels, goroutine'ler arasında veri paylaşımını güvenli bir şekilde yönetir.

**Örnek:**
```go
package main

import (
    "fmt"
)

func sayHello(c chan string) {
    c <- "Merhaba!"
}

func main() {
    c := make(chan string) // Channel oluştur
    go sayHello(c)         // Goroutine'i başlat
    message := <-c         // Channel'den mesaj al
    fmt.Println(message)   // Mesajı yazdır
}
```

Bu örnekte, `sayHello` fonksiyonu bir mesajı channel üzerinden gönderir ve ana fonksiyon bu mesajı alır.

### Sonuç
Goroutine'ler, Go dilinde eş zamanlı programlama yapmanın temel yoludur. `go` anahtar kelimesi ile goroutine'ler başlatılır ve `channels` ile veri paylaşımı sağlanır. Bu sayede, yüksek performanslı ve verimli uygulamalar geliştirmek mümkün olur.

# Örneklerle Pekiştirme

### Örnek: Basit Toplama İşlemi

Bu örnekte, iki sayıyı toplayan bir fonksiyon tanımlayacağız ve sonucu bir channel aracılığıyla geri göndereceğiz.

```go
package main

import (
    "fmt"
)

// Toplama işlemini yapan fonksiyon
func topla(a int, b int, c chan int) {
    sonuc := a + b
    c <- sonuc // Sonucu channel üzerinden gönder
}

func main() {
    c := make(chan int) // Channel oluştur

    // Goroutine'leri başlat
    go topla(3, 5, c)
    go topla(10, 20, c)

    // Channel'den sonuçları al
    toplam1 := <-c
    toplam2 := <-c

    fmt.Printf("Birinci toplam: %d\n", toplam1)
    fmt.Printf("İkinci toplam: %d\n", toplam2)
}
```

### Çıktı:
```
Birinci toplam: 8
İkinci toplam: 30
```

Bu örnekte, `topla` fonksiyonu iki sayıyı toplar ve sonucu bir channel üzerinden gönderir. Ana fonksiyon, channel'den iki sonucu alır ve ekrana yazdırır.

### Örnek: Üretici-Tüketici Modeli

Bu örnekte, bir üretici goroutine bir kanal aracılığıyla verileri üretirken, bir tüketici goroutine bu verileri tüketir. 

```go
package main

import (
    "fmt"
    "time"
)

// Üretici fonksiyonu
func üretici(c chan int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Üretici: %d üretildi\n", i)
        c <- i // Channel'e veri gönder
        time.Sleep(time.Second) // Üretim arasında bekleme
    }
    close(c) // Channel'i kapat
}

// Tüketici fonksiyonu
func tüketici(c chan int) {
    for num := range c { // Channel'den gelen verileri oku
        fmt.Printf("Tüketici: %d tüketildi\n", num)
        time.Sleep(2 * time.Second) // Tüketim arasında bekleme
    }
}

func main() {
    c := make(chan int) // Channel oluştur

    go üretici(c)  // Üretici goroutine'i başlat
    go tüketici(c) // Tüketici goroutine'i başlat

    // Ana goroutine'in bitmesini bekle
    time.Sleep(12 * time.Second)
}
```

### Çıktı:
```
Üretici: 0 üretildi
Tüketici: 0 tüketildi
Üretici: 1 üretildi
Tüketici: 1 tüketildi
Üretici: 2 üretildi
Üretici: 3 üretildi
Tüketici: 2 tüketildi
Üretici: 4 üretildi
Tüketici: 3 tüketildi
Tüketici: 4 tüketildi
```

Bu örnekte, `üretici` fonksiyonu 0'dan 4'e kadar sayıları üretir ve channel'e gönderir. `tüketici` fonksiyonu ise bu sayıları alır ve ekrana yazdırır. Üretici, her üretim arasında bir saniye beklerken, tüketici her tüketim için iki saniye bekler. Channel kapatıldığında, `tüketici` fonksiyonu verileri okumayı durdurur.


Tabii ki! Asenkron çalışma, bir fonksiyonun çalışmasının diğer fonksiyonların çalışmasını etkilemeden gerçekleşmesini sağlar. Go dilinde goroutine kullanarak iki fonksiyonun asenkron olarak nasıl çalıştığını gösteren bir örnek geliştirelim.

### Örnek: Asenkron Görevler

Bu örnekte, iki asenkron görev tanımlayacağız. Birinci görev, belirli bir süre boyunca bekleyip bir mesaj yazdıracak, ikinci görev ise hemen bir mesaj yazdıracak. Böylece bu iki fonksiyonun birbirinden bağımsız olarak çalıştığını gözlemleyeceğiz.

```go
package main

import (
    "fmt"
    "time"
)

// İlk asenkron görev
func uzunGorev() {
    time.Sleep(3 * time.Second) // 3 saniye bekle
    fmt.Println("Uzun görev tamamlandı!")
}

// İkinci asenkron görev
func kisaGorev() {
    fmt.Println("Kısa görev hemen tamamlandı!")
}

func main() {
    // Uzun görevi goroutine olarak başlat
    go uzunGorev()

    // Kısa görevi hemen çalıştır
    kisaGorev()

    // Ana goroutine'in bitmesini beklemek için bir süre bekleyelim
    time.Sleep(4 * time.Second) // Uzun görev için yeterli bekleme süresi
}
```

### Çıktı:
```
Kısa görev hemen tamamlandı!
Uzun görev tamamlandı!
```

### Açıklama:
1. **`uzunGorev` Fonksiyonu:** Bu fonksiyon, 3 saniye bekledikten sonra "Uzun görev tamamlandı!" mesajını yazdırır.
2. **`kisaGorev` Fonksiyonu:** Bu fonksiyon hemen çalışır ve "Kısa görev hemen tamamlandı!" mesajını yazdırır.
3. **Asenkron Çalışma:** Ana fonksiyonda, `uzunGorev` fonksiyonu `go` anahtar kelimesi ile bir goroutine olarak başlatılır. Bu, `uzunGorev` fonksiyonunun ana goroutine'den bağımsız olarak çalışmasına olanak tanır.
4. **Sonuç:** `kisaGorev` fonksiyonu hemen çalışarak "Kısa görev hemen tamamlandı!" mesajını ekrana yazdırır. Daha sonra, ana fonksiyon 4 saniye bekler; bu süre zarfında `uzunGorev` fonksiyonu tamamlanır ve "Uzun görev tamamlandı!" mesajı yazdırılır.

Bu örnek, Go'da asenkron görevlerin nasıl çalıştığını ve goroutine'lerin bağımsız olarak nasıl yürütülebileceğini gösteriyor. 