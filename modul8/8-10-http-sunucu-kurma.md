Go dilinde bir HTTP sunucusu kurmak oldukça basittir. Go, yerleşik `net/http` paketine sahiptir ve bu paket, HTTP sunucuları oluşturmayı kolaylaştırır. Bu kılavuzda, adım adım bir HTTP sunucusu kurmayı öğreneceğiz ve sunucuya bazı temel özellikler ekleyeceğiz.

## 1. Go HTTP Sunucusu Kurma

### 1.1. Proje Dizini Oluşturma

Öncelikle bir proje dizini oluşturup bu dizine geçelim:

```bash
mkdir go-http-server
cd go-http-server
```

### 1.2. Go Modülünü Başlatma

Projenizi modüler hale getirmek için Go modülünü başlatın:

```bash
go mod init go-http-server
```

### 1.3. Basit HTTP Sunucusu

Aşağıda, temel bir HTTP sunucusu oluşturacağız. Bu sunucu, gelen istekleri dinleyecek ve yanıt verecektir.

#### Örnek: Basit HTTP Sunucusu

Aşağıdaki kodu `main.go` dosyasına yazın:

```go
package main

import (
    "fmt"
    "net/http"
)

// HelloHandler, gelen istekleri karşılayacak olan bir handler fonksiyonudur.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // Yanıtın içeriğini ayarlıyoruz.
    w.Header().Set("Content-Type", "text/plain")
    // Yanıt olarak "Merhaba, Dünya!" yazıyoruz.
    fmt.Fprintln(w, "Merhaba, Dünya!")
}

func main() {
    // HTTP sunucusunu "/hello" yolunu dinleyecek şekilde ayarlıyoruz.
    http.HandleFunc("/hello", HelloHandler)

    // Sunucu 8080 portunda dinlemeye başlıyor.
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        // Hata meydana gelirse, hatayı ekrana yazdırıyoruz.
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 2. Kodu Çalıştırma

Aşağıdaki komutu terminalden çalıştırarak sunucuyu başlatabilirsiniz:

```bash
go run main.go
```

### Çıktı

```plaintext
Sunucu 8080 portunda dinliyor...
```

**Açıklama**: Bu mesaj, sunucunun başarıyla başlatıldığını ve 8080 portunda dinlediğini gösterir.

### 3. HTTP İsteği Gönderme

Sunucu çalışırken, başka bir terminal veya tarayıcı kullanarak sunucuya bir istek gönderebiliriz. Tarayıcınızda şu URL'yi ziyaret edin:

```
http://localhost:8080/hello
```

### Çıktı

Tarayıcıda aşağıdaki yanıtı göreceksiniz:

```
Merhaba, Dünya!
```

**Açıklama**: Sunucu, `/hello` yoluna gelen isteği yanıtlar ve "Merhaba, Dünya!" mesajını döner.

### 4. HTTP Yöntemleri

HTTP sunucusu, farklı HTTP yöntemlerini (GET, POST, PUT, DELETE vb.) işleyebilir. Aşağıdaki örnekte, hem GET hem de POST isteklerini nasıl yöneteceğimizi göstereceğiz.

#### Örnek: HTTP Yöntemlerini Kullanma

Aşağıdaki kodu `main.go` dosyasını güncelleyerek yazın:

```go
package main

import (
    "fmt"
    "net/http"
)

// HelloHandler, gelen istekleri karşılayacak olan bir handler fonksiyonudur.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // İstek yöntemini kontrol ediyoruz.
    if r.Method == http.MethodGet {
        // Yanıtın içeriğini ayarlıyoruz.
        w.Header().Set("Content-Type", "text/plain")
        // Yanıt olarak "Merhaba, GET!" yazıyoruz.
        fmt.Fprintln(w, "Merhaba, GET!")
    } else if r.Method == http.MethodPost {
        // Yanıtın içeriğini ayarlıyoruz.
        w.Header().Set("Content-Type", "text/plain")
        // Yanıt olarak "Merhaba, POST!" yazıyoruz.
        fmt.Fprintln(w, "Merhaba, POST!")
    } else {
        // Desteklenmeyen bir yöntem geldiğinde 405 hatası döndürüyoruz.
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintln(w, "Desteklenmeyen istek yöntemi!")
    }
}

