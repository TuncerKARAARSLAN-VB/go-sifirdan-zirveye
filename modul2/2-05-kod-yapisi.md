# Kod Yapısı

## 1. `package main`

`package` anahtar kelimesi, Go dilinde bir dosyanın hangi pakete ait olduğunu belirtir. Go programları paketler (modules) halinde düzenlenir ve her Go dosyası bir paketin parçasıdır.

- **`package main`**, Go'da özel bir anlam taşır. `main` paketi, Go uygulamasının başlangıç noktasıdır. Go programında yürütülecek fonksiyon `main` paketinde bulunmalıdır. Diğer bir deyişle, Go programının çalışmaya başlayacağı yer `package main` içindeki `main` fonksiyonudur.
- **Ana uygulama dosyaları**: Eğer `package main` yerine başka bir paket adı kullanırsanız, o dosya bir kütüphane veya yardımcı kod parçası olur, ancak doğrudan çalıştırılamaz.

## Örnek

```go
package main

func main() {
    // Bu fonksiyon programın başladığı yerdir.
    println("Program başlıyor...")
}
```

Bu örnekte `package main` kullanılmış ve bu dosyanın bir uygulama dosyası olduğu belirtilmiştir.

## 2. `import "fmt"`

`import` anahtar kelimesi, Go'da başka paketleri (kütüphaneleri) kodunuza dahil etmek için kullanılır. Bir paket içindeki fonksiyonları, tipleri veya diğer yapıları kullanmak istiyorsanız o paketi içe aktarmanız gerekir.

- **`fmt` paketi**, Go'nun standart kütüphanesinde bulunan bir formattır ve girdileri yazdırma (`print`), girdiyi okuma (`scan`) gibi temel işlevleri içerir. Adı "format" kelimesinden gelir ve sıklıkla kullanılır.
- `fmt` paketindeki en yaygın kullanılan fonksiyonlar `fmt.Println`, `fmt.Printf`, `fmt.Sprint` gibi çıktı yazdırma fonksiyonlarıdır.

## Örnek

```go
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, Dünya!") // Ekrana yazdırır
}
```

Bu örnekte, `import "fmt"` ile `fmt` paketi içe aktarılmıştır ve `fmt.Println` kullanılarak ekrana bir metin yazdırılmıştır.

## `import` İle Birden Fazla Paket Kullanımı

Go'da birden fazla paketi aynı anda import edebilirsiniz. Bunun için her paketi ayrı satırlarda yazabileceğiniz gibi, tek bir `import` bloğu içerisinde toplu halde de yazabilirsiniz.

## Örnek 1

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

## Örnek 2

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Printf("Merhaba %d\n", i)
    }
}
```

- `for` döngüsü, 1'den başlayarak 100'e kadar döner.
- `fmt.Printf` fonksiyonu, ekrana formatlı bir çıktı yazmak için kullanılır. `%d` ifadesi, sayısal bir değeri yazdırmak için kullanılır.
- `\n`, her "Merhaba" ifadesinden sonra yeni bir satıra geçmek için kullanılır.

Bu kodu çalıştırdığınızda ekrana 100 kez şu şekilde bir çıktı verecektir:

```
Merhaba 1
Merhaba 2
Merhaba 3
...
Merhaba 100
```

## Library

Go'da her bağımsız çalıştırılabilir programın `package main` ve `main` fonksiyonuna sahip olması gereklidir. Eğer `package main` olmadan bir Go dosyası yazarsanız, bu dosya bir kütüphane olarak başka Go programları tarafından kullanılabilir, ancak tek başına bir program olarak çalıştırılamaz.

Aşağıda `package main` kullanmadan bir Go dosyası örneği bulunuyor:

```go
package utility

import "fmt"

func init() {
    fmt.Println("Utility paketi içindeki init fonksiyonu çalıştı.")
}

func MesajYazdir() {
    fmt.Println("Utility paketindeki bir fonksiyon.")
}
```

Bu dosya `utility` isimli bir paket tanımlar ve başka bir Go programında şöyle kullanılabilir:

```go
package main

import "path/to/your/utility"

func main() {
    utility.MesajYazdir()
}
```

Burada dikkat edilmesi gereken nokta, `package main` olmadan bir Go dosyası sadece bir kütüphane olarak kullanılabilir ve tek başına çalıştırılamaz. Çalıştırılabilir bir Go programı yazmak için mutlaka `package main` ve `main` fonksiyonuna ihtiyaç vardır.