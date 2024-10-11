Mikro servis mimarisi geliştirme süreci, birçok aşamadan oluşur ve her aşama, uygulamanın başarısını etkileyen önemli adımları içerir. Aşağıda, mikro servis mimarisi geliştirme adımlarını detaylı bir şekilde ele alacağız ve her adım için örnekler sunacağız.

## 1. Proje Planlama ve Tasarım

### 1.1. Gereksinimleri Belirleme

Mikro servis mimarisine geçmeden önce, uygulamanızın gereksinimlerini ve kullanıcı ihtiyaçlarını belirlemeniz önemlidir. Gereksinimler, hangi işlevlerin gerekli olduğunu anlamanıza yardımcı olur. Bu aşamada aşağıdaki soruları yanıtlamaya çalışın:

- Kullanıcılar ne tür hizmetlere ihtiyaç duyuyor?
- Hangi iş süreçlerini otomatikleştirmek istiyoruz?
- Uygulamanın hangi bileşenleri var?

### 1.2. Servislerin Tanımlanması

Gereksinimlerinizi belirledikten sonra, mikro servislerinizi tasarlayın. Her mikro hizmet, belirli bir işlevi yerine getirmelidir. Örneğin, bir e-ticaret uygulaması için aşağıdaki mikro hizmetler tanımlanabilir:

- **Kullanıcı Servisi**: Kullanıcı kaydı, oturum açma ve profil yönetimi.
- **Ürün Servisi**: Ürün ekleme, güncelleme ve silme.
- **Sipariş Servisi**: Kullanıcı siparişlerini yönetme.
- **Ödeme Servisi**: Ödeme işlemlerini gerçekleştirme.

## 2. Teknoloji Seçimi

Mikro servis mimarisi geliştirmede kullanacağınız teknolojileri seçin. Bu aşamada dikkate almanız gereken noktalar:

- **Programlama Dili**: Go, Python, Java gibi dillerden birini seçebilirsiniz. Bu örnekte Go dilini kullanacağız.
- **Veritabanı**: Her mikro hizmetin kendi veritabanına sahip olması önerilir. MySQL, PostgreSQL veya MongoDB gibi veritabanlarını değerlendirebilirsiniz.
- **API İletişimi**: REST API veya gRPC gibi iletişim protokollerini kullanabilirsiniz.

## 3. Mikro Servis Geliştirme

### 3.1. Proje Yapılandırması

Her mikro hizmet için ayrı bir proje yapısı oluşturun. Aşağıda, basit bir "Kullanıcı Servisi" mikro hizmeti için bir yapı gösterilmektedir:

```plaintext
user-service/
│
├── main.go           // Ana uygulama dosyası
├── user.go           // Kullanıcı model dosyası
├── handler.go        // HTTP handler'ları
├── repository.go     // Veritabanı işlemleri
└── go.mod            // Go modül dosyası
```

### 3.2. Kullanıcı Modeli Tanımlama

```go
// user.go
package main

// Kullanıcı modelini tanımlama
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### 3.3. Veritabanı İşlemleri

Veritabanına bağlantı ve CRUD işlemleri için bir `repository.go` dosyası oluşturun.

```go
// repository.go
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL sürücüsü
)

var db *sql.DB

// Veritabanına bağlantı kurma
func initDB() {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        panic(err) // Hata durumunda panik yap
    }
}

// Kullanıcı ekleme
func createUser(user User) (int64, error) {
    result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId() // Eklenen kullanıcının ID'sini döndür
}
```

### 3.4. HTTP Handler’ları

HTTP isteklerini yönlendirmek için bir `handler.go` dosyası oluşturun.

```go
// handler.go
package main

import (
    "encoding/json"
    "net/http"
)

// Kullanıcı oluşturma endpoint'i
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    id, err := createUser(user) // Kullanıcıyı veritabanına ekle
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int64{"id": id}) // Kullanıcı ID'sini döndür
}
```

### 3.5. Ana Fonksiyonu Tamamlama

Ana fonksiyonu tamamlayarak mikro hizmetinizi başlatın.

```go
// main.go
package main

import (
    "net/http"
)

func main() {
    initDB() // Veritabanına bağlantı kur
    defer db.Close() // Uygulama sona erdiğinde veritabanı bağlantısını kapat

    http.HandleFunc("/users", createUserHandler) // /users endpoint'ine yönlendirme
    http.ListenAndServe(":8080", nil) // 8080 portunda sunucu başlatma
}
```

## 4. Test Etme

Mikro hizmetinizi test etmek için bir HTTP istemcisi (örneğin, Postman veya cURL) kullanabilirsiniz.

### 4.1. Kullanıcı Oluşturma

Yeni bir kullanıcı oluşturmak için aşağıdaki cURL komutunu çalıştırabilirsiniz:

```bash
curl -X POST http://localhost:8080/users -d '{"name": "Ali", "email": "ali@example.com"}' -H "Content-Type: application/json"
```

### 4.2. Çıktı Açıklaması

Başarılı bir istek sonucunda alacağınız çıktı:

```json
{
    "id": 1
}
```

Bu çıktı, başarıyla eklenen kullanıcının ID'sini gösterir.

## 5. Dağıtım ve Yönetim

Mikro hizmetlerinizi dağıtmak için Docker kullanabilirsiniz. Her mikro hizmeti bir Docker konteyneri olarak paketleyerek, kolayca dağıtım ve ölçeklendirme sağlayabilirsiniz.

### 5.1. Dockerfile Oluşturma

Mikro hizmetiniz için bir `Dockerfile` oluşturun.

```dockerfile
# Dockerfile
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o user-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
CMD ["./user-service"]
```

### 5.2. Docker İmajını Oluşturma

Aşağıdaki komutla Docker imajını oluşturun:

```bash
docker build -t user-service .
```

### 5.3. Docker Konteynerini Çalıştırma

Aşağıdaki komutla Docker konteynerini çalıştırın:

```bash
docker run -p 8080:8080 user-service
```

## 6. İzleme ve Hata Yönetimi

Mikro hizmetlerinizi izlemek için Prometheus ve Grafana gibi araçları kullanabilirsiniz. Bu araçlar, hizmetlerinizi performans metrikleri ile izlemenizi sağlar.

### 6.1. Hata Yönetimi

Mikro hizmetlerde hata yönetimi, uygulamanızın sağlamlığı açısından kritik öneme sahiptir. Her mikro hizmette hata durumlarını yönetmek için uygun hata mesajları ve durum kodları kullanmalısınız.

## Sonuç

Mikro servis mimarisi geliştirme süreci, iyi bir planlama ve tasarım ile başlar. Her mikro hizmetin belirli bir işlevi yerine getirmesi sağlanır. Proje yönetimi, teknoloji seçimi, geliştirme, test etme ve dağıtım adımları, mikro hizmetlerin başarılı bir şekilde yönetilmesi için önemlidir. İyi bir izleme ve hata yönetimi stratejisi ile mikro hizmet mimarisi uygulamanızın başarısını artırabilirsiniz.