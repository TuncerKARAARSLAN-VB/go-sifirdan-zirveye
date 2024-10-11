Docker Compose, birden fazla Docker konteynerini tanımlamak ve yönetmek için kullanılan bir araçtır. Go uygulamanız ile bir veritabanını (örneğin MySQL veya PostgreSQL) aynı anda çalıştırmak için Docker Compose kullanarak yapılandırma yapabilirsiniz. Bu, uygulamanızı ve veritabanınızı kolayca başlatmak, durdurmak ve yönetmek için mükemmel bir yoldur.

## 1. Docker Compose Nedir?

- **Amaç**: Birden fazla konteyneri bir arada yönetmek ve yapılandırmak.
- **Yapılandırma Dosyası**: `docker-compose.yml` dosyası, hizmetleri tanımlamak için kullanılır.
- **Hızlı Başlatma**: Tek bir komut ile birden fazla konteyneri başlatma veya durdurma imkanı sağlar.

## 2. Proje Yapısı

Proje yapımız aşağıdaki gibi olacak:

```
myapp/
│
├── docker-compose.yml
├── main.go
└── go.mod
```

### 2.1. Go Uygulaması Yazma

İlk olarak, basit bir Go uygulaması oluşturalım. Bu uygulama, MySQL veritabanı ile bağlantı kuracak ve veritabanına veri ekleyecektir.

#### main.go

```go
// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql" // MySQL sürücüsünü içe aktarıyoruz
)

func main() {
    // Veritabanı bağlantısını açıyoruz
    db, err := sql.Open("mysql", "user:password@tcp(db:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Veritabanı bağlantısını test ediyoruz
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Veritabanına bağlanıldı!")

    // Basit bir HTTP sunucusu oluşturuyoruz
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        _, err := db.Exec("INSERT INTO greetings (message) VALUES ('Merhaba, Docker Compose ile entegre!')")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Mesaj veritabanına eklendi!")
    })

    // Sunucuyu başlatıyoruz
    fmt.Println("Sunucu 8080 portunda çalışıyor...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```

### Açıklamalar

- **Veritabanı Bağlantısı**: `sql.Open` ile MySQL veritabanına bağlanıyoruz. Bağlantı dizesinde `user`, `password`, `db` gibi yer tutucuları kullanıyoruz. `db` konteynerin adıdır.
- **HTTP Sunucusu**: Basit bir HTTP sunucusu oluşturuyoruz. `/` yoluna bir istek geldiğinde veritabanına bir kayıt ekliyoruz.

### 2.2. Go Modül Dosyası

Go uygulamamızı çalıştırabilmek için bir modül dosyası oluşturalım. Aşağıdaki komut ile `go.mod` dosyasını oluşturabilirsiniz:

```bash
go mod init myapp
```

Ardından `go-sql-driver/mysql` bağımlılığını eklemek için:

```bash
go get -u github.com/go-sql-driver/mysql
```

### 2.3. Docker Compose Dosyası

Şimdi, `docker-compose.yml` dosyasını oluşturalım. Bu dosya, Go uygulamamızı ve MySQL veritabanını yapılandırmak için kullanacağız.

#### docker-compose.yml

```yaml
version: '3.8'

services:
  db:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: mydb
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    ports:
      - "3306:3306"

  app:
    build: .
    restart: always
    ports:
      - "8080:8080"
    depends_on:
      - db
```

### Açıklamalar

- **Version**: Kullanılan Docker Compose sürümünü belirtir.
- **Services**: Uygulama bileşenlerini tanımlar.
  - **db**: MySQL veritabanı servisi.
    - `image`: Kullanılacak MySQL imajı.
    - `restart`: Her zaman yeniden başlatılması gerektiğini belirtir.
    - `environment`: MySQL için gerekli ortam değişkenleri (şifreler ve veritabanı adı).
    - `ports`: Konteynerin portunu dışarıya açar.
  - **app**: Go uygulaması servisi.
    - `build`: Uygulamanın bulunduğu dizini belirtir.
    - `ports`: Go uygulamasının dinleyeceği port.
    - `depends_on`: `app` servisi, `db` servisine bağlıdır.

### 2.4. Dockerfile Oluşturma

Go uygulamamızı derlemek için bir Dockerfile oluşturalım.

#### Dockerfile

```dockerfile
# Golang imajını kullanarak uygulama derleme
FROM golang:1.19 AS builder

# Çalışma dizini
WORKDIR /app

# Go modül dosyalarını kopyala
COPY go.mod ./
COPY go.sum ./

# Modülleri indir
RUN go mod download

# Uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle
RUN go build -o myapp

# Minimal bir imaj ile çalıştırma
FROM alpine:latest

# Çalışma dizini
WORKDIR /app

# Derlenmiş uygulamayı kopyala
COPY --from=builder /app/myapp .

# Konteynerin dinleyeceği portu belirt
EXPOSE 8080

# Uygulamayı başlat
CMD ["./myapp"]
```

### Açıklamalar

- **Multi-Stage Yapı**: Uygulama derlemesi için `golang` imajı kullanılıyor, ardından minimal `alpine` imajı ile çalıştırılıyor.
- **Derleme ve Çalıştırma**: Uygulama derlendikten sonra çalıştırılacak dosya `myapp` olarak kopyalanıyor.

### 2.5. Uygulamanın Çalıştırılması

Şimdi, uygulamayı başlatmak için terminalde projenin bulunduğu dizine gidin ve aşağıdaki komutu çalıştırın:

```bash
docker-compose up --build
```

- `--build`: Dockerfile'ı kullanarak imajları yeniden oluşturur.

### 2.6. Uygulamayı Test Etme

Başlatma işlemi tamamlandığında, tarayıcınızı açarak `http://localhost:8080` adresine gidin. Aşağıdaki mesajı görmelisiniz:

```plaintext
Mesaj veritabanına eklendi!
```

### 2.7. Veritabanındaki Verileri Kontrol Etme

MySQL veritabanında eklediğimiz verileri kontrol etmek için aşağıdaki komut ile MySQL CLI'ye bağlanabilirsiniz:

```bash
docker exec -it myapp_db_1 mysql -u user -p
```

- `myapp_db_1`: Docker Compose tarafından oluşturulan MySQL konteynerinin adı.
- `-u user`: MySQL kullanıcı adı.
- `-p`: Parola isteyecek. Parola olarak `password` yazın.

Veritabanına bağlandıktan sonra aşağıdaki SQL sorgusu ile `greetings` tablosundaki verileri kontrol edebilirsiniz:

```sql
USE mydb;
SELECT * FROM greetings;
```

### 3. Sonuç

Bu makalede, Docker Compose kullanarak bir Go uygulaması ve MySQL veritabanı entegre ettik. Uygulamamız, HTTP istekleri alarak veritabanına veri ekleyebilme yeteneğine sahip. Docker Compose sayesinde uygulamanızı ve veritabanınızı hızlıca başlatıp yönetebilirsiniz. Docker ile geliştirme sürecinizi daha verimli hale getirebilir ve uygulamalarınızı kolayca dağıtabilirsiniz.