Go dilinde veritabanı işlemleri yapmak için `database/sql` paketi kullanılır. Bu paket, veritabanlarına erişmek için genel bir arayüz sağlar ve `MySQL`, `PostgreSQL` gibi farklı veritabanı sistemleri ile uyumlu çalışabilir. Aşağıda, Go ile `MySQL` ve `PostgreSQL` veritabanlarını nasıl kullanacağınızı detaylı bir şekilde anlatan örnekler yer almaktadır.

## 1. MySQL ile Çalışma

### 1.1 Gerekli Paketlerin Kurulumu

MySQL ile çalışabilmek için Go'da `go-sql-driver/mysql` paketine ihtiyacınız var. Bu paketi kurmak için terminalde aşağıdaki komutu çalıştırabilirsiniz:

```bash
go get -u github.com/go-sql-driver/mysql
```

### 1.2 MySQL Veritabanı Bağlantısı

Aşağıdaki kod parçası, bir MySQL veritabanına bağlanmayı ve basit bir tablo oluşturmayı göstermektedir.

#### 1.2.1 `main.go` Dosyasını Oluşturma

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // MySQL sürücüsünü import et
)

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "root:password@tcp(127.0.0.1:3306)/testdb" // DSN: Data Source Name
    db, err := sql.Open("mysql", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet ve çık
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Bağlantı sağlanamazsa hata kaydet
    }
    fmt.Println("MySQL veritabanına başarıyla bağlandı!") // Bağlantı başarılı

    // Tablo oluşturma
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        age INT
    )`
    if _, err := db.Exec(createTableQuery); err != nil {
        log.Fatal(err) // Tablo oluşturulamazsa hata kaydet
    }
    fmt.Println("users tablosu başarıyla oluşturuldu!")
}
```

### 1.3 Kodu Çalıştırma

```bash
go run main.go
```

### Çıktı

```plaintext
MySQL veritabanına başarıyla bağlandı!
users tablosu başarıyla oluşturuldu!
```

**Açıklama**: Kod çalıştırıldığında, veritabanına bağlantı sağlanmış ve `users` adlı bir tablo oluşturulmuştur.

### 1.4 Veri Ekleme

Şimdi, `users` tablosuna veri eklemek için bir fonksiyon yazalım.

```go
// Kullanıcı ekleme fonksiyonu
func insertUser(db *sql.DB, name string, age int) {
    query := "INSERT INTO users (name, age) VALUES (?, ?)" // Veri ekleme sorgusu
    result, err := db.Exec(query, name, age) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    id, err := result.LastInsertId() // Eklenen son verinin ID'sini al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Yeni kullanıcı eklendi: %d - %s (%d yaşında)\n", id, name, age) // Başarılı ekleme mesajı
}

// main fonksiyonuna ekleyin
insertUser(db, "Alice", 30) // Alice adında bir kullanıcı ekle
```

### Çıktı

```plaintext
Yeni kullanıcı eklendi: 1 - Alice (30 yaşında)
```

**Açıklama**: Kullanıcı başarıyla `users` tablosuna eklenmiştir.

### 1.5 Veri Sorgulama

Tablodan verileri sorgulamak için aşağıdaki kodu ekleyin:

```go
// Kullanıcıları listeleme fonksiyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    // Sonuçları döngü ile işle
    for rows.Next() {
        var id int
        var name string
        var age int
        if err := rows.Scan(&id, &name, &age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", id, name, age) // Kullanıcı bilgilerini yazdır
    }
}

// main fonksiyonuna ekleyin
listUsers(db) // Kullanıcıları listele
```

### Çıktı

```plaintext
Kullanıcı: 1 - Alice (30 yaşında)
```

**Açıklama**: `users` tablosundaki veriler başarıyla listelenmiştir.

## 2. PostgreSQL ile Çalışma

PostgreSQL ile çalışmak için `pq` paketini kullanacağız. Aşağıda PostgreSQL ile nasıl bağlantı kuracağınızı ve veri ekleyeceğinizi gösteren örnek bulunmaktadır.

