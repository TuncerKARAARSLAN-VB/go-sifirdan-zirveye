## Go ile SQL Komutları ve Transaction Yönetimi

Go programlama dilinde SQL komutları ve transaction yönetimi, veritabanı işlemlerinin güvenli ve etkili bir şekilde gerçekleştirilmesini sağlar. Bu bölümde, temel SQL komutları ile transaction işlemlerinin nasıl gerçekleştirileceği detaylı bir şekilde açıklanacaktır.

### 1. SQL Komutları

SQL, veritabanı üzerinde veri ekleme, okuma, güncelleme ve silme işlemlerini gerçekleştirmek için kullanılan bir dildir. Aşağıda en yaygın SQL komutları ve örnekleri verilmiştir:

#### 1.1. CREATE

Yeni bir tablo oluşturmak için `CREATE` komutu kullanılır.

**Örnek:**

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    age INT
);
```

#### 1.2. INSERT

Veritabanına yeni bir kayıt eklemek için `INSERT` komutu kullanılır.

**Örnek:**

```sql
INSERT INTO users (name, age) VALUES ('Alice', 30);
```

#### 1.3. SELECT

Veritabanından veri okumak için `SELECT` komutu kullanılır.

**Örnek:**

```sql
SELECT * FROM users;
```

#### 1.4. UPDATE

Mevcut bir kaydı güncellemek için `UPDATE` komutu kullanılır.

**Örnek:**

```sql
UPDATE users SET age = 31 WHERE name = 'Alice';
```

#### 1.5. DELETE

Veritabanından bir kaydı silmek için `DELETE` komutu kullanılır.

**Örnek:**

```sql
DELETE FROM users WHERE name = 'Alice';
```

### 2. Transaction Yönetimi

Transaction, bir grup SQL komutunun atomik bir şekilde (ya tamamı ya da hiçbiri) çalıştırılmasını sağlayan bir mekanizmadır. Transaction kullanarak veritabanındaki değişikliklerin tutarlılığını koruyabiliriz.

#### 2.1. Transaction Başlatma

Transaction başlatmak için `Begin` fonksiyonu kullanılır. Aşağıda, transaction işlemi için kullanılacak temel fonksiyonları ve örnekleri göreceksiniz.

### 2.2. Transaction Fonksiyonları

1. **Begin**: Transaction başlatır.
2. **Commit**: Yapılan değişiklikleri kalıcı hale getirir.
3. **Rollback**: Yapılan değişiklikleri geri alır.

#### 2.3. Örnek Kod

Aşağıdaki örnek, bir veritabanına bağlanmayı, kullanıcı eklemeyi ve transaction işlemlerini göstermektedir.

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
    // MySQL veritabanı bağlantısı
    dsn := "user:password@tcp(localhost:3306)/testdb" // DSN: Data Source Name
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

    // Transaction başlatma
    tx, err := db.Begin() // Yeni bir transaction başlat
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }

    // Kullanıcı ekleme
    if err := insertUser(tx, "Alice", 30); err != nil {
        tx.Rollback() // Hata varsa transaction'ı geri al
        log.Fatal(err) // Hata varsa kaydet
    }

    if err := insertUser(tx, "Bob", 25); err != nil {
        tx.Rollback() // Hata varsa transaction'ı geri al
        log.Fatal(err) // Hata varsa kaydet
    }

    // Transaction'ı başarılı bir şekilde tamamla
    if err := tx.Commit(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }

    fmt.Println("Kullanıcılar başarıyla eklendi!")
}

// Kullanıcı ekleme fonksiyonu
func insertUser(tx *sql.Tx, name string, age int) error {
    query := "INSERT INTO users (name, age) VALUES (?, ?)" // Kullanıcı ekleme sorgusu
    _, err := tx.Exec(query, name, age) // Sorguyu çalıştır
    return err // Hata varsa döndür
}
```

### 3. Çıktı

```plaintext
MySQL veritabanına başarıyla bağlandı!
Kullanıcılar başarıyla eklendi!
```

### 4. Örnek Transaction Durumunda Hata Yönetimi

Aşağıdaki örnek, bir transaction sırasında hata oluşursa rollback işlemi yaparak durumu nasıl yöneteceğinizi gösterir.

```go
func main() {
    // MySQL veritabanı bağlantısı
    dsn := "user:password@tcp(localhost:3306)/testdb" // DSN: Data Source Name
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

    // Transaction başlatma
    tx, err := db.Begin() // Yeni bir transaction başlat
    if err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }

    // Kullanıcı ekleme
    if err := insertUser(tx, "Alice", 30); err != nil {
        tx.Rollback() // Hata varsa transaction'ı geri al
        log.Fatal(err) // Hata varsa kaydet
    }

    // Intentional error for testing rollback
    if err := insertUser(tx, "", 25); err != nil {
        tx.Rollback() // Hata varsa transaction'ı geri al
        log.Fatal("Kullanıcı adı boş olamaz: ", err) // Hata mesajı ile kaydet
    }

    // Transaction'ı başarılı bir şekilde tamamla
    if err := tx.Commit(); err != nil {
        log.Fatal(err) // Hata varsa kaydet
    }

    fmt.Println("Kullanıcılar başarıyla eklendi!")
}

// Kullanıcı ekleme fonksiyonu
func insertUser(tx *sql.Tx, name string, age int) error {
    query := "INSERT INTO users (name, age) VALUES (?, ?)" // Kullanıcı ekleme sorgusu
    _, err := tx.Exec(query, name, age) // Sorguyu çalıştır
    return err // Hata varsa döndür
}
```

### 5. Çıktı (Hata Durumu)

```plaintext
MySQL veritabanına başarıyla bağlandı!
Kullanıcı adı boş olamaz: Error 1364: Field 'name' doesn't have a default value
```

## Sonuç

Yukarıda, Go dilinde SQL komutları ve transaction yönetiminin nasıl yapılacağını detaylı bir şekilde inceledik. `CREATE`, `INSERT`, `SELECT`, `UPDATE`, ve `DELETE` komutlarının yanı sıra, transaction işlemleri ile veritabanı üzerinde güvenli bir şekilde nasıl işlem yapılabileceği gösterildi. Her iki örnek ile hata yönetimi ve rollback işlemlerinin nasıl gerçekleştirileceği hakkında bilgi verilmiştir. Bu sayede veritabanı işlemlerinin güvenilirliğini artırabiliriz.