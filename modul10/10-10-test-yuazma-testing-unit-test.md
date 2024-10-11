# Go ile Test Yazma ve Unit Testing

Go, test yazma konusunda yerleşik destek sağlar. `testing` paketi, testler yazmak ve çalıştırmak için gerekli araçları sunar. Bu makalede, Go'da test yazma yöntemleri, unit testing ile ilgili temel kavramlar ve test yazma örnekleri detaylı bir şekilde ele alınacaktır.

## 1. Test Yazma Temelleri

Go’da bir test yazmak için `*_test.go` uzantılı bir dosya oluşturmanız gerekmektedir. Test fonksiyonları `Test` ile başlamalı ve `*testing.T` parametresi almalıdır. Aşağıda test yazma temelleri ile ilgili bir örnek bulunmaktadır.

### 1.1. Basit Bir Fonksiyon ve Testi

Öncelikle test edilecek basit bir fonksiyon yazalım.

```go
// math.go
package mathutil

// Toplama fonksiyonu
func Add(a int, b int) int {
    return a + b
}
```

Yukarıdaki kodda `Add` adında basit bir toplama fonksiyonu tanımladık. Şimdi bu fonksiyon için bir test yazalım.

### 1.2. Test Dosyasını Oluşturma

Test dosyası için yeni bir dosya oluşturun: `math_test.go`.

```go
// math_test.go
package mathutil

import "testing"

// Add fonksiyonu için test
func TestAdd(t *testing.T) {
    result := Add(1, 2) // Toplama işlemi
    expected := 3       // Beklenen sonuç

    // Sonucu kontrol et
    if result != expected {
        t.Errorf("Add(1, 2) = %d; beklenen %d", result, expected) // Hata mesajı
    }
}
```

Bu test fonksiyonu, `Add` fonksiyonunu çağırır ve sonucu beklenen değerle karşılaştırır. Eğer sonuç beklenenden farklıysa, bir hata mesajı iletildi.

## 2. Testleri Çalıştırma

Testleri çalıştırmak için terminalde aşağıdaki komutu kullanabilirsiniz:

```bash
go test
```

### 2.1. Çıktı

Eğer her şey doğruysa, terminalde şu çıktı görünecektir:

```plaintext
PASS
ok      yourmodule/mathutil  0.001s
```

Eğer hata varsa, terminalde şu şekilde bir çıktı göreceksiniz:

```plaintext
--- FAIL: TestAdd (0.00s)
    math_test.go:10: Add(1, 2) = 4; beklenen 3
FAIL
exit status 1
FAIL    yourmodule/mathutil  0.001s
```

## 3. Diğer Test Senaryoları

### 3.1. Hata Kontrolü

Hata kontrolü için `t.Fatal` veya `t.Errorf` kullanabilirsiniz. Örneğin, bir hata durumunu test etmek için şöyle bir fonksiyon yazabiliriz:

```go
// math.go
package mathutil

// Divide fonksiyonu
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("b sıfır olamaz")
    }
    return a / b, nil
}
```

Şimdi `Divide` fonksiyonu için test yazalım.

```go
// math_test.go
package mathutil

import "testing"

// Divide fonksiyonu için test
func TestDivide(t *testing.T) {
    _, err := Divide(10, 0) // Hata bekleniyor

    // Hata kontrolü
    if err == nil {
        t.Fatal("Divide(10, 0) için hata bekleniyordu, ancak alınamadı") // Hata mesajı
    }
}
```

### 3.2. Testler İçin `Table Driven Tests`

Go, "table-driven tests" olarak bilinen bir test yazma yaklaşımını destekler. Bu yaklaşım, aynı test fonksiyonunu birden fazla giriş için çalıştırarak kodu daha düzenli hale getirir.

**Örnek:**

```go
// math_test.go
package mathutil

import "testing"

// Add fonksiyonu için test
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

    for _, test := range tests {
        result := Add(test.a, test.b) // Toplama işlemi
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; beklenen %d", test.a, test.b, result, test.expected) // Hata mesajı
        }
    }
}
```

### 4. Benchmark Testleri

Go, performans testleri yapmak için benchmark testlerine de destek verir. Benchmark testleri, belirli bir kod parçasının ne kadar sürede çalıştığını ölçmek için kullanılır. Benchmark fonksiyonları `Benchmark` ile başlar ve `*testing.B` parametresi alır.

**Örnek:**

```go
// math_test.go
package mathutil

import "testing"

// Benchmark testi
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2) // Toplama işlemi
    }
}
```

Benchmark testlerini çalıştırmak için şu komutu kullanabilirsiniz:

```bash
go test -bench=.
```

### 4.1. Çıktı

Eğer benchmark testini çalıştırdıysanız, terminalde aşağıdaki gibi bir çıktı göreceksiniz:

```plaintext
BenchmarkAdd-8    1000000000    0.457 ns/op
PASS
ok      yourmodule/mathutil  0.001s
```

Bu çıktı, `Add` fonksiyonunun her çağrısının ortalama 0.457 nanosecond sürdüğünü gösterir.

## 5. Özet

Bu makalede Go'da test yazma ve unit testing konusunu detaylı bir şekilde inceledik. İşte önemli noktalar:

- Test fonksiyonları `*_test.go` dosyalarında yer almalıdır.
- Her test fonksiyonu `Test` ile başlamalı ve `*testing.T` parametresi almalıdır.
- Hata kontrolü için `t.Error`, `t.Errorf`, `t.Fatal` ve `t.Fatalf` kullanılabilir.
- Table-driven tests ile düzenli ve okunabilir testler yazılabilir.
- Benchmark testleri, kodun performansını değerlendirmek için kullanılır.

Go’nun test framework’ü, yazılım geliştirme sürecinde kalite ve güvenilirlik sağlamak için oldukça etkilidir. Geliştiricilerin, uygulamalarını test ederek hata oranını azaltmaları ve kod kalitesini artırmaları hedeflenmektedir.