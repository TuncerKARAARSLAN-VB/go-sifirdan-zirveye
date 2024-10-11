# Go'da Test Fonksiyonları ve Test Dosyaları Üretme

Go dilinde test yazma, yazılım geliştirmenin önemli bir parçasıdır. Bu bölümde, Go'da test fonksiyonları oluşturmayı ve test dosyaları üretmeyi detaylı bir şekilde ele alacağız. Örnekler ve açıklamalar ile birlikte, test yazımının nasıl yapıldığını anlamaya çalışacağız.

## 1. Test Fonksiyonları

Go'da test fonksiyonları, genellikle `*_test.go` uzantılı dosyaların içinde yer alır ve `testing` paketini kullanarak yazılır. Her test fonksiyonu, `Test` ile başlar ve `*testing.T` parametresini alır.

### 1.1. Temel Test Fonksiyonu Örneği

Öncelikle basit bir toplama fonksiyonu yazalım ve bu fonksiyon için bir test oluşturalım.

#### 1.1.1. Toplama Fonksiyonu

```go
// math.go
package mathutil

// Add fonksiyonu, iki tam sayıyı toplar.
func Add(a int, b int) int {
    return a + b
}
```

Bu `Add` fonksiyonu, iki tam sayıyı alıp toplar ve sonucu döndürür.

#### 1.1.2. Test Fonksiyonu

Şimdi bu fonksiyon için bir test dosyası oluşturalım.

```go
// math_test.go
package mathutil

import "testing"

// TestAdd fonksiyonu, Add fonksiyonunu test eder.
func TestAdd(t *testing.T) {
    result := Add(1, 2) // Add fonksiyonu çağrılıyor
    expected := 3       // Beklenen sonuç

    // Sonucu kontrol et
    if result != expected {
        // Hata mesajı, eğer sonuç beklenenden farklıysa
        t.Errorf("Add(1, 2) = %d; beklenen %d", result, expected)
    }
}
```

### 1.2. Test Fonksiyonu Açıklamaları

- `TestAdd`: Bu, `Add` fonksiyonunu test eden bir test fonksiyonudur.
- `t *testing.T`: Bu parametre, test sırasında hata bildirimleri yapmak için kullanılır.
- `result := Add(1, 2)`: `Add` fonksiyonu çağrılarak sonuç alınır.
- `t.Errorf(...)`: Eğer beklenen sonuç ile gerçek sonuç eşleşmezse, hata mesajı yazdırılır.

### 1.3. Testi Çalıştırma

Testi çalıştırmak için terminalde aşağıdaki komutu kullanabilirsiniz:

```bash
go test
```

#### 1.3.1. Çıktı

Eğer test başarıyla geçtiyse, terminalde şu çıktıyı görürsünüz:

```plaintext
PASS
ok      yourmodule/mathutil  0.001s
```

Eğer bir hata varsa, terminalde aşağıdaki gibi bir hata mesajı göreceksiniz:

```plaintext
--- FAIL: TestAdd (0.00s)
    math_test.go:10: Add(1, 2) = 4; beklenen 3
FAIL
exit status 1
FAIL    yourmodule/mathutil  0.001s
```

## 2. Birden Fazla Test Fonksiyonu Yazma

Go'da birden fazla test fonksiyonu yazmak mümkündür. Her fonksiyon, farklı senaryoları test edebilir.

### 2.1. Hata Kontrolü

Hata kontrolü yaparak daha sağlam bir test yazabiliriz. Örneğin, sıfıra bölme durumunu kontrol eden bir fonksiyon ve buna ait test:

#### 2.1.1. Bölme Fonksiyonu

```go
// math.go
package mathutil

import "fmt"

// Divide fonksiyonu, iki tam sayıyı böler.
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("b sıfır olamaz") // Hata durumu
    }
    return a / b, nil // Başarılı sonuç
}
```

#### 2.1.2. Bölme Test Fonksiyonu

```go
// math_test.go
package mathutil

import "testing"

// TestDivide fonksiyonu, Divide fonksiyonunu test eder.
func TestDivide(t *testing.T) {
    _, err := Divide(10, 0) // Hata durumu bekleniyor

    // Hata kontrolü
    if err == nil {
        t.Fatal("Divide(10, 0) için hata bekleniyordu, ancak alınamadı")
    }
}
```

