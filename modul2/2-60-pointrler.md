# Pointerler

Elbette! Go dilinde pointer'lar, bellek adreslerini tutan değişkenlerdir ve veri yönetimini daha etkili hale getirmek için kullanılır. Pointer'lar ile değişkenlerin bellek adresleri üzerinden erişilebilir, böylece bellek kullanımı optimize edilebilir. Bu eğitimde, pointer parametre almayı ve pointer değer döndürmeyi detaylı bir şekilde inceleyeceğiz.

## 1. Pointer Nedir?

Pointer, bir değişkenin bellek adresini tutan bir değişkendir. Go dilinde pointer'lar, `*` sembolü ile tanımlanır. Örneğin:

```go
var x int = 10
var p *int = &x // x'in adresini p'ye atıyoruz
```

Burada `&` operatörü, bir değişkenin adresini almak için kullanılır.

## 2. Pointer Parametre Alma

Pointer parametre almanın faydaları, büyük yapıları ve dizileri fonksiyonlara geçirirken bellek kullanımı ve kopyalama maliyetlerini azaltmaktır. Aşağıda pointer parametre almanın nasıl yapılacağını gösteren bir örnek verilmiştir.

### 2.1 Örnek: Pointer Parametre ile Değişiklik Yapma

```go
package main

import "fmt"

// Değişkenin değerini değiştiren fonksiyon
func updateValue(x *int) {
    *x = 20 // Pointer üzerinden değer güncelleniyor
}

func main() {
    value := 10
    fmt.Println("Başlangıç Değeri:", value) // 10

    updateValue(&value) // Pointer ile değeri güncelleyerek fonksiyona gönderiyoruz
    fmt.Println("Güncellenmiş Değer:", value) // 20
}
```

#### Açıklamalar

- **updateValue**: Bu fonksiyon, bir `int` pointer'ı alır ve işaret ettiği değeri değiştirir. `*x` ifadesi, pointer üzerinden gösterilen değere erişir.
- **main**: `value` değişkeninin adresi `&value` ile fonksiyona gönderilir. Fonksiyon, bu adres üzerinden değeri değiştirir.

## 3. Pointer Değer Döndürme

Fonksiyonlar, pointer döndürebilir ve böylece çağıran fonksiyona bellek adresini iletebilir. Bu, özellikle karmaşık veri yapıları için yararlıdır.

### 3.1 Örnek: Pointer Değeri Döndürme

```go
package main

import "fmt"

// Yeni bir değişken oluşturan ve pointer döndüren fonksiyon
func createPointer() *int {
    value := 30 // Yerel değişken
    return &value // Değişkenin adresini döndürüyoruz
}

func main() {
    p := createPointer() // Pointer'ı alıyoruz
    fmt.Println("Pointer'dan Alınan Değer:", *p) // 30 yazdırılır
}
```

#### Açıklamalar

- **createPointer**: Bu fonksiyon, yeni bir `int` değişken oluşturur ve onun adresini döndürür. 
- **main**: Fonksiyonu çağırarak döndürülen pointer üzerinden değeri alır.

### 3.2 Bellek Sorunu

Yukarıdaki örnekte dikkat edilmesi gereken bir nokta var: Fonksiyon içerisinde tanımlanan `value` değişkeni, fonksiyon tamamlandığında bellekten silinir. Bu, **nil** pointer hatalarına yol açabilir. Aşağıdaki örnekte bu durumu gösterelim:

```go
package main

import "fmt"

// Yanlış bir pointer döndüren fonksiyon
func createInvalidPointer() *int {
    var value int = 40 // Yerel değişken
    return &value // Hatalı: value, fonksiyon bitiminde geçersiz hale gelir
}

func main() {
    p := createInvalidPointer() // Pointer'ı alıyoruz
    fmt.Println("Pointer'dan Alınan Değer:", *p) // Hata: geçersiz bellek erişimi
}
```

### Açıklama

- Bu kod çalıştırıldığında, geçersiz bellek erişimi hatası alırsınız. Çünkü `value` değişkeni, fonksiyon tamamlandığında geçerliliğini yitirir. Bunun önüne geçmek için ya global değişkenler kullanmalı ya da heap'te (dinamik bellek) yer alan değişkenler oluşturmalısınız.

## 4. Bellekte Dinamik Değişken Oluşturma

Go'da dinamik bellek yönetimi için `new` ve `make` fonksiyonları kullanılabilir. Bu yöntemlerle bellek yönetimi sorunlarının üstesinden gelebiliriz.

### 4.1 `new` ile Pointer Oluşturma

```go
package main

import "fmt"

// new ile dinamik bellek yönetimi
func createDynamicPointer() *int {
    p := new(int) // int türünde yeni bir pointer oluştur
    *p = 50      // Pointer'ın gösterdiği değeri atıyoruz
    return p    // Pointer'ı döndür
}

func main() {
    p := createDynamicPointer() // Dinamik olarak oluşturulmuş pointer
    fmt.Println("Dinamik Pointer'dan Alınan Değer:", *p) // 50 yazdırılır
}
```

### Açıklamalar

- **new**: `new(int)` ifadesi, bellekte `int` türünde yeni bir alan oluşturur ve pointer olarak döndürür. Bu bellek alanı, fonksiyon dışında da geçerlidir.

## 5. Özet

- **Pointer Parametre Alma**: Pointer'lar, fonksiyonlara veri geçerken daha verimli bir yöntem sağlar. Fonksiyon içindeki değişiklikler, çağıran fonksiyonda da görünür.
- **Pointer Değer Döndürme**: Fonksiyonlar, pointer döndürerek bellek adresini dışarıya iletebilir. Ancak bu, dikkatli bir şekilde yönetilmelidir.
- **Dinamik Bellek Yönetimi**: `new` ve `make` fonksiyonları, dinamik bellek yönetimi için kullanılır ve geçerli bellek adresleri döndürür.

Bu eğitim ile Go'da pointer parametre alma ve pointer değer döndürme konularında detaylı bilgi sahibi oldunuz. Daha fazla bilgi ya da örnek isterseniz, lütfen belirtin!