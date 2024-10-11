CI/CD (Continuous Integration / Continuous Deployment), yazılım geliştirme süreçlerini otomatikleştiren bir yöntemdir. GitHub, bu süreçleri desteklemek için GitHub Actions adında bir hizmet sunar. Bu yazıda, GitHub Actions kullanarak bir CI/CD otomatik dağıtım sürecini nasıl oluşturabileceğinizi detaylı olarak anlatacağım. 

### 1. CI/CD Nedir?

- **Continuous Integration (CI)**: Geliştiricilerin kod değişikliklerini sık sık birleştirdiği bir uygulamadır. Her bir değişiklik, otomatik testler ve yapı işlemleri ile doğrulanır.
- **Continuous Deployment (CD)**: Testler başarıyla tamamlandığında, kod değişiklikleri otomatik olarak üretim ortamına dağıtılır.

### 2. GitHub Actions Nedir?

- **Hizmet**: GitHub'ın sunduğu CI/CD otomasyonu için bir platformdur.
- **Yapılandırma**: `.github/workflows` klasöründe YAML dosyaları ile tanımlanır.
- **Olay Tabanlı**: Belirli olaylara (commit, pull request, vb.) göre tetiklenebilir.

### 3. Proje Yapısı

Aşağıdaki gibi bir proje yapısına sahip olduğumuzu varsayıyoruz:

```
myapp/
├── .github/
│   └── workflows/
│       └── ci-cd.yml
├── main.go
└── go.mod
```

### 4. Go Uygulaması Yazma

Basit bir Go uygulaması oluşturalım. Bu uygulama, bir HTTP sunucusu çalıştıracak.

#### main.go

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba, GitHub Actions CI/CD!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Sunucu 8080 portunda çalışıyor...")
    http.ListenAndServe(":8080", nil)
}
```

### Açıklamalar

- **HTTP Sunucusu**: `/` yoluna gelen isteklere "Merhaba, GitHub Actions CI/CD!" yanıtını verir.
- **Port**: Uygulama, 8080 portunda dinler.

### 5. GitHub Actions Workflow Oluşturma

Şimdi, CI/CD sürecini tanımlamak için bir GitHub Actions workflow dosyası oluşturalım.

#### .github/workflows/ci-cd.yml

```yaml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main # Ana branş üzerinde yapılan her push tetikleyecek
  pull_request:
    branches:
      - main # Ana branş üzerinde açılan her PR tetikleyecek

jobs:
  build:
    runs-on: ubuntu-latest # GitHub Actions çalıştırma ortamı
    steps:
      - name: Checkout repository
        uses: actions/checkout@v2 # Depoyu çek

      - name: Set up Go
        uses: actions/setup-go@v2 # Go ortamını kur
        with:
          go-version: '1.19' # Kullanılacak Go sürümü

      - name: Install dependencies
        run: go mod tidy # Bağımlılıkları yükle

      - name: Build
        run: go build -o myapp # Uygulamayı derle

      - name: Run tests
        run: go test ./... # Testleri çalıştır

      - name: Deploy
        run: echo "Deploying to production..." # Gerçek dağıtım komutları buraya yazılacak
```

### Açıklamalar

- **name**: Workflow'un adı.
- **on**: Hangi olayların workflow'u tetikleyeceğini belirler.
  - `push`: Ana branşa yapılan her push işlemi.
  - `pull_request`: Ana branşa açılan her PR.
- **jobs**: Her bir işin yapılandırması.
  - **build**: Derleme ve test adımlarını içerir.
    - `runs-on`: Kullanılacak işletim sistemi (örneğin `ubuntu-latest`).
    - **steps**: Her bir adım.
      - `actions/checkout`: Depoyu klonlamak için kullanılır.
      - `actions/setup-go`: Go ortamını kurar.
      - `go mod tidy`: Go bağımlılıklarını yükler.
      - `go build`: Uygulamayı derler.
      - `go test`: Testleri çalıştırır.
      - **Deploy**: Dağıtım komutları buraya yazılacak.

### 6. Dağıtım Sürecini Otomatikleştirme

Gerçek bir dağıtım süreci için, deployment işlemi için bir bulut sağlayıcısı (AWS, Heroku, DigitalOcean vb.) kullanabilirsiniz. Örneğin, Heroku'ya dağıtım yapmak için aşağıdaki adımları takip edebilirsiniz:

#### Deploy Adımını Güncelleme

```yaml
      - name: Deploy
        run: |
          curl -s https://cli-assets.heroku.com/install.sh | sh # Heroku CLI yükleme
          heroku git:remote -a your-heroku-app-name # Heroku uygulamanızın adı
          git add -A
          git commit -m "Deploy to Heroku" --allow-empty # Değişiklikleri ekle
          git push heroku main # Değişiklikleri Heroku'ya gönder
```

### Açıklamalar

- **Heroku CLI**: Heroku CLI'sini yükler.
- **git remote**: Heroku uygulamasına bağlanır.
- **git push**: Uygulamayı Heroku'ya gönderir.

### 7. Uygulamanın Test Edilmesi

Yapılandırmamız tamamlandığında, aşağıdaki adımları izleyin:

1. **Değişiklik Yapma**: `main.go` dosyasında basit bir değişiklik yapın.
2. **Commit ve Push**: Değişikliklerinizi `main` branşına gönderin.

```bash
git add .
git commit -m "Update main.go"
git push origin main
```

### 8. Sonuç

- **GitHub Actions**: Herhangi bir push veya pull request tetiklendiğinde CI/CD süreciniz otomatik olarak başlayacak.
- **Bağımlılıkların Yüklenmesi**: `go mod tidy` ile bağımlılıklar yüklenir.
- **Uygulamanın Derlenmesi**: `go build` komutuyla uygulama derlenir.
- **Testlerin Çalıştırılması**: `go test` komutuyla testler çalıştırılır.
- **Otomatik Dağıtım**: Herhangi bir hata yoksa, kodunuz otomatik olarak belirlediğiniz ortama dağıtılır.

Bu süreç, yazılım geliştirme sürecinizde verimliliği artırır ve daha hızlı bir geliştirme döngüsü sağlar. CI/CD, kod değişikliklerinizi hızlı ve güvenli bir şekilde üretim ortamına aktararak yazılım teslimatını kolaylaştırır.