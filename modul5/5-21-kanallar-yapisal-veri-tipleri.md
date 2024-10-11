# Kanallar ve Yapısal Veri Tipleri

Bir goroutine içinden yapı (struct) veya dizi (array) gibi karmaşık veri türlerini de kanala gönderebilirsiniz. Go dilinde kanallar, herhangi bir türdeki verileri iletmek için kullanılabilir; bu türler arasında kullanıcı tanımlı yapılar (struct) ve diziler de bulunmaktadır.

## Örnek 1: Struct Gönderme

Aşağıda, bir struct tanımlayıp bu struct'ı bir goroutine içinden bir kanala gönderen bir örnek bulabilirsiniz:

```go
package main

import (
    "fmt"
)

// Kullanıcı tanımlı struct
type Kişi struct {
    İsim string
    Yaş  int
}

// Üretici fonksiyonu, kanala Kişi struct'larını gönderir
func üretici(c chan Kişi) {
    kişiler := []Kişi{
        {"Ali", 30},
        {"Ayşe", 25},
        {"Mehmet", 40},
    }

    for _, kişi := range kişiler {
        c <- kişi // Kanal üzerinden Kişi struct'ını gönderiyoruz
    }
    close(c) // Kanalı kapatıyoruz
}

func main() {
    c := make(chan Kişi) // Kişi türünde bir kanal oluşturuyoruz

    go üretici(c) // Üretici goroutine'i başlatıyoruz

    // Tüketici, kanaldan gelen Kişi struct'larını okur
    for kişi := range c { // Kanaldan gelen Kişi struct'larını okuyoruz
        fmt.Printf("Alınan Kişi: İsim: %s, Yaş: %d\n", kişi.İsim, kişi.Yaş) // Kişi bilgilerini yazdırıyoruz
    }
}
```

### Çıktı

Bu kod çalıştırıldığında aşağıdaki çıktıyı verir:

```
Alınan Kişi: İsim: Ali, Yaş: 30
Alınan Kişi: İsim: Ayşe, Yaş: 25
Alınan Kişi: İsim: Mehmet, Yaş: 40
```

## Örnek 2: Dizi Gönderme

Aşağıda, bir dizi gönderip bu diziyi bir goroutine içinden bir kanala gönderen başka bir örnek bulunmaktadır:

```go
package main

import (
    "fmt"
)

// Üretici fonksiyonu, kanala int dizilerini gönderir
func üretici(c chan []int) {
    diziler := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    for _, dizi := range diziler {
        c <- dizi // Kanal üzerinden dizi gönderiyoruz
    }
    close(c) // Kanalı kapatıyoruz
}

func main() {
    c := make(chan []int) // int dizisi türünde bir kanal oluşturuyoruz

    go üretici(c) // Üretici goroutine'i başlatıyoruz

    // Tüketici, kanaldan gelen int dizilerini okur
    for dizi := range c { // Kanaldan gelen int dizilerini okuyoruz
        fmt.Println("Alınan Dizi:", dizi) // Diziyi yazdırıyoruz
    }
}
```

### Çıktı

Bu kod çalıştırıldığında aşağıdaki çıktıyı verir:

```
Alınan Dizi: [1 2 3]
Alınan Dizi: [4 5 6]
Alınan Dizi: [7 8 9]
```

## Açıklamalar

1. **Struct Gönderme**: İlk örnekte, `Kişi` adında bir struct tanımladık ve bu struct'ları bir kanala gönderiyoruz.
2. **Dizi Gönderme**: İkinci örnekte, int dizilerini bir kanala gönderen bir goroutine oluşturduk.
3. **Kanal Kapatma**: Her iki örnekte de, tüm veriler gönderildikten sonra kanalı kapatıyoruz. Bu, tüketici goroutine'ine daha fazla veri gelmeyeceğini belirtir.

Bu tür verilerin kanallarda güvenli bir şekilde iletilmesi, eş zamanlı programlamanın önemli bir parçasıdır. Eğer başka sorularınız veya daha fazla örnek ihtiyacınız varsa, lütfen belirtin!