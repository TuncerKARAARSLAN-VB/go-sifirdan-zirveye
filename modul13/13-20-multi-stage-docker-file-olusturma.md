Multi-stage Dockerfile, birden fazla aşama kullanarak Docker imajları oluşturmanızı sağlayan bir özelliktir. Bu, özellikle derleme aşamasında daha büyük bir temel imaj kullanmak yerine, yalnızca gerekli olan bileşenleri içeren daha küçük ve daha hafif bir imaj elde etmenize olanak tanır. Multi-stage yapılandırması, genellikle derleme sürecinde kullanılan geçici dosyaları ve bağımlılıkları temizlemeye yardımcı olur, bu da daha verimli ve taşınabilir Docker imajları oluşturur.

## 1. Multi-Stage Dockerfile Nedir?

- **Amaç**: Birden fazla aşama kullanarak, her aşamada gerekli bileşenleri oluşturup, en son aşamada sadece gereken dosyaları içeren minimal bir Docker imajı oluşturmak.
- **Avantajlar**:
  - Daha küçük imaj boyutları.
  - Daha az karmaşık bağımlılık yönetimi.
  - Daha hızlı dağıtım ve çalıştırma süreleri.

## 2. Multi-Stage Dockerfile Oluşturma

### 2.1. Örnek Uygulama: Go Uygulaması

Aşağıda, basit bir Go uygulaması ile multi-stage Dockerfile oluşturacağız. Uygulamamız, HTTP üzerinden gelen istekleri yanıtlayan basit bir sunucu olacaktır.

#### 2.1.1. Go Uygulaması Yazma

Öncelikle, aşağıdaki gibi basit bir Go uygulaması oluşturun:

#### main.go

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba, Multi-Stage Dockerfile ile kapsüllenmiş uygulama!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Sunucu 8080 portunda çalışıyor...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 2.2. Multi-Stage Dockerfile Oluşturma

Şimdi, yukarıda yazdığımız Go uygulamasını kapsüllemek için bir Multi-Stage Dockerfile oluşturalım.

#### Dockerfile

```dockerfile
# İlk aşama: Go uygulamasını derleme
FROM golang:1.19 AS builder

# Çalışma dizinini ayarlama
WORKDIR /app

# Go modül dosyalarını kopyalama
COPY go.mod ./
COPY go.sum ./

# Modülleri indirme
RUN go mod download

# Uygulamayı kopyalama
COPY . .

# Uygulamayı derleme
RUN go build -o myapp

# İkinci aşama: Minimal bir imaj ile çalıştırma
FROM alpine:latest

# Çalışma dizinini ayarlama
WORKDIR /app

# Derlenmiş uygulamayı kopyalama
COPY --from=builder /app/myapp .

# Konteynerin çalıştırılacağı komut
CMD ["./myapp"]

# Konteynerin dinleyeceği port
EXPOSE 8080
```

### Açıklamalar

1. **İlk Aşama (Builder)**:
   - `FROM golang:1.19 AS builder`: Go dilini kullanarak uygulamamızı derleyeceğimiz temel imajı belirtir. `AS builder` kısmı, bu aşamayı `builder` adıyla etiketler.
   - `WORKDIR /app`: Çalışma dizinini `/app` olarak ayarlıyoruz.
   - `COPY go.mod ./` ve `COPY go.sum ./`: Go modül dosyalarını kopyalıyoruz.
   - `RUN go mod download`: Modülleri indiriyoruz.
   - `COPY . .`: Uygulama dosyalarını konteyner içine kopyalıyoruz.
   - `RUN go build -o myapp`: Uygulamayı derleyerek `myapp` adlı bir yürütülebilir dosya oluşturuyoruz.

2. **İkinci Aşama**:
   - `FROM alpine:latest`: Çok hafif olan Alpine Linux imajını kullanarak minimal bir ortam oluşturuyoruz.
   - `WORKDIR /app`: Çalışma dizinini yine `/app` olarak ayarlıyoruz.
   - `COPY --from=builder /app/myapp .`: İlk aşamadan derlenmiş uygulamayı kopyalıyoruz. `--from=builder` ifadesi, uygulamanın hangi aşamadan kopyalanacağını belirtir.
   - `CMD ["./myapp"]`: Konteyner başlatıldığında çalıştırılacak komut.
   - `EXPOSE 8080`: Konteynerin dinleyeceği portu belirtir.

### 2.3. Docker İmajı Oluşturma

Terminalde projenizin bulunduğu dizine gidin ve aşağıdaki komutla Docker imajını oluşturun:

```bash
docker build -t myapp .
```

### 2.4. Docker Konteynerini Çalıştırma

Oluşturduğumuz Docker imajını çalıştırmak için aşağıdaki komutu kullanın:

```bash
docker run -d -p 8080:8080 myapp
```

### 2.5. Uygulamayı Test Etme

Tarayıcınızı açın ve `http://localhost:8080` adresine gidin. Aşağıdaki mesajı görmelisiniz:

```plaintext
Merhaba, Multi-Stage Dockerfile ile kapsüllenmiş uygulama!
```

## 3. Multi-Stage Dockerfile'ın Avantajları

- **Küçük İmaj Boyutları**: Sadece gerekli olan dosyaları içeren daha küçük imajlar elde edersiniz.
- **Hızlı Dağıtım**: Küçük imaj boyutları, daha hızlı dağıtım ve başlatma süreleri sağlar.
- **Temiz Bağımlılık Yönetimi**: İlk aşamada kullanılan bağımlılıklar, son aşamada kopyalanmadığı için temiz bir imaj oluşturur.

## 4. Sonuç

Bu makalede, Multi-Stage Dockerfile oluşturma konusunu detaylı bir şekilde ele aldık. Basit bir Go uygulaması ile bir Docker imajı oluşturarak, çok aşamalı bir yapı kullanmanın avantajlarını gösterdik. Multi-stage yapılandırması, uygulama geliştirme sürecinizi daha verimli hale getirirken, daha küçük ve taşınabilir imajlar elde etmenizi sağlar. Bu sayede, uygulamalarınızı daha hızlı bir şekilde geliştirebilir ve dağıtabilirsiniz.