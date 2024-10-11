# Stock App
 
Bu proje, bir kütüphane yönetim sistemi için kapsamlı bir Web API geliştirmek amacıyla tasarlanmıştır. Proje, kullanıcıların kitap satın alma, iade işlemleri, depo girişi, stok çıkışı ve hurdaya çıkarma gibi işlemleri gerçekleştirmesine olanak tanır. Projenin güvenliği için kullanıcı kimlik doğrulama işlemleri JWT (JSON Web Token) ile sağlanacak ve SSL ile güvenli iletişim sağlanacaktır. 

### Proje Planı

1. **Proje Tanımı**: Kullanıcılar, kitap satın alma ve iade işlemleri gerçekleştirebilir. Ayrıca, depo girişleri, stok çıkışları ve hurdaya çıkarma işlemleri yapılabilir. Kullanıcılar, sisteme giriş yaptıktan sonra JWT ile kimlik doğrulaması yaparlar.

2. **Kullanılacak Teknolojiler**:
   - **Go**: Sunucu tarafı kodlama için kullanılacak.
   - **Gorilla Mux**: HTTP yönlendirme için kullanılacak.
   - **JWT**: Kullanıcı kimlik doğrulaması için kullanılacak.
   - **GORM**: Veritabanı ile etkileşim için kullanılacak (ORM).
   - **PostgreSQL**: Veritabanı olarak kullanılacak.
   - **SSL**: API iletişimi için güvenli bir bağlantı sağlamak için kullanılacak.

3. **Veritabanı Tasarımı**:
   - **Users**: Kullanıcı bilgilerini tutar.
   - **Books**: Kitap bilgilerini tutar.
   - **Loans**: Ödünç alınan kitapları tutar.
   - **Purchases**: Satın alınan kitapları tutar.
   - **StockMovements**: Depoya giriş ve stok çıkış işlemlerini tutar.

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
    ID       uint   `json:"id" gorm:"primaryKey"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"` // Kitap miktarı
}

// Loan model
type Loan struct {
    ID     uint `json:"id" gorm:"primaryKey"`
    UserID uint `json:"user_id"`
    BookID uint `json:"book_id"`
}

// Purchase model
type Purchase struct {
    ID     uint `json:"id" gorm:"primaryKey"`
    BookID uint `json:"book_id"`
    Quantity int `json:"quantity"`
}

// StockMovement model
type StockMovement struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    BookID    uint   `json:"book_id"`
    MovementType string `json:"movement_type"` // "in" veya "out"
    Quantity  int    `json:"quantity"`
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
    db.AutoMigrate(&User{}, &Book{}, &Loan{}, &Purchase{}, &StockMovement{})
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
    expirationTime := time.Now().Add(30 * time.Minute) // Token süresi 30 dakika
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

// Kitap satın alma
func PurchaseBook(w http.ResponseWriter, r *http.Request) {
    var purchase Purchase
    json.NewDecoder(r.Body).Decode(&purchase)

    db := ConnectDatabase()
    db.Create(&purchase)

    // Stok güncelleme
    var book Book
    db.First(&book, purchase.BookID)
    book.Quantity += purchase.Quantity
    db.Save(&book)

    // Stok hareketi kaydı
    db.Create(&StockMovement{BookID: purchase.BookID, MovementType: "in", Quantity: purchase.Quantity})

    w.WriteHeader(http.StatusCreated)
}

// Kitap ödünç alma
func LoanBook(w http.ResponseWriter, r *http.Request) {
    var loan Loan
    json.NewDecoder(r.Body).Decode(&loan)

    db := ConnectDatabase()
    db.Create(&loan)

    // Stok güncelleme
    var book Book
    db.First(&book, loan.BookID)
    book.Quantity--
    db.Save(&book)

    // Stok hareketi kaydı
    db.Create(&StockMovement{BookID: loan.BookID, MovementType: "out", Quantity: 1})

    w.WriteHeader(http.StatusCreated)
}

// Kitap iade etme
func ReturnBook(w http.ResponseWriter, r *http.Request) {
    var loan Loan
    json.NewDecoder(r.Body).Decode(&loan)

    db := ConnectDatabase()
    db.Delete(&loan)

    // Stok güncelleme
    var book Book
    db.First(&book, loan.BookID)
    book.Quantity++
    db.Save(&book)

    // Stok hareketi kaydı
    db.Create(&StockMovement{BookID: loan.BookID, MovementType: "in", Quantity: 1})

    w.WriteHeader(http.StatusNoContent)
}

// Depo çıkışı işlemi
func StockOut(w http.ResponseWriter, r *http.Request) {
    var stockMovement StockMovement
    json.NewDecoder(r.Body).Decode(&stockMovement)

    db := ConnectDatabase()
    db.Create(&stockMovement)

    var book Book
    db.First(&book, stockMovement.BookID)
    book.Quantity -= stockMovement.Quantity
    db.Save(&book)

    // Stok hareketi kaydı
    db.Create(&StockMovement{BookID: stockMovement.BookID, MovementType: "out", Quantity: stockMovement.Quantity})

    w.WriteHeader(http.StatusCreated)
}

