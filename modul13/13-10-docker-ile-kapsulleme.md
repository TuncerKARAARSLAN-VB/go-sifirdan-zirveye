Docker, uygulamaların kapsüllenmesini ve taşınabilirliğini sağlamak için kullanılan bir platformdur. Uygulamaları ve bağımlılıklarını bir araya getirerek konteyner adı verilen hafif sanal ortamlarda çalıştırılmasını sağlar. Bu makalede, Docker ile kapsülleme konusunu detaylı bir şekilde inceleyeceğiz, örneklerle açıklamalar yapacağız ve çıktılarda yeterli açıklamalar ekleyeceğiz.

## 1. Docker Nedir?

- **Kapsülleme**: Docker, uygulamaları ve tüm bağımlılıklarını (kütüphaneler, ayarlar vb.) tek bir paket içinde tutarak taşınabilir hale getirir.
- **Konteyner**: Docker, uygulamanın çalışması için gereken her şeyi içeren bir konteyner oluşturur. Konteynerler, işletim sistemi kaynaklarını paylaşır ve birbirlerinden izole çalışır.
- **Taşınabilirlik**: Bir konteyner, geliştirme ortamında çalıştığı gibi üretim ortamında da çalışabilir. Bu, "çalışıyor ama neden çalışmıyor?" sorununu minimize eder.

## 2. Docker Kurulumu

Docker’ı kurmak için işletim sisteminize uygun olan sürümü [Docker resmi web sitesinden](https://docs.docker.com/get-docker/) indirebilirsiniz. Aşağıdaki adımlarla Docker'ı kurabilirsiniz:

### 2.1. Docker Kurulumu

1. **Docker'ı İndirin**: İşletim sisteminize uygun Docker sürümünü indirin.
2. **Kurulumu Tamamlayın**: İndirdiğiniz dosyayı çalıştırarak kurulum işlemini tamamlayın.
3. **Docker'ı Başlatın**: Kurulumdan sonra Docker servisini başlatın. Windows ve Mac kullanıcıları için, Docker Desktop uygulamasını açarak başlayabilirsiniz. Linux kullanıcıları için terminalden aşağıdaki komutla Docker'ı başlatın:

   ```bash
   sudo systemctl start docker
   ```

4. **Docker Kurulumunu Doğrulayın**: Terminalde aşağıdaki komutu çalıştırarak Docker'ın doğru kurulduğunu kontrol edin:

   ```bash
   docker --version
   ```

   Bu komut, Docker versiyonunu gösterecektir.

## 3. Docker ile Uygulama Kapsülleme

### 3.1. Basit Bir Uygulama Oluşturma

Bu bölümde, basit bir Go uygulaması oluşturacak ve bunu Docker ile kapsülleyeceğiz.

#### 3.1.1. Go Uygulaması Yazma

Aşağıdaki gibi basit bir Go uygulaması oluşturun:

#### main.go

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba, Docker ile Kapsüllenmiş Uygulama!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Sunucu 8080 portunda çalışıyor...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 3.2. Dockerfile Oluşturma

Docker, uygulamanızı kapsüllemek için bir yapılandırma dosyasına (Dockerfile) ihtiyaç duyar. Bu dosya, konteynerinizin nasıl oluşturulacağını belirtir.

#### Dockerfile

Aşağıdaki Dockerfile dosyasını oluşturun:

```dockerfile
# Temel imaj olarak Go'yu kullanıyoruz
FROM golang:1.19

# Çalışma dizinini ayarlama
WORKDIR /app

# Go modül dosyasını kopyalama
COPY go.mod ./
COPY go.sum ./

# Modülleri indirme
RUN go mod download

# Uygulamayı kopyalama
COPY . .

# Uygulamayı derleme
RUN go build -o myapp

# Konteynerin çalıştırılacağı komut
CMD ["./myapp"]

# Konteynerin dinleyeceği port
EXPOSE 8080
```

#### Açıklama:

- **FROM golang:1.19**: Bu satır, Go 1.19 sürümünü içeren bir temel imaj kullanır.
- **WORKDIR /app**: Çalışma dizinini `/app` olarak ayarlıyoruz. Bu dizin, uygulama dosyalarımızın bulunduğu yerdir.
- **COPY go.mod ./ ve COPY go.sum ./**: Go modül dosyalarını konteyner içine kopyalıyoruz. Bu, bağımlılıkları daha verimli bir şekilde yönetmemizi sağlar.
- **RUN go mod download**: Bağımlılıkları indiriyoruz.
- **COPY . .**: Proje dizinimizi konteynerin çalışma dizinine kopyalıyoruz.
- **RUN go build -o myapp**: Uygulamayı derliyoruz ve `myapp` adıyla bir yürütülebilir dosya oluşturuyoruz.
- **CMD ["./myapp"]**: Konteyner başlatıldığında çalıştırılacak komut.
- **EXPOSE 8080**: Konteynerin dinleyeceği portu belirtiyoruz.

### 3.3. Docker İmajı Oluşturma

Terminalde projenizin bulunduğu dizine gidin ve aşağıdaki komutla Docker imajını oluşturun:

```bash
docker build -t myapp .
```

Bu komut, bulunduğunuz dizindeki Dockerfile'ı kullanarak `myapp` adında bir Docker imajı oluşturur.

### 3.4. Docker Konteynerini Çalıştırma

Oluşturduğumuz Docker imajını çalıştırmak için aşağıdaki komutu kullanın:

```bash
docker run -d -p 8080:8080 myapp
```

#### Açıklamalar:

- **-d**: Konteyneri arka planda çalıştırır.
- **-p 8080:8080**: Konteynerin 8080 portunu yerel makinenizin 8080 portuna yönlendirir.

### 3.5. Uygulamayı Test Etme

Tarayıcınızı açın ve `http://localhost:8080` adresine gidin. Aşağıdaki mesajı görmelisiniz:

```plaintext
Merhaba, Docker ile Kapsüllenmiş Uygulama!
```

Bu mesaj, uygulamanın başarılı bir şekilde Docker konteynerinde çalıştığını gösterir.

## 4. Docker ile Kapsüllemenin Avantajları

- **Taşınabilirlik**: Uygulamanızı bir konteyner içinde paketlediğinizde, bu konteyneri her yerde çalıştırabilirsiniz.
- **Yalıtım**: Her konteyner, diğer konteynerlerden izole bir ortamda çalışır. Bu, farklı uygulamaların çakışmasını önler.
- **Yeniden Üretilebilirlik**: Uygulamanızı ve tüm bağımlılıklarını kapsüllemeniz, diğer geliştiricilerin veya ortamlardaki kurulumların tutarlı olmasını sağlar.

## 5. Sonuç

Bu makalede, Docker ile uygulama kapsülleme konusunu detaylı bir şekilde ele aldık. Basit bir Go uygulaması oluşturduk, bunu bir Docker konteynerinde kapsülledik ve çalıştırdık. Docker, modern yazılım geliştirme süreçlerinde önemli bir rol oynamakta ve uygulamaların daha hızlı, daha güvenilir bir şekilde dağıtımını sağlamaktadır. Docker'ı kullanarak uygulamalarınızı verimli bir şekilde geliştirebilir ve dağıtabilirsiniz.