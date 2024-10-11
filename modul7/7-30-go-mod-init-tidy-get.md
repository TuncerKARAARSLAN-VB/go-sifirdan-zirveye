Go dilinde modül yönetimi, projelerin bağımlılıklarını ve sürüm kontrollerini düzenlemek için kullanılır. Go 1.11 ile birlikte tanıtılan modül sistemi, projelerin daha iyi yönetilmesini sağlamaktadır. Bu modül yönetim sistemi, `go mod` komutlarıyla yönetilir. Bu yazıda `go mod init`, `go mod tidy` ve `go get` komutlarını detaylı bir şekilde inceleyeceğiz ve her birinin kullanımına dair örnekler vereceğiz.

## 1. go mod init

`go mod init`, bir Go projesinin modül olarak başlatılmasını sağlayan komuttur. Bu komut çalıştırıldığında, belirtilen modül adıyla birlikte bir `go.mod` dosyası oluşturulur. Bu dosya, modülün adını, sürüm bilgilerini ve bağımlılıkları tutar.

### Örnek: go mod init Kullanımı

```bash
# Proje dizinine gidin
mkdir myproject
cd myproject

# Go modülünü başlat
go mod init github.com/kullanici/myproject
```

### Çıktı:

```
go: creating new go.mod: module github.com/kullanici/myproject
```

**Açıklama**:
- `mkdir myproject`: Yeni bir proje dizini oluşturur.
- `cd myproject`: Oluşturulan dizine geçiş yapar.
- `go mod init github.com/kullanici/myproject`: Belirtilen isimle bir Go modülü başlatır ve `go.mod` dosyası oluşturur.

Oluşan `go.mod` dosyası aşağıdaki gibi görünebilir:

```go
module github.com/kullanici/myproject

go 1.20
```

## 2. go mod tidy

`go mod tidy`, proje için gerekli olan bağımlılıkları temizler ve eksik olanları ekler. Yani, `go.mod` dosyasında ve proje dizinindeki kodlarda bulunan gereksiz bağımlılıkları kaldırırken, gerekli olanları da ekler. Bu, projede kullanılan bağımlılıkların güncel ve doğru bir şekilde tanımlanmasını sağlar.

### Örnek: go mod tidy Kullanımı

Öncelikle, bir modül oluşturalım ve bağımlılık ekleyelim:

```bash
# Örnek bir Go dosyası oluşturun
touch main.go
```

`main.go` dosyası içerisine basit bir örnek yazalım:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("Rastgele sayı:", rand.Intn(100))
}
```

### Bağımlılıkları Ekleme

Yukarıdaki kodda `math/rand` paketi kullanılmıştır. Ancak bu paketi projeye bağımlılık olarak eklemedik. Şimdi `go mod tidy` komutunu çalıştıralım.

```bash
# Projeyi temizle
go mod tidy
```

### Çıktı:

```
go: finding module for package math/rand
go: found math/rand in module runtime
```

**Açıklama**:
- `go mod tidy`: Projedeki bağımlılıkları kontrol eder ve eksik olanları ekler. `math/rand` paketi için gerekli bağımlılıkları bulur ve `go.mod` dosyasına ekler.

Oluşan `go.mod` dosyası aşağıdaki gibi görünebilir:

```go
module github.com/kullanici/myproject

go 1.20

require (
    math/rand v0.0.0-20220928120000-7deaf59f2f12 // örnek bir versiyon
)
```

## 3. go get

`go get`, bir modül veya paket için bağımlılıkları indiren ve bunları `go.mod` dosyasına ekleyen bir komuttur. Belirli bir modülün en son sürümünü ya da belirli bir sürümünü indirmek için kullanılır.

### Örnek: go get Kullanımı

Öncelikle, yeni bir bağımlılık eklemek için `go get` komutunu kullanabiliriz. Örneğin, `gorilla/mux` adlı popüler bir HTTP yönlendirme kütüphanesini ekleyelim:

```bash
# gorilla/mux modülünü ekle
go get github.com/gorilla/mux
```

### Çıktı:

```
go: downloading github.com/gorilla/mux v1.8.0
go: added github.com/gorilla/mux v1.8.0
```

**Açıklama**:
- `go get github.com/gorilla/mux`: `gorilla/mux` kütüphanesinin en son sürümünü indirir ve `go.mod` dosyasına ekler.
- İndirme işlemi sırasında, kütüphanenin sürümü ve diğer bağımlılıklar hakkında bilgi verir.

`go.mod` dosyası artık aşağıdaki gibi görünebilir:

```go
module github.com/kullanici/myproject

go 1.20

require (
    github.com/gorilla/mux v1.8.0
)
```

### Özet

- **go mod init**: Yeni bir Go modülü başlatmak için kullanılır ve `go.mod` dosyası oluşturur.
- **go mod tidy**: Proje için gerekli olan bağımlılıkları kontrol eder, eksik olanları ekler ve gereksiz olanları kaldırır.
- **go get**: Belirli bir modülü veya paketi indirir ve `go.mod` dosyasına ekler.

Bu komutlar, Go modül yönetiminin temel taşlarını oluşturur ve projelerin daha düzenli ve sürdürülebilir bir şekilde geliştirilmesine olanak tanır. Go modülleriyle ilgili daha fazla bilgi için [Go resmi belgelerine](https://golang.org/doc/modules/) göz atabilirsiniz.