// Hurdaya çıkarma işlemi
func MarkAsScrap(w http.ResponseWriter, r *http.Request) {
    var

 book Book
    json.NewDecoder(r.Body).Decode(&book)

    db := ConnectDatabase()
    db.Delete(&book)

    // Hurdaya çıkarma kaydı
    db.Create(&StockMovement{BookID: book.ID, MovementType: "scrap", Quantity: 1})

    w.WriteHeader(http.StatusNoContent)
}
```

### Aşama 6: Middleware ile JWT Doğrulama

Kullanıcıların kimlik doğrulamasını sağlamak için bir middleware fonksiyonu oluşturalım.

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
    router.HandleFunc("/purchase", VerifyToken(PurchaseBook)).Methods("POST")
    router.HandleFunc("/loan", VerifyToken(LoanBook)).Methods("POST")
    router.HandleFunc("/return", VerifyToken(ReturnBook)).Methods("DELETE")
    router.HandleFunc("/stockout", VerifyToken(StockOut)).Methods("POST")
    router.HandleFunc("/scrap", VerifyToken(MarkAsScrap)).Methods("DELETE")

    log.Println("Sunucu 8080 portunda çalışıyor...")
    log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", router)) // SSL için
}
```

### Açıklamalar

- **Router**: API yönlendirmelerini tanımlar.
- **HTTP Sunucusu**: SSL üzerinden 8080 portunda dinlemeye başlar.

### Aşama 8: SSL Sertifikası Oluşturma

Güvenli bir bağlantı sağlamak için bir SSL sertifikası oluşturmalısınız. Aşağıdaki komutla bir sertifika ve anahtar oluşturabilirsiniz:

```bash
openssl req -new -x509 -days 365 -keyout server.key -out server.crt
```

### Aşama 9: Projeyi Çalıştırma

Projeyi çalıştırmak için PostgreSQL veritabanınızı ayarladıktan sonra aşağıdaki komutu çalıştırın:

```bash
go run main.go
```

### Aşama 10: API'yi Test Etme

API'yi test etmek için Postman veya cURL kullanabilirsiniz. Aşağıda bazı örnek istekler bulunmaktadır:

#### 1. Kullanıcı Girişi

```bash
curl -X POST https://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password123"}'
```

#### 2. Tüm Kitapları Alma

```bash
curl -X GET https://localhost:8080/books -H "Authorization: Bearer <TOKEN>"
```

#### 3. Kitap Satın Alma

```bash
curl -X POST https://localhost:8080/purchase -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"book_id": 1, "quantity": 5}'
```

#### 4. Kitap Ödünç Alma

```bash
curl -X POST https://localhost:8080/loan -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"user_id": 1, "book_id": 2}'
```

#### 5. Kitap İade Etme

```bash
curl -X DELETE https://localhost:8080/return -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"user_id": 1, "book_id": 2}'
```

#### 6. Depo Çıkışı

```bash
curl -X POST https://localhost:8080/stockout -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"book_id": 1, "quantity": 2}'
```

#### 7. Hurdaya Çıkarma

```bash
curl -X DELETE https://localhost:8080/scrap -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"id": 1}'
```

### Sonuç

Bu proje, kitap satın alma, iade işlemleri, depo girişleri, stok çıkışları ve hurdaya çıkarma işlemlerini yöneten bir web API'si olarak çalışır. Kullanıcılar kimlik doğrulama işlemi için JWT kullanırken, API SSL üzerinden güvenli bir şekilde çalışır. Proje, gerçek hayatta kullanılabilir bir temel oluşturur ve genişletilebilir. İleride eklemeler yapmak veya mevcut özellikleri geliştirmek oldukça kolaydır.


# Client

Aşağıda, yukarıda geliştirdiğimiz Web API projesine bağlanarak kullanıcı giriş işlemi yapabilen ve stok girişi yapabilen basit bir HTML web istemcisi oluşturacağız. Bu istemci, jQuery kullanarak API ile iletişim kuracak ve gerekli HTTP isteklerini gönderecektir. 

### HTML Web İstemcisi

Aşağıdaki adımlar, kullanıcıların API üzerinden giriş yapmalarını ve stok girişi yapmalarını sağlayan bir istemci oluşturur.

#### 1. HTML Dosyası Oluşturma

İlk olarak, `index.html` adlı bir dosya oluşturalım ve içerisine aşağıdaki kodları ekleyelim.

