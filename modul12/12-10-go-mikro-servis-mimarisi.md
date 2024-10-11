## Microservices Mimarisi Nedir?

Microservices, bir uygulamanın işlevselliğini küçük, bağımsız hizmetler (servisler) olarak yapılandıran bir yazılım mimarisi modelidir. Her mikro hizmet, belirli bir işlevi yerine getirir ve genellikle kendi veritabanına, iş mantığına ve arayüzüne sahiptir. Microservices mimarisi, büyük ve karmaşık uygulamaların daha yönetilebilir ve esnek bir şekilde geliştirilmesini, dağıtılmasını ve bakımını sağlar.

### 1. Microservices'ın Avantajları

- **Bağımsız Geliştirme**: Her mikro hizmet, bağımsız olarak geliştirilip dağıtılabilir. Bu, geliştirme süreçlerini hızlandırır.
- **Teknoloji Seçimi**: Her hizmet, farklı teknolojiler ve programlama dilleri kullanılarak geliştirilebilir.
- **Esneklik**: Mikro hizmetlerin bağımsız yapısı, uygulamanın belirli bölümlerini değiştirmeyi veya yeniden yapılandırmayı kolaylaştırır.
- **Hata İzolasyonu**: Bir mikro hizmette meydana gelen hata, diğer hizmetleri etkilemez, bu da sistemin genel sağlamlığını artırır.

### 2. Microservices'ın Dezavantajları

- **Dağıtık Sistem Karmaşıklığı**: Birden fazla hizmetin yönetimi ve izlenmesi, bir monolitik uygulamaya göre daha karmaşık olabilir.
- **Ağ Gecikmesi**: Hizmetler arası iletişim genellikle ağ üzerinden yapılır, bu da gecikmelere yol açabilir.
- **Veri Yönetimi**: Her mikro hizmetin kendi veritabanına sahip olması, veri tutarlılığı sorunlarına yol açabilir.

## 3. Microservices Mimarisi Bileşenleri

### 3.1. API Gateway

API Gateway, tüm mikro hizmetlerin dış dünyaya açılan kapısıdır. İstemciler, API Gateway üzerinden mikro hizmetlere erişir. API Gateway, yönlendirme, kimlik doğrulama, yük dengeleme ve izleme gibi görevleri yerine getirir.

### 3.2. Service Discovery

Mikro hizmetlerin dinamik olarak bulunduğu bir sistemdir. Bir mikro hizmet, diğer mikro hizmetlerin adreslerini bulmak için bir Service Discovery aracına ihtiyaç duyar. Bu, hizmetlerin adreslerini merkezi bir yerde tutmayı sağlar.

### 3.3. Database per Service

Her mikro hizmet, kendi veritabanına sahip olmalıdır. Bu, hizmetlerin bağımsız çalışmasını ve veri modelinin her hizmet için özelleştirilmesini sağlar.

## 4. Microservices ile Uygulama Geliştirme

Aşağıda, Go dilinde bir mikro hizmet uygulaması geliştirmek için temel bir örnek verilmiştir. Bu örnek, basit bir kullanıcı yönetim sistemi içerecektir.

### 4.1. Kullanıcı Mikro Hizmeti

**Kullanıcı Mikro Hizmeti**: Kullanıcı bilgilerini yöneten bir mikro hizmet.

#### 4.1.1. Kullanıcı Mikro Hizmetini Oluşturma

```go
// user_service.go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Basit bir kullanıcı veritabanı
var users = []User{
    {ID: 1, Name: "Ali", Email: "ali@example.com"},
    {ID: 2, Name: "Ayşe", Email: "ayse@example.com"},
}

// Kullanıcıları listeleme
func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Ana fonksiyon
func main() {
    http.HandleFunc("/users", getUsers) // /users endpoint'ine yönlendirme
    http.ListenAndServe(":8080", nil)   // 8080 portunda sunucu başlatma
}
```

### 4.1.2. Çalıştırma

Yukarıdaki kodu `user_service.go` adlı bir dosyaya kaydedin ve terminalden çalıştırın:

```bash
go run user_service.go
```

### 4.1.3. Kullanıcı Bilgilerine Erişim

Terminalde, kullanıcı bilgilerine erişmek için aşağıdaki komutu kullanabilirsiniz:

```bash
curl http://localhost:8080/users
```

### 4.1.4. Çıktı Açıklaması

```
[
    {"id":1,"name":"Ali","email":"ali@example.com"},
    {"id":2,"name":"Ayşe","email":"ayse@example.com"}
]
```

Bu çıktı, iki kullanıcının bilgilerini JSON formatında gösterir.

## 5. Diğer Mikro Hizmetleri Oluşturma

Mikro hizmetlerinizi ihtiyacınıza göre genişletebilir ve diğer hizmetlerle entegre edebilirsiniz. Örneğin, bir **Sipariş Mikro Hizmeti** oluşturabilir ve bu hizmetin kullanıcı mikro hizmeti ile etkileşimde bulunmasını sağlayabilirsiniz.

## 6. Microservices ile İletişim

### 6.1. HTTP İletişimi

Mikro hizmetler, HTTP üzerinden REST API’leri ile birbirleriyle iletişim kurabilir.

### 6.2. Mesajlaşma Sistemleri

Mikro hizmetler, RabbitMQ, Kafka gibi mesajlaşma sistemleri kullanarak asenkron olarak da iletişim kurabilir.

## 7. Microservices Mimarisi Uygulama Örnekleri

### 7.1. E-Ticaret Uygulaması

Bir e-ticaret uygulaması aşağıdaki mikro hizmetleri içerebilir:

- **Kullanıcı Hizmeti**: Kullanıcı bilgilerini yönetir.
- **Ürün Hizmeti**: Ürün bilgilerini yönetir.
- **Sipariş Hizmeti**: Kullanıcı siparişlerini yönetir.
- **Ödeme Hizmeti**: Ödeme işlemlerini gerçekleştirir.

### 7.2. Sosyal Medya Uygulaması

Bir sosyal medya uygulaması da aşağıdaki mikro hizmetleri içerebilir:

- **Kullanıcı Profili Hizmeti**: Kullanıcı profillerini yönetir.
- **Gönderi Hizmeti**: Kullanıcı gönderilerini yönetir.
- **Yorum Hizmeti**: Gönderilere yapılan yorumları yönetir.
- **Bildirim Hizmeti**: Kullanıcılara bildirim gönderir.

## 8. Microservices Mimarisi için En İyi Uygulamalar

- **Küçük ve Bağımsız Hizmetler**: Her mikro hizmet, belirli bir işlevi yerine getirmelidir.
- **Hizmet Sınırları**: Mikro hizmetler arası iletişim net bir şekilde tanımlanmalıdır.
- **Versiyon Yönetimi**: Mikro hizmetlerin versiyonları iyi yönetilmelidir.
- **Monitoring ve Logging**: Mikro hizmetlerin performansı izlenmeli ve loglama yapılmalıdır.

## Sonuç

Microservices mimarisi, büyük ve karmaşık uygulamaları yönetilebilir küçük parçalara ayırarak geliştirmenize olanak tanır. Her mikro hizmetin bağımsız olarak geliştirilmesi, test edilmesi ve dağıtılması, yazılım geliştirme süreçlerini hızlandırır ve esnekliği artırır. Ancak, bu mimarinin karmaşıklıkları da göz önünde bulundurulmalı ve iyi bir planlama yapılmalıdır.