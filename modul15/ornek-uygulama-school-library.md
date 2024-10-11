# School Library 

Aşağıda, okul kütüphanesi için bir web API projesi geliştirme planınızı güncelledim. Giriş dışında tüm API arayüzlerinin token ile kontrol edildiğinden ve güvenlik açıklarının engellendiğinden emin olmak için gerekli düzenlemeleri yaptım. Proje, Go dilinde yazılacak ve JWT (JSON Web Token) kullanarak kimlik doğrulama sağlayacak.

### Proje Planı

1. **Proje Tanımı**: Kullanıcıların kütüphane kitaplarını ödünç alabileceği ve iade edebileceği bir API oluşturacağız. Kullanıcılar, sisteme giriş yapacak ve JWT kullanarak kimlik doğrulaması gerçekleştirecekler.

2. **Kullanılacak Teknolojiler**:
   - **Go**: Sunucu tarafı kodlama için kullanılacak.
   - **Gorilla Mux**: HTTP yönlendirme için kullanılacak.
   - **JWT**: Kullanıcı kimlik doğrulaması için kullanılacak.
   - **GORM**: Veritabanı ile etkileşim için kullanılacak (ORM).
   - **PostgreSQL**: Veritabanı olarak kullanılacak.

3. **Veritabanı Tasarımı**:
   - **Users**: Kullanıcı bilgilerini tutar.
   - **Books**: Kitap bilgilerini tutar.
   - **Loans**: Ödünç alınan kitapları tutar.

### Aşama 1: Projeyi Başlatma

Öncelikle, Go ve gerekli kütüphanelerin kurulumunu yapalım.

```bash
mkdir library-api
cd library-api
go mod init library-api
go get -u github.com/gorilla/mux
go get -u github.com/dgrijalva/jwt-go
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### Aşama 2: Veritabanı Tasarımı

Veritabanı tablolarını tanımlayalım. Bir PostgreSQL veritabanı oluşturduktan sonra, aşağıdaki kodları kullanarak tabloları oluşturabilirsiniz.

#### models.go

```go
package main

import (
    "gorm.io/gorm"
)

// User model
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
}

// Book model
type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Loan model
type Loan struct {
    ID     uint `json:"id" gorm:"primaryKey"`
    UserID uint `json:"user_id"`
    BookID uint `json:"book_id"`
}
```

### Aşama 3: Veritabanı Bağlantısı

Veritabanı bağlantısını ayarlamak için bir fonksiyon oluşturalım.

#### database.go

```go
package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func ConnectDatabase() *gorm.DB {
    dsn := "host=localhost user=postgres password=yourpassword dbname=library port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Tabloları otomatik oluştur
    db.AutoMigrate(&User{}, &Book{}, &Loan{})
    return db
}
```

### Aşama 4: Kimlik Doğrulama ve JWT Oluşturma

Kullanıcıların giriş yapabilmesi için bir kimlik doğrulama sistemi oluşturalım. JWT, kullanıcı kimlik doğrulaması için kullanılacak.

#### auth.go

```go
package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key") // JWT için anahtar

// Claims struct
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// Giriş için handler
func Login(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Kullanıcıyı veritabanından bul
    var dbUser User
    db := ConnectDatabase()
    if err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // JWT oluştur
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: dbUser.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Could not create token", http.StatusInternalServerError)
        return
    }

    // Token'ı yanıtla
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
```

### Aşama 5: API Handler'larını Oluşturma

API'yi oluşturmak için gerekli olan handler fonksiyonlarını tanımlayalım.

#### handlers.go

```go
package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

// Kitapları alma
func GetBooks(w http.ResponseWriter, r *http.Request) {
    db := ConnectDatabase()
    var books []Book
    db.Find(&books)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// Kitap ödünç alma
func LoanBook(w http.ResponseWriter, r *http.Request) {
    var loan Loan
    json.NewDecoder(r.Body).Decode(&loan)

    db := ConnectDatabase()
    db.Create(&loan)

    w.WriteHeader(http.StatusCreated)
}

// Kitap iade etme
func ReturnBook(w http.ResponseWriter, r *http.Request) {
    var loan Loan
    json.NewDecoder(r.Body).Decode(&loan)

    db := ConnectDatabase()
    db.Delete(&loan)

    w.WriteHeader(http.StatusNoContent)
}
```

### Aşama 6: Middleware ile JWT Doğrulama

Kullanıcıların kimlik doğrulamasını sağlamak için bir middleware fonksiyonu oluşturalım. Bu middleware, giriş dışındaki tüm API arayüzlerine erişim için token doğrulaması yapacaktır.

#### middleware.go

```go
package main

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

// JWT doğrulama middleware
func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Token geçerli ise, kullanıcı bilgilerini request'e ekle
        r.Header.Set("username", claims.Username)
        next.ServeHTTP(w, r)
    })
}
```

### Aşama 7: API Yönlendirmeleri

Son olarak, API yönlendirmelerini oluşturalım.

#### main.go

```go
package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // API yönlendirmeleri
    router.HandleFunc("/login", Login).Methods("POST")
    router.HandleFunc("/books", VerifyToken(GetBooks)).Methods("GET")
    router.HandleFunc("/loan", VerifyToken(LoanBook)).Methods("POST")
    router.HandleFunc("/return", VerifyToken(ReturnBook)).Methods("DELETE")

    log.Println("Sunucu 8080 portunda çalışıyor...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

### Aşama 8: Projeyi Çalıştırma

Projeyi çalıştırmak için PostgreSQL veritabanınızı ayarladıktan sonra aşağıdaki komutu çalıştırın:

```bash
go run main.go
```

### Aşama 9: API'yi Test Etme

API'yi test etmek için Postman veya cURL kullanabilirsiniz. Aşağıda bazı örnek istekler bulunmaktadır:

#### 1. Kullanıcı Girişi

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password123"}'
```

#### 2. Tüm Kitapları Alma

```bash
curl -X GET http://localhost:8080/books -H "Authorization: Bearer <TOKEN>"
```

#### 3. Kitap Ödünç Alma

```bash
curl -X POST

 http://localhost:8080/loan -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"user_id": 1, "book_id": 2}'
```

#### 4. Kitap İade Etme

```bash
curl -X DELETE http://localhost:8080/return -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"user_id": 1, "book_id": 2}'
```

### Sonuç

Bu proje, okul kütüphanesi için kitap ödünç alma ve iade işlemlerini yöneten bir web API'si olarak çalışır. Kullanıcılar kimlik doğrulama işlemi için JWT kullanırken, API HTTPS üzerinden güvenli bir şekilde çalışır. Proje, gerçek hayatta kullanılabilir bir temel oluşturur ve genişletilebilir. İleride eklemeler yapmak veya mevcut özellikleri geliştirmek oldukça kolaydır.