### 2.2. Test Fonksiyonu Açıklamaları

- `Divide(10, 0)`: Bu, sıfıra bölme durumunu test etmek için çağrılır.
- `t.Fatal(...)`: Eğer hata alınmazsa, bu metod çağrılır ve testi durdurur.

## 3. Table Driven Tests

Table-driven tests, Go'da test yazmanın en popüler yöntemlerinden biridir. Bu yöntem, aynı test fonksiyonunu birden fazla giriş için çalıştırmayı sağlar.

### 3.1. Örnek

#### 3.1.1. Toplama Fonksiyonu için Table Driven Test

```go
// math_test.go
package mathutil

import "testing"

// TestAddTableDriven fonksiyonu, Add fonksiyonu için tablo bazlı test yapar.
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {2, 3, 5},
        {10, 10, 20},
        {-1, 1, 0},
    }

    // Testleri döngü ile kontrol et
    for _, test := range tests {
        result := Add(test.a, test.b) // Toplama işlemi
        if result != test.expected {
            // Hata mesajı
            t.Errorf("Add(%d, %d) = %d; beklenen %d", test.a, test.b, result, test.expected)
        }
    }
}
```

### 3.2. Table Driven Test Açıklamaları

- `tests`: Test senaryolarının bir listesini tutan bir yapı dizisidir.
- `for _, test := range tests`: Her bir test senaryosunu döngü ile kontrol eder.
- `result := Add(test.a, test.b)`: Her test senaryosu için `Add` fonksiyonu çağrılır.

## 4. Benchmark Testleri

Go, performans testleri yapmak için de bir altyapı sağlar. Benchmark testleri, belirli bir kod parçasının ne kadar sürede çalıştığını ölçer.

### 4.1. Örnek

#### 4.1.1. Benchmark Testi Yazma

```go
// math_test.go
package mathutil

import "testing"

// BenchmarkAdd fonksiyonu, Add fonksiyonunun performansını test eder.
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2) // Toplama işlemi
    }
}
```

### 4.2. Benchmark Testi Çalıştırma

Benchmark testlerini çalıştırmak için terminalde şu komutu kullanabilirsiniz:

```bash
go test -bench=.
```

#### 4.2.1. Çıktı

Eğer benchmark testi başarıyla geçerse, terminalde aşağıdaki gibi bir çıktı alacaksınız:

```plaintext
BenchmarkAdd-8    1000000000    0.457 ns/op
PASS
ok      yourmodule/mathutil  0.001s
```

Bu çıktı, `Add` fonksiyonunun her çağrısının ortalama 0.457 nanosecond sürdüğünü gösterir.

## 5. Test Dosyaları Üretme

Test dosyaları, genellikle projelerinizin aynı dizininde yer alır ve `*_test.go` uzantısına sahiptir. Bu dosyalar, ilgili kodun işlevselliğini kontrol etmek için yazılmış test fonksiyonlarını içerir.

### 5.1. Örnek Proje Yapısı

Örnek bir proje yapısı aşağıdaki gibidir:

```
/yourmodule
    ├── math.go
    ├── math_test.go
```

### 5.2. Test Dosyası Oluşturma

Yeni bir test dosyası oluşturmak için, mevcut kod dosyanızın dizininde yeni bir dosya oluşturun. Örneğin:

- `math_test.go`

İçerisinde test fonksiyonlarını yazabilirsiniz.

## 6. Özet

Bu makalede Go'da test fonksiyonları yazma ve test dosyaları oluşturma konularını detaylı bir şekilde inceledik. İşte önemli noktalar:

- Test fonksiyonları `*_test.go` dosyalarında yer almalıdır.
- Her test fonksiyonu `Test` ile başlamalı ve `*testing.T` parametresi almalıdır.
- Hata kontrolü ve tablodan sürüklenen testler, daha düzenli ve etkili testler yazmamıza yardımcı olur.
- Benchmark testleri, kod performansını değerlend

irmek için kullanılır.

Go’nun test framework’ü, yazılım geliştirme sürecinde kod kalitesini artırmak için oldukça etkilidir. Bu, yazılım geliştiricilerin kodlarını güvenilir bir şekilde test etmelerini ve hata oranını azaltmalarını sağlar.