### 2.1 Gerekli Paketlerin Kurulumu

PostgreSQL ile çalışmak için aşağıdaki komutu kullanarak `pq` paketini kurabilirsiniz:

```bash
go get -u github.com/lib/pq
```

### 2.2 PostgreSQL Veritabanı Bağlantısı

Aşağıdaki kod, PostgreSQL veritabanına bağlanmayı ve basit bir tablo oluşturmayı göstermektedir.

#### 2.2.1 `main.go` Dosyasını Güncelleme

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL sürücüsünü import et
)

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "user=postgres password=password dbname=testdb sslmode=disable" // DSN: Data Source Name
    db, err := sql.Open("postgres", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet ve çık
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Bağlantı sağlanamazsa hata kaydet
    }
    fmt.Println("PostgreSQL veritabanına başarıyla bağlandı!") // Bağlantı başarılı

    // Tablo oluşturma
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        age INT
    )`
    if _, err := db.Exec(createTableQuery); err != nil {
        log.Fatal(err) // Tablo oluşturulamazsa hata kaydet
    }
    fmt.Println("users tablosu başarıyla oluşturuldu!") // Tablo başarıyla oluşturuldu
}
```

### 2.3 Kodu Çalıştırma

```bash
go run main.go
```

### Çıktı

```plaintext
PostgreSQL veritabanına başarıyla bağlandı!
users tablosu başarıyla oluşturuldu!
```

**Açıklama**: PostgreSQL veritabanına bağlantı sağlanmış ve `users` adlı bir tablo oluşturulmuştur.

### 2.4 Veri Ekleme

PostgreSQL tablosuna veri eklemek için aynı şekilde bir fonksiyon yazabiliriz:

```go
// Kullanıcı ekleme fonksiyonu
func insertUser(db *sql.DB, name string, age int) {
    query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id" // Veri ekleme sorgusu
    var id int
    err := db.QueryRow(query, name, age).Scan(&id) // Sorguyu çalıştır ve dönen ID'yi al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Yeni kullanıcı eklendi: %d - %s (%d yaşında)\n", id, name, age) // Başarılı ekleme mesajı
}

// main fonksiyonuna ekleyin
insertUser(db, "Bob", 25) // Bob adında bir kullanıcı ekle
```

### Çıktı

```plaintext
Yeni kullanıcı eklendi: 1 - Bob (25 yaşında)
```

**Açıklama**: Kullanıcı başarıyla `users` tablosuna eklenmiştir.

### 2.5 Veri Sorgulama

PostgreSQL tablosundan verileri sorgulamak için aşağıdaki kodu ekleyin:

```go
// Kullanıcıları listeleme fonks

iyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    // Sonuçları döngü ile işle
    for rows.Next() {
        var id int
        var name string
        var age int
        if err := rows.Scan(&id, &name, &age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", id, name, age) // Kullanıcı bilgilerini yazdır
    }
}

// main fonksiyonuna ekleyin
listUsers(db) // Kullanıcıları listele
```

### Çıktı

```plaintext
Kullanıcı: 1 - Bob (25 yaşında)
```

**Açıklama**: `users` tablosundaki veriler başarıyla listelenmiştir.

## 3. Özet

- **MySQL ve PostgreSQL**: Go ile bu iki veritabanına nasıl bağlanacağınızı, tablo oluşturacağınızı, veri ekleyeceğinizi ve verileri sorgulayacağınızı gösterdik.
- **`database/sql` Paketi**: Veritabanı işlemlerinde standart bir arayüz sunar.
- **Sürücü Kullanımı**: Her veritabanı için uygun sürücüyü kullanmanız gerekmektedir (`mysql` için `go-sql-driver/mysql`, `PostgreSQL` için `pq`).

Bu örnekler, Go ile veritabanı işlemleri yaparken size temel bir anlayış kazandıracaktır. Gerçek uygulamalarda hata yönetimi, veritabanı bağlantı havuzları ve daha karmaşık sorgular gibi konuları da göz önünde bulundurmalısınız.