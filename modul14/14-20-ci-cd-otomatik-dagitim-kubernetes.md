CI/CD otomatik dağıtım sürecinde Kubernetes, uygulamaları konteynerleştirmek ve yönetmek için popüler bir platformdur. Kubernetes, dağıtım, ölçeklendirme ve uygulama yönetimini kolaylaştırırken, CI/CD süreçleri ile entegrasyonu otomatikleştirir. Bu yazıda, CI/CD süreçlerini kullanarak Kubernetes'e otomatik dağıtım sürecini detaylı bir şekilde anlatacağım.

### 1. CI/CD Nedir?

- **Continuous Integration (CI)**: Geliştiricilerin kod değişikliklerini sık sık entegre etme uygulamasıdır. Her değişiklik otomatik testlerle doğrulanır.
- **Continuous Deployment (CD)**: Başarılı testlerden sonra kodun otomatik olarak üretim ortamına dağıtılmasıdır.

### 2. Kubernetes Nedir?

Kubernetes, konteynerleştirilmiş uygulamaları otomatikleştirmek, dağıtmak, ölçeklendirmek ve yönetmek için kullanılan bir açık kaynak platformdur. Aşağıdaki temel bileşenleri içerir:

- **Pod**: En küçük dağıtım birimi. Bir veya daha fazla konteyneri içerebilir.
- **Deployment**: Pod'ların yönetimi için kullanılan bir kaynak türüdür. Uygulamaların güncellenmesi ve ölçeklendirilmesi için kullanılır.
- **Service**: Pod'lara erişimi sağlayan bir soyutlama katmanıdır.

### 3. Proje Yapısı

Aşağıdaki gibi bir proje yapısına sahip olduğumuzu varsayıyoruz:

```
myapp/
├── .github/
│   └── workflows/
│       └── ci-cd.yml
├── k8s/
│   └── deployment.yaml
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
    fmt.Fprintf(w, "Merhaba, Kubernetes CI/CD!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Sunucu 8080 portunda çalışıyor...")
    http.ListenAndServe(":8080", nil)
}
```

### Açıklamalar

- **HTTP Sunucusu**: `/` yoluna gelen isteklere "Merhaba, Kubernetes CI/CD!" yanıtını verir.
- **Port**: Uygulama, 8080 portunda dinler.

### 5. Kubernetes Deployment Tanımlama

Kubernetes'te uygulamayı dağıtmak için bir `deployment.yaml` dosyası oluşturalım.

#### k8s/deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-deployment
spec:
  replicas: 3 # 3 kopya oluştur
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp-container
        image: myapp:latest # Uygulama konteynerinin ismi
        ports:
        - containerPort: 8080 # Konteynerin dinlediği port
```

### Açıklamalar

- **replicas**: Uygulamanın kaç kopyasının çalışacağını belirtir.
- **selector**: Pod'ları tanımlamak için kullanılan etiketlerdir.
- **template**: Pod tanımını içerir.
- **containers**: Konteyner yapılandırmasını içerir.

### 6. GitHub Actions ile CI/CD Workflow Oluşturma

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

      - name: Build Docker Image
        run: |
          echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
          docker build -t myapp:latest . # Docker imajını oluştur

      - name: Push Docker Image
        run: docker push myapp:latest # Docker imajını Docker Hub'a gönder

      - name: Deploy to Kubernetes
        run: |
          kubectl apply -f k8s/deployment.yaml # Kubernetes'e deployment dosyasını uygula
```

### Açıklamalar

- **Docker İmajı Oluşturma**: `docker build` komutuyla Docker imajı oluşturulur.
- **Docker İmajını Gönderme**: `docker push` komutuyla Docker Hub'a gönderilir.
- **Kubernetes'e Dağıtım**: `kubectl apply` komutuyla Kubernetes ortamına dağıtım yapılır.

### 7. Ortam Değişkenlerini Ayarlama

Kubernetes'e erişim sağlamak için bazı ortam değişkenlerini ayarlamak gerekebilir. GitHub'da depo ayarlarına gidin ve aşağıdaki ortam değişkenlerini ekleyin:

- `DOCKER_USERNAME`: Docker Hub kullanıcı adı.
- `DOCKER_PASSWORD`: Docker Hub şifre.
- `KUBE_CONFIG`: Kubernetes erişim bilgileri (Kubeconfig dosyası).

### 8. Uygulamanın Test Edilmesi

Yapılandırmamız tamamlandığında, aşağıdaki adımları izleyin:

1. **Değişiklik Yapma**: `main.go` dosyasında basit bir değişiklik yapın.
2. **Commit ve Push**: Değişikliklerinizi `main` branşına gönderin.

```bash
git add .
git commit -m "Update main.go"
git push origin main
```

### 9. Sonuç

- **GitHub Actions**: Herhangi bir push veya pull request tetiklendiğinde CI/CD süreciniz otomatik olarak başlayacak.
- **Bağımlılıkların Yüklenmesi**: `go mod tidy` ile bağımlılıklar yüklenir.
- **Uygulamanın Derlenmesi**: `go build` komutuyla uygulama derlenir.
- **Docker İmajı Oluşturma ve Gönderme**: Docker imajı oluşturulup Docker Hub'a gönderilir.
- **Kubernetes'e Dağıtım**: Başarılı bir şekilde, uygulama Kubernetes ortamına dağıtılır.

Bu süreç, yazılım geliştirme sürecinizde verimliliği artırır ve daha hızlı bir geliştirme döngüsü sağlar. CI/CD süreçlerinin Kubernetes ile entegrasyonu, uygulamanızın güvenilir ve ölçeklenebilir bir şekilde dağıtılmasını sağlar.