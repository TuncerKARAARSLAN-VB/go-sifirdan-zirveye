Go programlama dilinde CRUD (Create, Read, Update, Delete) işlemleri, veritabanı ile etkileşimde bulunmanın temel yollarını temsil eder. Aşağıda, MySQL ve PostgreSQL veritabanlarını kullanarak bir CRUD uygulaması geliştireceğiz. Her bir işlem için detaylı açıklamalar ve örnekler sunulacaktır.

## 1. Gerekli Paketlerin Kurulumu

Öncelikle, veritabanı işlemleri için gerekli paketleri yüklememiz gerekiyor. Terminalden aşağıdaki komutları çalıştırarak gerekli kütüphaneleri yükleyin:

### MySQL için

```bash
go get -u github.com/go-sql-driver/mysql
```

### PostgreSQL için

```bash
go get -u github.com/lib/pq
```

## 2. MySQL ile CRUD Uygulaması

### 2.1 Veritabanı Bağlantısı ve Tablo Oluşturma

Aşağıdaki kod, bir MySQL veritabanına bağlanmayı ve `users` tablosunu oluşturmayı gösterir.

#### 2.1.1 `main.go` Dosyasını Oluşturma

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // MySQL sürücüsü
)

// Kullanıcı yapısı
type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "root:password@tcp(127.0.0.1:3306)/testdb" // DSN: Data Source Name
    db, err := sql.Open("mysql", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Println("MySQL veritabanına başarıyla bağlandı!")

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

### 2.2 CRUD Fonksiyonlarının Tanımlanması

Aşağıdaki fonksiyonlar, CRUD işlemlerini gerçekleştirecek:

#### 2.2.1 Kullanıcı Ekleme (Create)

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
```

#### 2.2.2 Kullanıcı Okuma (Read)

```go
// Kullanıcıları listeleme fonksiyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    fmt.Println("Kullanıcılar:")
    // Sonuçları döngü ile işle
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", user.ID, user.Name, user.Age) // Kullanıcı bilgilerini yazdır
    }
}
```

#### 2.2.3 Kullanıcı Güncelleme (Update)

```go
// Kullanıcı güncelleme fonksiyonu
func updateUser(db *sql.DB, id int, name string, age int) {
    query := "UPDATE users SET name = ?, age = ? WHERE id = ?" // Güncelleme sorgusu
    result, err := db.Exec(query, name, age, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Güncellenen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı güncelleme mesajı
}
```

#### 2.2.4 Kullanıcı Silme (Delete)

```go
// Kullanıcı silme fonksiyonu
func deleteUser(db *sql.DB, id int) {
    query := "DELETE FROM users WHERE id = ?" // Silme sorgusu
    result, err := db.Exec(query, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Silinen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı silme mesajı
}
```

### 2.3 Fonksiyonları Ana Fonksiyona Eklemek

Ana fonksiyona CRUD işlemlerini gerçekleştiren fonksiyonları ekleyelim.

```go
func main() {
    // ... (önceki kod)

    // CRUD işlemleri
    insertUser(db, "Alice", 30) // Alice adında bir kullanıcı ekle
    insertUser(db, "Bob", 25)   // Bob adında bir kullanıcı ekle
    listUsers(db)               // Kullanıcıları listele

    updateUser(db, 1, "Alice Smith", 31) // Alice'in bilgilerini güncelle
    listUsers(db)                        // Güncelleme sonrası kullanıcıları listele

    deleteUser(db, 2) // Bob'u sil
    listUsers(db)     // Son durumu listele
}
```

### 2.4 Tam Kod ve Çıktı

#### Tam Kod

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // MySQL sürücüsü
)

// Kullanıcı yapısı
type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "root:password@tcp(127.0.0.1:3306)/testdb" // DSN: Data Source Name
    db, err := sql.Open("mysql", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Println("MySQL veritabanına başarıyla bağlandı!")

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

    // CRUD işlemleri
    insertUser(db, "Alice", 30) // Alice adında bir kullanıcı ekle
    insertUser(db, "Bob", 25)   // Bob adında bir kullanıcı ekle
    listUsers(db)               // Kullanıcıları listele

    updateUser(db, 1, "Alice Smith", 31) // Alice'in bilgilerini güncelle
    listUsers(db)                        // Güncelleme sonrası kullanıcıları listele

    deleteUser(db, 2) // Bob'u sil
    listUsers(db)     // Son durumu listele
}

// Kullanıcı ekleme fonksiyonu
func insertUser(db *sql.DB, name string, age int) {
    query := "INSERT INTO users (name, age) VALUES (?, ?)" // Veri ekleme sorgusu
    result, err := db.Exec(query, name, age) // Sorguyu çalıştır
    if err != nil

 {
        log.Fatal(err) // Hata varsa kaydet
    }
    id, err := result.LastInsertId() // Eklenen son verinin ID'sini al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Yeni kullanıcı eklendi: %d - %s (%d yaşında)\n", id, name, age) // Başarılı ekleme mesajı
}

// Kullanıcıları listeleme fonksiyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    fmt.Println("Kullanıcılar:")
    // Sonuçları döngü ile işle
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", user.ID, user.Name, user.Age) // Kullanıcı bilgilerini yazdır
    }
}

// Kullanıcı güncelleme fonksiyonu
func updateUser(db *sql.DB, id int, name string, age int) {
    query := "UPDATE users SET name = ?, age = ? WHERE id = ?" // Güncelleme sorgusu
    result, err := db.Exec(query, name, age, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Güncellenen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı güncelleme mesajı
}

// Kullanıcı silme fonksiyonu
func deleteUser(db *sql.DB, id int) {
    query := "DELETE FROM users WHERE id = ?" // Silme sorgusu
    result, err := db.Exec(query, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Silinen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı silme mesajı
}
```

#### Çıktı

```plaintext
MySQL veritabanına başarıyla bağlandı!
users tablosu başarıyla oluşturuldu!
Yeni kullanıcı eklendi: 1 - Alice (30 yaşında)
Yeni kullanıcı eklendi: 2 - Bob (25 yaşında)
Kullanıcılar:
Kullanıcı: 1 - Alice (30 yaşında)
Kullanıcı: 2 - Bob (25 yaşında)
Güncellenen kullanıcı sayısı: 1
Kullanıcılar:
Kullanıcı: 1 - Alice Smith (31 yaşında)
Kullanıcı: 2 - Bob (25 yaşında)
Silinen kullanıcı sayısı: 1
Kullanıcılar:
Kullanıcı: 1 - Alice Smith (31 yaşında)
```

## 3. PostgreSQL ile CRUD Uygulaması

### 3.1 Veritabanı Bağlantısı ve Tablo Oluşturma

Aşağıdaki kod, bir PostgreSQL veritabanına bağlanmayı ve `users` tablosunu oluşturmayı gösterir.

#### 3.1.1 `main.go` Dosyasını Oluşturma

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL sürücüsü
)

// Kullanıcı yapısı
type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "user=postgres password=password dbname=testdb sslmode=disable" // DSN: Data Source Name
    db, err := sql.Open("postgres", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Println("PostgreSQL veritabanına başarıyla bağlandı!")

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
    fmt.Println("users tablosu başarıyla oluşturuldu!")
}
```

### 3.2 CRUD Fonksiyonlarının Tanımlanması

Aşağıdaki fonksiyonlar, CRUD işlemlerini gerçekleştirecek:

#### 3.2.1 Kullanıcı Ekleme (Create)

```go
// Kullanıcı ekleme fonksiyonu
func insertUser(db *sql.DB, name string, age int) {
    query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id" // Veri ekleme sorgusu
    var id int
    err := db.QueryRow(query, name, age).Scan(&id) // Sorguyu çalıştır ve ID'yi al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Yeni kullanıcı eklendi: %d - %s (%d yaşında)\n", id, name, age) // Başarılı ekleme mesajı
}
```

#### 3.2.2 Kullanıcı Okuma (Read)

```go
// Kullanıcıları listeleme fonksiyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    fmt.Println("Kullanıcılar:")
    // Sonuçları döngü ile işle
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", user.ID, user.Name, user.Age) // Kullanıcı bilgilerini yazdır
    }
}
```

#### 3.2.3 Kullanıcı Güncelleme (Update)

```go
// Kullanıcı güncelleme fonksiyonu
func updateUser(db *sql.DB, id int, name string, age int) {
    query := "UPDATE users SET name = $1, age = $2 WHERE id = $3" // Güncelleme sorgusu
    result, err := db.Exec(query, name, age, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Güncellenen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı güncelleme mesajı
}
```

#### 3.2.4 Kullanıcı Silme (Delete)

```go
// Kullanıcı silme fonksiyonu
func deleteUser(db *sql.DB, id int) {
    query := "DELETE FROM users WHERE id = $1" // Silme sorgusu
    result, err := db.Exec(query, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Silinen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı silme mesajı
}
```

### 3.3 Fonksiyonları Ana Fonksiyona Eklemek

Ana fonksiyona CRUD işlemlerini gerçekleştiren fonksiyonları ekleyelim.

```go
func main() {
    // ... (önceki kod)

    // CRUD işlemleri
    insertUser(db, "Alice", 30) // Alice adında bir kullanıcı ekle
    insertUser(db, "Bob", 25)   // Bob adında bir kullanıcı ekle
    listUsers(db)               // Kullanıcıları listele

    updateUser(db, 1, "Alice Smith", 31) // Alice'in bilgilerini güncelle
    listUsers(db)                       

 // Güncelleme sonrası kullanıcıları listele

    deleteUser(db, 2) // Bob'u sil
    listUsers(db)     // Son durumu listele
}
```

### 3.4 Tam Kod ve Çıktı

#### Tam Kod

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL sürücüsü
)

// Kullanıcı yapısı
type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Veritabanı bağlantı dizesi
    dsn := "user=postgres password=password dbname=testdb sslmode=disable" // DSN: Data Source Name
    db, err := sql.Open("postgres", dsn) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer db.Close() // Fonksiyon bitince bağlantıyı kapat

    // Bağlantıyı kontrol et
    if err := db.Ping(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Println("PostgreSQL veritabanına başarıyla bağlandı!")

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
    fmt.Println("users tablosu başarıyla oluşturuldu!")

    // CRUD işlemleri
    insertUser(db, "Alice", 30) // Alice adında bir kullanıcı ekle
    insertUser(db, "Bob", 25)   // Bob adında bir kullanıcı ekle
    listUsers(db)               // Kullanıcıları listele

    updateUser(db, 1, "Alice Smith", 31) // Alice'in bilgilerini güncelle
    listUsers(db)                        // Güncelleme sonrası kullanıcıları listele

    deleteUser(db, 2) // Bob'u sil
    listUsers(db)     // Son durumu listele
}

// Kullanıcı ekleme fonksiyonu
func insertUser(db *sql.DB, name string, age int) {
    query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id" // Veri ekleme sorgusu
    var id int
    err := db.QueryRow(query, name, age).Scan(&id) // Sorguyu çalıştır ve ID'yi al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Yeni kullanıcı eklendi: %d - %s (%d yaşında)\n", id, name, age) // Başarılı ekleme mesajı
}

// Kullanıcıları listeleme fonksiyonu
func listUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users") // Sorgu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    defer rows.Close() // Fonksiyon bitince sonuçları kapat

    fmt.Println("Kullanıcılar:")
    // Sonuçları döngü ile işle
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil { // Sonuçları al
            log.Fatal(err) // Hata varsa kaydet
        }
        fmt.Printf("Kullanıcı: %d - %s (%d yaşında)\n", user.ID, user.Name, user.Age) // Kullanıcı bilgilerini yazdır
    }
}

// Kullanıcı güncelleme fonksiyonu
func updateUser(db *sql.DB, id int, name string, age int) {
    query := "UPDATE users SET name = $1, age = $2 WHERE id = $3" // Güncelleme sorgusu
    result, err := db.Exec(query, name, age, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Güncellenen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı güncelleme mesajı
}

// Kullanıcı silme fonksiyonu
func deleteUser(db *sql.DB, id int) {
    query := "DELETE FROM users WHERE id = $1" // Silme sorgusu
    result, err := db.Exec(query, id) // Sorguyu çalıştır
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    rowsAffected, err := result.RowsAffected() // Etkilenen satır sayısını al
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }
    fmt.Printf("Silinen kullanıcı sayısı: %d\n", rowsAffected) // Başarılı silme mesajı
}
```

#### Çıktı

```plaintext
PostgreSQL veritabanına başarıyla bağlandı!
users tablosu başarıyla oluşturuldu!
Yeni kullanıcı eklendi: 1 - Alice (30 yaşında)
Yeni kullanıcı eklendi: 2 - Bob (25 yaşında)
Kullanıcılar:
Kullanıcı: 1 - Alice (30 yaşında)
Kullanıcı: 2 - Bob (25 yaşında)
Güncellenen kullanıcı sayısı: 1
Kullanıcılar:
Kullanıcı: 1 - Alice Smith (31 yaşında)
Kullanıcı: 2 - Bob (25 yaşında)
Silinen kullanıcı sayısı: 1
Kullanıcılar:
Kullanıcı: 1 - Alice Smith (31 yaşında)
```

## Sonuç

Yukarıdaki örneklerde, Go dilini kullanarak MySQL ve PostgreSQL veritabanlarında CRUD işlemlerinin nasıl gerçekleştirileceğini detaylı bir şekilde açıkladık. Her işlem için ilgili fonksiyonlar, bağlantı oluşturma, tablo oluşturma ve işlem sonuçlarıyla birlikte çıktılar sağlandı. Her iki veritabanında da benzer bir yapı ile CRUD işlemlerini gerçekleştirebiliriz.