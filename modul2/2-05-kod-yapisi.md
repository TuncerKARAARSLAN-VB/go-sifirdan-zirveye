### 1. `package main`

`package` anahtar kelimesi, Go dilinde bir dosyanın hangi pakete ait olduğunu belirtir. Go programları paketler (modules) halinde düzenlenir ve her Go dosyası bir paketin parçasıdır.

- **`package main`**, Go'da özel bir anlam taşır. `main` paketi, Go uygulamasının başlangıç noktasıdır. Go programında yürütülecek fonksiyon `main` paketinde bulunmalıdır. Diğer bir deyişle, Go programının çalışmaya başlayacağı yer `package main` içindeki `main` fonksiyonudur.
- **Ana uygulama dosyaları**: Eğer `package main` yerine başka bir paket adı kullanırsanız, o dosya bir kütüphane veya yardımcı kod parçası olur, ancak doğrudan çalıştırılamaz.

#### Örnek:
```go
package main

func main() {
    // Bu fonksiyon programın başladığı yerdir.
    println("Program başlıyor...")
}
```

Bu örnekte `package main` kullanılmış ve bu dosyanın bir uygulama dosyası olduğu belirtilmiştir.

### 2. `import "fmt"`

`import` anahtar kelimesi, Go'da başka paketleri (kütüphaneleri) kodunuza dahil etmek için kullanılır. Bir paket içindeki fonksiyonları, tipleri veya diğer yapıları kullanmak istiyorsanız o paketi içe aktarmanız gerekir.

- **`fmt` paketi**, Go'nun standart kütüphanesinde bulunan bir formattır ve girdileri yazdırma (`print`), girdiyi okuma (`scan`) gibi temel işlevleri içerir. Adı "format" kelimesinden gelir ve sıklıkla kullanılır.
- `fmt` paketindeki en yaygın kullanılan fonksiyonlar `fmt.Println`, `fmt.Printf`, `fmt.Sprint` gibi çıktı yazdırma fonksiyonlarıdır.

#### Örnek:
```go
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, Dünya!") // Ekrana yazdırır
}
```

Bu örnekte, `import "fmt"` ile `fmt` paketi içe aktarılmıştır ve `fmt.Println` kullanılarak ekrana bir metin yazdırılmıştır.

### `import` İle Birden Fazla Paket Kullanımı
Go'da birden fazla paketi aynı anda import edebilirsiniz. Bunun için her paketi ayrı satırlarda yazabileceğiniz gibi, tek bir `import` bloğu içerisinde toplu halde de yazabilirsiniz.

#### Örnek:
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Pi sayısı:", math.Pi)
}
```

Bu örnekte hem `fmt` hem de `math` paketleri içe aktarılmıştır. `math.Pi` kullanılarak matematiksel sabit olan Pi sayısı ekrana yazdırılmıştır.

# fmt Fonksiyonları: Fonksiyon Listesi

Go dilindeki `fmt` paketi, girdileri formatlama, çıktı yazdırma ve okuma gibi işlevler için çeşitli fonksiyonlar sunar. Aşağıda `fmt` paketinde yaygın olarak kullanılan bazı fonksiyonları listeleyen bir tablo bulunuyor:

| Fonksiyon          | Açıklama                                                                 |
|--------------------|--------------------------------------------------------------------------|
| **`fmt.Print`**     | Argümanları yazdırır, yeni satıra geçmez.                                |
| **`fmt.Println`**   | Argümanları yazdırır ve yeni satıra geçer.                               |
| **`fmt.Printf`**    | Argümanları belirtilen biçimde (formatlı) yazdırır.                      |
| **`fmt.Sprintf`**   | Argümanları belirtilen biçimde bir string olarak döner.                  |
| **`fmt.Sprint`**    | Argümanları bir string olarak döner.                                     |
| **`fmt.Sprintln`**  | Argümanları bir string olarak döner ve yeni satıra geçer.                |
| **`fmt.Fprint`**    | Argümanları belirtilen `io.Writer` (örneğin bir dosya) üzerine yazar.    |
| **`fmt.Fprintln`**  | Argümanları belirtilen `io.Writer` üzerine yazar ve yeni satıra geçer.   |
| **`fmt.Fprintf`**   | Argümanları belirtilen `io.Writer` üzerine formatlı yazar.               |
| **`fmt.Scan`**      | Kullanıcıdan girdi alır, boşlukla ayrılan değerleri okur.                |
| **`fmt.Scanln`**    | Kullanıcıdan girdi alır, yeni satıra kadar olan değerleri okur.          |
| **`fmt.Scanf`**     | Formatlı bir şekilde kullanıcıdan girdi alır.                            |
| **`fmt.Sscan`**     | Bir string'den boşlukla ayrılan değerleri okur.                          |
| **`fmt.Sscanln`**   | Bir string'den yeni satıra kadar olan değerleri okur.                    |
| **`fmt.Sscanf`**    | Bir string'den formatlı değerleri okur.                                  |
| **`fmt.Fscan`**     | Belirtilen `io.Reader`'dan boşlukla ayrılan değerleri okur.              |
| **`fmt.Fscanln`**   | Belirtilen `io.Reader`'dan yeni satıra kadar olan değerleri okur.        |
| **`fmt.Fscanf`**    | Belirtilen `io.Reader`'dan formatlı değerleri okur.                      |

### Fonksiyon Açıklamaları:

1. **Yazdırma Fonksiyonları**: 
   - `Print`, `Println`, ve `Printf` gibi fonksiyonlar, ekrana yazdırma işlevi görürler.
   - `Fprint`, `Fprintln`, `Fprintf` ise çıktı hedefini belirleyebileceğiniz (`io.Writer` arayüzüne sahip olan) fonksiyonlardır. Örneğin dosya yazımı.
   - `Sprint`, `Sprintln`, ve `Sprintf` ise yazdırılan çıktıyı bir `string` olarak döner.

2. **Girdi Alma Fonksiyonları**:
   - `Scan`, `Scanln`, ve `Scanf`, kullanıcıdan veya standart girdiden (örneğin komut satırından) veri almak için kullanılır.
   - `Sscan`, `Sscanln`, ve `Sscanf` ise bir `string` girdiden veri okur.
   - `Fscan`, `Fscanln`, ve `Fscanf`, belirtilen `io.Reader`'dan veri okumak için kullanılır.