func main() {
    // HTTP sunucusunu "/hello" yolunu dinleyecek şekilde ayarlıyoruz.
    http.HandleFunc("/hello", HelloHandler)

    // Sunucu 8080 portunda dinlemeye başlıyor.
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        // Hata meydana gelirse, hatayı ekrana yazdırıyoruz.
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 5. GET ve POST İstekleri Gönderme

#### GET İsteği

GET isteğini göndermek için tarayıcınızı kullanarak şu URL'yi ziyaret edin:

```
http://localhost:8080/hello
```

### Çıktı

```plaintext
Merhaba, GET!
```

**Açıklama**: Sunucu, GET isteğine yanıt olarak "Merhaba, GET!" döner.

#### POST İsteği

POST isteği göndermek için `curl` komutunu kullanabilirsiniz. Terminalde şu komutu çalıştırın:

```bash
curl -X POST http://localhost:8080/hello
```

### Çıktı

```plaintext
Merhaba, POST!
```

**Açıklama**: Sunucu, POST isteğine yanıt olarak "Merhaba, POST!" döner.

### 6. Hata Yönetimi

Go'da hata yönetimi oldukça önemlidir. Yukarıdaki örnekte, `http.ListenAndServe` çağrısında bir hata meydana gelirse, bu hata konsola yazdırılmaktadır.

### 7. JSON Yanıtı Gönderme

Birçok modern uygulama JSON formatında veri gönderir. Aşağıda, JSON formatında yanıt döndüren bir örnek vereceğiz.

#### Örnek: JSON Yanıtı

Aşağıdaki kodu `main.go` dosyasını güncelleyerek yazın:

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// User, JSON yanıtında kullanılacak bir yapıdır.
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// UserHandler, kullanıcı bilgilerini döndüren bir handler fonksiyonudur.
func UserHandler(w http.ResponseWriter, r *http.Request) {
    // Kullanıcı bilgilerini tanımlıyoruz.
    user := User{Name: "Ali Veli", Email: "ali@veli.com"}

    // Yanıtın içeriğini JSON formatında ayarlıyoruz.
    w.Header().Set("Content-Type", "application/json")

    // Kullanıcı bilgilerini JSON formatında yazıyoruz.
    json.NewEncoder(w).Encode(user)
}

func main() {
    // HTTP sunucusunu "/user" yolunu dinleyecek şekilde ayarlıyoruz.
    http.HandleFunc("/user", UserHandler)

    // Sunucu 8080 portunda dinlemeye başlıyor.
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        // Hata meydana gelirse, hatayı ekrana yazdırıyoruz.
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 8. JSON İsteği Gönderme

Sunucu çalışırken, aşağıdaki URL'yi tarayıcınızda ziyaret edin:

```
http://localhost:8080/user
```

### Çıktı

```json
{"name":"Ali Veli","email":"ali@veli.com"}
```

**Açıklama**: Sunucu, `/user` yoluna gelen isteği JSON formatında yanıtlar.

## Özet

- **HTTP Sunucusu Oluşturma**: `net/http` paketi kullanılarak basit bir HTTP sunucusu oluşturduk.
- **HTTP Yöntemleri**: GET ve POST isteklerini nasıl yöneteceğimizi gösterdik.
- **JSON Yanıtı**: Kullanıcı bilgilerini JSON formatında döndüren bir örnek yaptık.
- **Hata Yönetimi**: Sunucu başlatılırken bir hata oluşursa, bu hatayı konsola yazdırıyoruz.

Go dilinde HTTP sunucusu oluşturmak, veri sunmak ve API geliştirmek oldukça kolaydır. Bu temel bilgileri kullanarak daha karmaşık uygulamalar geliştirebilirsiniz.
