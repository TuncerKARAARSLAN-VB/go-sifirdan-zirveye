**Middleware**, web uygulamalarında isteklere ve yanıtlara ek işlevler eklemek için kullanılan bir tasarım desenidir. Go dilinde, middleware kullanarak HTTP isteklerini işlemek, doğrulama yapmak, günlüğe kaydetmek veya hata yönetimi gibi ek işlevler ekleyebiliriz. Middleware, genellikle istek işleme akışının bir parçası olarak çalışır ve zincirleme olarak birbirine bağlanabilir.

### Middleware Nedir?

Middleware, bir web uygulamasındaki HTTP isteklerinin ve yanıtlarının işlenmesi sırasında yapılan arka planda çalışan kod parçalarıdır. Genellikle, middleware işlevleri, isteklerin ve yanıtların belirli aşamalarında çalışır. Örneğin:

- İstek geldiğinde belirli bir işlem yapılabilir (örneğin, oturum doğrulama).
- İstemciden gelen veriler üzerinde doğrulama yapılabilir.
- Yanıt oluşturulmadan önce yanıt üzerinde değişiklikler yapılabilir.
- Hatalar, merkezi bir yerde işlenebilir.

### 1. Basit Middleware Örneği

Aşağıda, basit bir middleware örneği ile HTTP isteklerini nasıl yönlendireceğimizi ve ek işlevler ekleyeceğimizi göreceğiz.

#### 1.1 Proje Dizini Oluşturma

```bash
mkdir middleware-example
cd middleware-example
go mod init middleware-example
```

#### 1.2 `main.go` Dosyasını Oluşturma

`main.go` dosyasını oluşturun ve aşağıdaki kodu yazın:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware, isteğin zamanını kaydeden bir middleware'dir.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now() // İşlem başlangıç zamanını al
        next.ServeHTTP(w, r) // Bir sonraki handler'ı çağır
        duration := time.Since(start) // İşlem süresini hesapla
        log.Printf("Request %s %s took %v", r.Method, r.URL.Path, duration) // İstek bilgilerini günlüğe kaydet
    })
}

// HelloHandler, basit bir "Hello, World!" yanıtı döndüren bir handler'dır.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!") // Yanıt olarak "Hello, World!" yazdır
}

// main fonksiyonu, uygulamanın başlangıç noktasıdır.
func main() {
    mux := http.NewServeMux() // Yeni bir HTTP yönlendirici oluştur
    mux.HandleFunc("/", HelloHandler) // "/" rotasına HelloHandler'ı ekle

    // Middleware'i uygulama
    loggedMux := LoggingMiddleware(mux) // LoggingMiddleware'i kullanarak yönlendirici oluştur

    // HTTP sunucusunu başlat
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", loggedMux) // Middleware uygulaması ile sunucuyu başlat
    if err != nil {
        log.Fatal("Sunucu başlatılamadı:", err)
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

**Açıklama**: Sunucu başarıyla başlatıldı ve 8080 portunda dinlemeye başladı.

### 3. HTTP İsteğini Gönderme

Tarayıcı veya `curl` kullanarak sunucuya istek gönderebilirsiniz:

```bash
curl http://localhost:8080/
```

### Çıktı

```plaintext
Hello, World!
```

**Açıklama**: Sunucu "Hello, World!" yanıtını döndürdü.

### 4. Günlük Kayıtları

Sunucunun konsolunda aşağıdaki gibi bir günlük kaydı göreceksiniz:

```plaintext
Request GET / took 1.234567ms
```

**Açıklama**: Middleware, isteğin ne kadar sürdüğünü ve istek bilgilerini günlüğe kaydetti.

### 5. Ek Middleware Örnekleri

Middleware kullanarak başka işlevler de ekleyebiliriz. Örneğin, bir doğrulama middleware'i oluşturabiliriz.

#### 5.1. Doğrulama Middleware

Kullanıcıların yalnızca belirli bir token ile erişebileceği bir rota oluşturabiliriz. Aşağıdaki kodu `main.go` dosyanıza ekleyin:

```go
// AuthMiddleware, yalnızca geçerli bir token ile erişimi kontrol eden bir middleware'dir.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization") // Authorization başlığını al
        if token != "Bearer mysecrettoken" { // Eğer token geçersizse
            http.Error(w, "Unauthorized", http.StatusUnauthorized) // Yetkisiz hatası döndür
            return
        }
        next.ServeHTTP(w, r) // Geçerli ise bir sonraki handler'ı çağır
    })
}

// main fonksiyonu içinde
mux.HandleFunc("/protected", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is a protected route!") // Korunan rotada yanıt döndür
})))
```

### 6. Kodu Çalıştırma ve Test Etme

Sunucuyu yeniden başlatın ve korunan rota için istek gönderin:

```bash
curl http://localhost:8080/protected
```

### Çıktı

```plaintext
Unauthorized
```

**Açıklama**: Token verilmediği için erişim reddedildi.

#### 6.1. Geçerli Token ile İstek Gönderme

Geçerli bir token ile korunan rotaya istek gönderin:

```bash
curl -H "Authorization: Bearer mysecrettoken" http://localhost:8080/protected
```

### Çıktı

```plaintext
This is a protected route!
```

**Açıklama**: Geçerli bir token sağlandığı için korunan rotaya erişim sağlandı.

### 7. Özet

- **Middleware**: İstek işleme akışında ek işlevler eklemek için kullanılır.
- **Günlük Kaydı**: Middleware ile isteklerin sürelerini ve bilgilerini günlüğe kaydedebiliriz.
- **Doğrulama**: Geçerli bir token kontrolü ile korunan rotalar oluşturabiliriz.
- **Zincirleme Middleware**: Middleware'leri birden fazla işlevi bir arada kullanmak için zincirleme yapabiliriz.

Middleware, web uygulamanızın esnekliğini ve işlevselliğini artırmak için güçlü bir araçtır. Uygulamanıza kolayca yeni işlevler ekleyebilir ve istekleri yönetebilirsiniz.