```html
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Stok Girişi</title>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }
        input[type="text"], input[type="password"], input[type="number"] {
            width: 100%;
            padding: 10px;
            margin: 5px 0;
        }
        button {
            padding: 10px 20px;
            background-color: #007BFF;
            color: white;
            border: none;
            cursor: pointer;
        }
        button:hover {
            background-color: #0056b3;
        }
        #response {
            margin-top: 20px;
            color: green;
        }
    </style>
</head>
<body>
    <h1>Stok Girişi</h1>
    <h2>Giriş Yap</h2>
    <input type="text" id="username" placeholder="Kullanıcı Adı" required>
    <input type="password" id="password" placeholder="Şifre" required>
    <button id="login-btn">Giriş Yap</button>

    <div id="stock-entry" style="display:none;">
        <h2>Stok Girişi</h2>
        <input type="number" id="bookId" placeholder="Kitap ID'si" required>
        <input type="number" id="quantity" placeholder="Miktar" required>
        <button id="stock-btn">Stok Girişi Yap</button>
    </div>

    <div id="response"></div>

    <script>
        // Kullanıcı giriş fonksiyonu
        $(document).ready(function() {
            $("#login-btn").click(function() {
                const username = $("#username").val();
                const password = $("#password").val();

                $.ajax({
                    url: 'https://localhost:8080/login',
                    type: 'POST',
                    contentType: 'application/json',
                    data: JSON.stringify({ username: username, password: password }),
                    success: function(data) {
                        // Token'ı al ve sakla
                        localStorage.setItem('token', data.token);
                        $("#response").text("Giriş başarılı! Stok girişi yapabilirsiniz.").css("color", "green");
                        $("#stock-entry").show();
                    },
                    error: function() {
                        $("#response").text("Giriş başarısız! Lütfen kullanıcı adınızı ve şifrenizi kontrol edin.").css("color", "red");
                    }
                });
            });

            // Stok girişi yapma fonksiyonu
            $("#stock-btn").click(function() {
                const bookId = $("#bookId").val();
                const quantity = $("#quantity").val();
                const token = localStorage.getItem('token');

                $.ajax({
                    url: 'https://localhost:8080/stockout',
                    type: 'POST',
                    contentType: 'application/json',
                    headers: {
                        'Authorization': 'Bearer ' + token
                    },
                    data: JSON.stringify({ book_id: bookId, quantity: quantity }),
                    success: function() {
                        $("#response").text("Stok girişi başarılı!").css("color", "green");
                    },
                    error: function() {
                        $("#response").text("Stok girişi başarısız!").css("color", "red");
                    }
                });
            });
        });
    </script>
</body>
</html>
```

### Açıklama

1. **HTML Yapısı**: Basit bir form yapısı ile kullanıcıdan kullanıcı adı ve şifre alır, ardından stok girişi yapabilmesi için gerekli alanları gösterir.

2. **jQuery**: jQuery kütüphanesi kullanarak API ile etkileşim kurmak için AJAX istekleri gönderir.
   - Kullanıcı giriş işlemi için `/login` endpoint'ine bir POST isteği gönderir. Başarılı olursa, dönen token'ı yerel depolama (localStorage) içinde saklar ve stok girişi alanını gösterir.
   - Kullanıcı stok girişi yapmak istediğinde, `/stockout` endpoint'ine gerekli bilgilerle birlikte bir POST isteği gönderir. Burada `Authorization` başlığına token eklenir.

3. **Stil**: Temel stil ile kullanıcı arayüzü daha okunaklı hale getirilmiştir.

### Projeyi Çalıştırma

1. **HTML Dosyasını Kaydetme**: `index.html` dosyasını oluşturup yukarıdaki kodu yapıştırın.
2. **Sunucu**: Go API sunucunuzun çalıştığından emin olun.
3. **Tarayıcıda Açma**: `index.html` dosyasını bir tarayıcıda açın.
4. **Giriş Yapma**: Kullanıcı adı ve şifre ile giriş yaparak token'ı alabilir ve stok girişi yapabilirsiniz.

### Önemli Notlar

- **CORS**: Eğer API ve istemci farklı portlarda çalışıyorsa, CORS (Cross-Origin Resource Sharing) ayarlarını yapmalısınız. API sunucusunun CORS ayarlarını yapılandırmak için `github.com/rs/cors` kütüphanesini kullanabilirsiniz.
- **HTTPS**: İstemci uygulamasının HTTPS ile çalıştığından emin olun. Bu, güvenli bir bağlantı sağlar ve API'ye isteklerinizi güvenli bir şekilde gönderir.
- **Hatalar**: Hata mesajları kullanıcıya bildirilecektir. Bu, kullanıcı deneyimini iyileştirir ve olası sorunları anlamalarına yardımcı olur.

Bu istemci, yukarıda tanımlanan API ile etkileşimde bulunmak için kullanılabilir. Herhangi bir geliştirme veya değişiklik yapmak için ihtiyacınıza göre düzenleyebilirsiniz.