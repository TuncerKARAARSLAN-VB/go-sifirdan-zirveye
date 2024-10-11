# Go'da Mocklama ve Hata Yakalama

Go programlama dilinde mocklama ve hata yakalama, yazılım geliştirme sürecinde önemli iki kavramdır. Mocklama, bağımlılıkları simüle etmek için kullanılırken, hata yakalama, uygulamanızın sağlamlığını ve güvenilirliğini artırmak için kritik bir rol oynar. Bu yazıda, mocklama ve hata yakalama konularını detaylı bir şekilde ele alacağız ve örneklerle açıklayacağız.

## 1. Mocklama

Mocklama, bir bağımlılığı veya dış bileşeni taklit etmek için kullanılan bir tekniktir. Genellikle birim testleri sırasında, gerçek bağımlılıkların yerine geçecek sahte (mock) nesneler oluşturulur. Bu sayede testler daha hızlı ve güvenilir bir şekilde çalışır.

### 1.1. Mocklama İçin `github.com/stretchr/testify/mock` Kütüphanesi

Go'da mocklama için yaygın olarak kullanılan bir kütüphane `testify`'dir. Bu kütüphane, mock nesneleri kolayca oluşturmanıza olanak tanır.

#### 1.1.1. Mock Bir Arayüz Oluşturma

Öncelikle bir arayüz oluşturalım. Bu arayüz, dış bir servisi temsil edecek.

```go
// service.go
package service

// Notifier arayüzü, bir bildirim göndermek için bir metot tanımlar.
type Notifier interface {
    Notify(message string) error
}
```

#### 1.1.2. Gerçek Bir Servis

Gerçek bir servisi aşağıdaki gibi tanımlayabiliriz:

```go
// real_service.go
package service

import "fmt"

// RealNotifier, Notifier arayüzünü uygulayan gerçek bir yapıdır.
type RealNotifier struct{}

// Notify metodu, bir bildirim gönderir.
func (rn *RealNotifier) Notify(message string) error {
    fmt.Println("Gönderilen mesaj:", message) // Mesaj konsola yazılır
    return nil // Hata yok
}
```

#### 1.1.3. Mock Servisi

Mock servisi, `Notifier` arayüzünü taklit etmek için kullanılır:

```go
// mock_service.go
package service

import "github.com/stretchr/testify/mock"

// MockNotifier, Notifier arayüzünü taklit eden bir mock yapısıdır.
type MockNotifier struct {
    mock.Mock
}

// Notify metodu, mock davranışını tanımlamak için kullanılır.
func (mn *MockNotifier) Notify(message string) error {
    args := mn.Called(message) // Gelen mesajı alır
    return args.Error(0)       // Hata döndürür
}
```

#### 1.1.4. Test Dosyası

Şimdi mock servisi kullanarak bir test yazalım:

```go
// service_test.go
package service

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// TestNotify fonksiyonu, Notify metodunu test eder.
func TestNotify(t *testing.T) {
    // MockNotifier oluştur
    mockNotifier := new(MockNotifier)

    // Beklenen çağrıyı ayarla
    mockNotifier.On("Notify", "Test mesajı").Return(nil)

    // Notify metodunu çağır
    err := mockNotifier.Notify("Test mesajı")

    // Hata kontrolü
    assert.Nil(t, err) // Hata yoksa başarılı
    mockNotifier.AssertExpectations(t) // Beklentileri kontrol et
}
```

### 1.2. Test Çıktısı

Bu testi çalıştırmak için terminalde aşağıdaki komutu kullanabilirsiniz:

```bash
go test
```

Eğer test başarılıysa, aşağıdaki gibi bir çıktı alırsınız:

```plaintext
PASS
ok      yourmodule/service   0.001s
```

## 2. Hata Yakalama

Go'da hata yakalama, uygulamanızın sağlamlığını sağlamak için kritik bir beceridir. Hatalar, genellikle `error` türünde döndürülür ve bu hataları yakalamak, uygulamanızın beklenmedik bir şekilde çökmesini önler.

### 2.1. Hata Yönetimi

Go'da hata yönetimi, genellikle aşağıdaki gibi yapılır:

```go
package main

import (
    "errors"
    "fmt"
)

// Divide fonksiyonu, iki sayıyı böler.
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("b sıfır olamaz") // Hata durumu
    }
    return a / b, nil // Başarılı sonuç
}

func main() {
    // Bölme işlemi
    result, err := Divide(10, 0)
    if err != nil {
        fmt.Println("Hata:", err) // Hata durumunu yakala
        return // Hata varsa işlemi durdur
    }
    fmt.Println("Sonuç:", result) // Sonuç yazdır
}
```

### 2.2. Hata Kontrolü

`Divide` fonksiyonunu çağırdıktan sonra, hatayı kontrol ediyoruz. Hata varsa, hatayı yazdırıp işlemi sonlandırıyoruz.

### 2.3. Çıktı

Eğer yukarıdaki kodu çalıştırırsanız, şu çıktıyı alırsınız:

```plaintext
Hata: b sıfır olamaz
```

Eğer `Divide(10, 2)` gibi bir çağrı yaparsanız, aşağıdaki gibi bir çıktı alırsınız:

```plaintext
Sonuç: 5
```

## 3. Hata Yakalama ve Mocklama Birlikte

Hata yakalama ve mocklama genellikle birlikte kullanılır. Örneğin, mocklama ile bir servis oluştururken, belirli bir hatayı simüle edebilirsiniz.

### 3.1. Mock Servis ile Hata Simülasyonu

Mock servis içinde bir hata durumu oluşturmak için aşağıdaki gibi bir test yazabilirsiniz:

```go
// service_test.go
package service

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

// TestNotifyWithError fonksiyonu, hata durumu simüle eder.
func TestNotifyWithError(t *testing.T) {
    mockNotifier := new(MockNotifier)

    // Hata durumu ayarla
    mockNotifier.On("Notify", "Test mesajı").Return(errors.New("mock hata"))

    // Notify metodunu çağır
    err := mockNotifier.Notify("Test mesajı")

    // Hata kontrolü
    assert.NotNil(t, err) // Hata olmalı
    assert.Equal(t, "mock hata", err.Error()) // Beklenen hata mesajı
    mockNotifier.AssertExpectations(t) // Beklentileri kontrol et
}
```

### 3.2. Test Çıktısı

Bu testi çalıştırdığınızda, aşağıdaki gibi bir çıktı alırsınız:

```plaintext
PASS
ok      yourmodule/service   0.001s
```

Eğer hata durumunu doğru bir şekilde yakaladıysanız, test başarılı olacaktır.

## 4. Özet

Bu makalede, Go'da mocklama ve hata yakalama konularını detaylı bir şekilde ele aldık. İşte önemli noktalar:

- **Mocklama**, bağımlılıkları taklit etmek için kullanılır ve genellikle birim testleri sırasında uygulanır.
- **`testify` kütüphanesi**, mock nesneleri oluşturmak için yaygın olarak kullanılır.
- Hata yakalama, uygulamanızın sağlamlığını sağlamak için kritik bir rol oynar.
- Hataları yakalamak için genellikle `error` türü kullanılır ve `nil` kontrolü yapılır.

Mocklama ve hata yakalama, yazılım geliştirme sürecinde kodun güvenilirliğini ve bakımını artırmak için oldukça önemlidir. Bu becerileri uygulamak, daha sağlam ve hatasız bir yazılım geliştirmeye katkı sağlayacaktır.