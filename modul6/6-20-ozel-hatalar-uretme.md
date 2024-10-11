# Özel Hatalar

Go dilinde özel hata üretimi, hata yönetiminin önemli bir parçasıdır. Özel hatalar, hataların daha anlamlı ve açıklayıcı bir şekilde yönetilmesine yardımcı olur. Bu eğitimde, özel hata türlerinin nasıl oluşturulacağı, kullanılacağı ve örneklerle açıklanacağı üzerinde duracağız.

## 1. Özel Hata Türleri

Go dilinde özel hata türleri, `error` arayüzünü uygulayan yapı (struct) olarak tanımlanabilir. Bu özel hata türleri, hata mesajları ve hata ile ilgili ek bilgiler taşıyabilir.

### 1.1 Özel Hata Türü Oluşturma

Aşağıda, özel bir hata türü oluşturan basit bir örnek bulunmaktadır:

```go
package main

import (
    "fmt"
)

// Özel hata türü
type MyError struct {
    Code    int
    Message string
}

// MyError struct'ı için Error() metodu
func (e *MyError) Error() string {
    return fmt.Sprintf("Hata Kodu: %d, Hata Mesajı: %s", e.Code, e.Message)
}

// Hata döndüren bir fonksiyon
func someFunction() error {
    return &MyError{Code: 404, Message: "Bulunamadı"} // 404 hatası
}

func main() {
    err := someFunction() // Hata döndüren fonksiyonu çağırıyoruz
    if err != nil {
        fmt.Println("Hata:", err) // Hata mesajını yazdırıyoruz
    }
}
```

#### Açıklamalar

- **MyError**: Bu yapı, bir hata kodu ve mesajı içerir.
- **Error() Metodu**: `Error()` metodu, hata bilgilerini formatlayarak döndürür.
- **someFunction**: Bu fonksiyon, özel bir hata döndürür. Ana fonksiyon içinde bu hata kontrol edilir.

### 1.2 Özel Hata Türlerini Kullanma

Özel hata türleri, hata türünü belirlemek ve farklı hatalar arasında ayırt etme yeteneği sağlar. Aşağıdaki örnekte, belirli bir hata türüne göre işlem yapacağız.

```go
package main

import (
    "fmt"
)

// Özel hata türü
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Hata Kodu: %d, Hata Mesajı: %s", e.Code, e.Message)
}

// Belirli bir hata koduna göre işlem yapan bir fonksiyon
func errorHandler(err error) {
    if myErr, ok := err.(*MyError); ok {
        fmt.Printf("Özel Hata Yakalandı: %s\n", myErr)
        switch myErr.Code {
        case 404:
            fmt.Println("Kaynak bulunamadı. Lütfen kontrol edin.")
        case 500:
            fmt.Println("Sunucu hatası. Daha sonra tekrar deneyin.")
        default:
            fmt.Println("Bilinmeyen hata.")
        }
    } else {
        fmt.Println("Genel Hata:", err)
    }
}

func someFunction() error {
    return &MyError{Code: 404, Message: "Bulunamadı"} // 404 hatası
}

func main() {
    err := someFunction() // Hata döndüren fonksiyonu çağırıyoruz
    errorHandler(err) // Hata işleme fonksiyonunu çağırıyoruz
}
```

#### Açıklamalar

- **errorHandler**: Bu fonksiyon, gelen hatanın türünü kontrol eder. Eğer hata, `MyError` türündeyse, hata koduna göre özel işlemler yapar.
- **main**: Hata döndüren bir fonksiyon çağrılır ve hata işleme fonksiyonu ile hata yönetimi gerçekleştirilir.

## 2. Hata Kapsülleme

Go 1.13 ile birlikte, hataları daha iyi yönetmek için hata kapsülleme (`error wrapping`) özelliği eklendi. Bu, hataları daha anlamlı bir şekilde iletmenizi sağlar.

### 2.1 Hata Kapsülleme Örneği

```go
package main

import (
    "errors"
    "fmt"
)

// Özel hata türü
type MyError struct {
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}

// Hata döndüren bir fonksiyon
func someFunction() error {
    return &MyError{Message: "Özel hata oluştu!"}
}

// Başka bir fonksiyon, hatayı sararak döndürür
func anotherFunction() error {
    err := someFunction()
    if err != nil {
        return fmt.Errorf("anotherFunction içinde hata: %w", err) // Hata kapsülleme
    }
    return nil
}

func main() {
    err := anotherFunction() // Hata döndüren fonksiyonu çağırıyoruz
    if err != nil {
        fmt.Println("Hata:", err) // Hata mesajını yazdırıyoruz
        if errors.Is(err, &MyError{}) { // Özel hatayı kontrol et
            fmt.Println("Bu bir MyError hatasıdır.")
        }
    }
}
```

#### Açıklamalar

- **fmt.Errorf**: Hata kapsülleme işlemi için `%w` formatı kullanılır. Bu, hatanın üzerine yeni bir hata mesajı ekler.
- **errors.Is**: `errors.Is` fonksiyonu, hatanın belirli bir türde olup olmadığını kontrol eder.

## 3. Hata Yönetiminde İyi Uygulamalar

- **Hataları anlamlı bir şekilde tanımlayın**: Hatalar, kullanıcıların ve geliştiricilerin anlamasını kolaylaştıracak şekilde tasarlanmalıdır.
- **Hataları kayıt altına alın**: Hata günlükleri, hata ayıklama ve izleme süreçlerinde yardımcı olabilir.
- **Kapsülleme kullanın**: Hata kapsülleme, hataların kökenine inmeyi kolaylaştırır.
- **Hata türlerini tanımlayın**: Özel hata türleri, hatalarınızı daha iyi yönetmenizi sağlar.

## 4. Sonuç

Go dilinde özel hata üretimi, hata yönetiminde büyük bir esneklik sağlar. Özel hata türlerini kullanarak anlamlı hata mesajları oluşturabilir, hata kapsülleme ile hataları daha iyi yönetebilirsiniz. Bu eğitim ile, Go'da özel hata yönetimi hakkında bilgi edindiniz. Eğer daha fazla örnek veya detay isterseniz, lütfen belirtin!