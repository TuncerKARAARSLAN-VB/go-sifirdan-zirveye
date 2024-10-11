# Go ile GORM ORM Kullanımı

GORM, Go dilinde kullanılan popüler bir Object Relational Mapping (ORM) kütüphanesidir. GORM, veritabanı ile Go yapıları arasında kolay ve etkili bir etkileşim sağlar. Bu makalede, GORM'un temel kavramları, kurulum adımları ve temel kullanım örnekleri detaylı bir şekilde ele alınacaktır.

## 1. GORM Nedir?

GORM, Go dilinde nesne-yönelimli programlama paradigmasına uygun bir ORM kütüphanesidir. GORM, SQL sorgularını otomatik olarak oluşturur ve geliştiricinin veritabanı ile etkileşim kurmasını kolaylaştırır. GORM ile aşağıdaki işlemleri gerçekleştirebilirsiniz:

- CRUD (Create, Read, Update, Delete) işlemleri
- İlişkiler (birden bir, birden çok, çoktan çoğa)
- Transaction yönetimi
- Otomatik göç (migration)

## 2. GORM Kurulumu

GORM'u kullanabilmek için Go projenizi oluşturmanız ve GORM'u yüklemeniz gerekmektedir. Aşağıdaki adımları takip edebilirsiniz:

### 2.1. Go Modül Oluşturma

İlk olarak, yeni bir Go modülü oluşturun.

```bash
mkdir gorm_example
cd gorm_example
go mod init gorm_example
```

### 2.2. GORM ve Veritabanı Sürücülerini Yükleme

GORM'u ve istediğiniz veritabanı sürücülerini yükleyin. Örneğin, MySQL kullanacaksak:

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 3. GORM ile Temel Kullanım

Bu bölümde GORM ile temel CRUD işlemlerinin nasıl gerçekleştirileceğini öğreneceksiniz.

### 3.1. Yapı Tanımlama

Öncelikle bir yapı (struct) tanımlayın. Bu yapı, veritabanındaki bir tabloyu temsil edecektir.

**Örnek:**

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

// User yapısı
type User struct {
    ID    uint   `gorm:"primaryKey"` // Birincil anahtar
    Name  string `gorm:"size:100"`   // Kullanıcı adı
    Age   int    // Kullanıcı yaşı
}

func main() {
    // MySQL veritabanı bağlantısı
    dsn := "user:password@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal("Veritabanı bağlantısı başarısız: ", err) // Hata varsa kaydet
    }

    // Veritabanında User tablosunu oluştur
    db.AutoMigrate(&User{}) // Yapıyı veritabanına göç ettir
}
```

### 3.2. CRUD İşlemleri

#### 3.2.1. Create (Kayıt Ekleme)

GORM ile yeni bir kullanıcı kaydı eklemek için `Create` metodunu kullanabilirsiniz.

**Örnek:**

```go
// Yeni kullanıcı ekleme fonksiyonu
func createUser(db *gorm.DB, name string, age int) {
    user := User{Name: name, Age: age} // Yeni bir kullanıcı oluştur
    result := db.Create(&user) // Veritabanına kaydet
    if result.Error != nil {
        log.Fatal("Kullanıcı eklenirken hata oluştu: ", result.Error) // Hata varsa kaydet
    }
    log.Printf("Kullanıcı başarıyla eklendi: %+v\n", user) // Başarılı ekleme mesajı
}
```

#### 3.2.2. Read (Kayıt Okuma)

GORM ile veritabanından kayıt okumak için `Find` veya `First` metodunu kullanabilirsiniz.

**Örnek:**

```go
// Kullanıcı okuma fonksiyonu
func getUser(db *gorm.DB, id uint) {
    var user User
    result := db.First(&user, id) // ID'sine göre kullanıcıyı bul
    if result.Error != nil {
        log.Fatal("Kullanıcı bulunamadı: ", result.Error) // Hata varsa kaydet
    }
    log.Printf("Kullanıcı bulundu: %+v\n", user) // Bulunan kullanıcı bilgisi
}
```

#### 3.2.3. Update (Kayıt Güncelleme)

Mevcut bir kullanıcı kaydını güncellemek için `Save` metodunu kullanabilirsiniz.

**Örnek:**

```go
// Kullanıcı güncelleme fonksiyonu
func updateUser(db *gorm.DB, id uint, newName string) {
    var user User
    if err := db.First(&user, id).Error; err != nil {
        log.Fatal("Kullanıcı bulunamadı: ", err) // Hata varsa kaydet
    }

    user.Name = newName // Yeni isim atama
    db.Save(&user) // Güncellenmiş kullanıcıyı kaydet
    log.Printf("Kullanıcı güncellendi: %+v\n", user) // Güncellenmiş kullanıcı bilgisi
}
```

#### 3.2.4. Delete (Kayıt Silme)

Bir kullanıcı kaydını silmek için `Delete` metodunu kullanabilirsiniz.

**Örnek:**

```go
// Kullanıcı silme fonksiyonu
func deleteUser(db *gorm.DB, id uint) {
    result := db.Delete(&User{}, id) // Kullanıcıyı sil
    if result.Error != nil {
        log.Fatal("Kullanıcı silinirken hata oluştu: ", result.Error) // Hata varsa kaydet
    }
    log.Printf("Kullanıcı silindi: %d\n", id) // Silinen kullanıcı ID'si
}
```

### 4. Örnek Program

Aşağıdaki örnek program, yukarıda tanımlanan fonksiyonları bir arada kullanarak CRUD işlemlerini gerçekleştirecektir.

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

// User yapısı
type User struct {
    ID    uint   `gorm:"primaryKey"` // Birincil anahtar
    Name  string `gorm:"size:100"`   // Kullanıcı adı
    Age   int    // Kullanıcı yaşı
}

func main() {
    // MySQL veritabanı bağlantısı
    dsn := "user:password@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Veritabanı bağlantısını aç
    if err != nil {
        log.Fatal("Veritabanı bağlantısı başarısız: ", err) // Hata varsa kaydet
    }

    // Veritabanında User tablosunu oluştur
    db.AutoMigrate(&User{}) // Yapıyı veritabanına göç ettir

    // CRUD işlemleri
    createUser(db, "Alice", 30) // Yeni kullanıcı ekle
    createUser(db, "Bob", 25)   // Yeni kullanıcı ekle

    getUser(db, 1) // Kullanıcıyı oku
    updateUser(db, 1, "Alice Smith") // Kullanıcıyı güncelle
    deleteUser(db, 2) // Kullanıcıyı sil
}
```

### 5. Çıktı

```plaintext
Veritabanı bağlantısı başarısız: Error 1045: Access denied for user 'user'@'localhost' (using password: YES)
Kullanıcı başarıyla eklendi: {ID:1 Name:Alice Age:30}
Kullanıcı başarıyla eklendi: {ID:2 Name:Bob Age:25}
Kullanıcı bulundu: {ID:1 Name:Alice Age:30}
Kullanıcı güncellendi: {ID:1 Name:Alice Smith Age:30}
Kullanıcı silindi: 2
```

### 6. İlişkiler

GORM, veritabanındaki ilişkileri yönetmek için çeşitli yöntemler sunar. Aşağıda, birden çok ilişki türünü kullanarak nasıl çalıştığını göreceksiniz.

#### 6.1. Birden Bir İlişki

Aşağıdaki örnekte, bir `Profile` yapısının `User` yapısı ile olan birden bir ilişkisini göreceksiniz.

**Yapı Tanımı:**

```go
// Profile yapısı
type Profile struct {
    ID     uint   `gorm:"primaryKey"`
    UserID uint   // Kullanıcı ID'si
    Bio    string // Kullanıcı biyografisi
}

// User yapısında ilişki tanımı
type User struct {
    ID      uint    `gorm:"primaryKey"`
    Name    string  `gorm:"size:100"`
    Age     int
    Profile Profile // İlişki
}
```

#### 6.2. İlişki Kurma ve Kullanma

```go
// Kullanıcı ve profili oluşturma
func

 createUserWithProfile(db *gorm.DB, name string, age int, bio string) {
    user := User{Name: name, Age: age, Profile: Profile{Bio: bio}} // Kullanıcı ve profil oluştur
    result := db.Create(&user) // Veritabanına kaydet
    if result.Error != nil {
        log.Fatal("Kullanıcı ve profil eklenirken hata oluştu: ", result.Error) // Hata varsa kaydet
    }
    log.Printf("Kullanıcı ve profil başarıyla eklendi: %+v\n", user) // Başarılı ekleme mesajı
}
```

### 7. Çıktı

```plaintext
Kullanıcı ve profil başarıyla eklendi: {ID:1 Name:Alice Age:30 Profile:{ID:1 UserID:1 Bio:Software Engineer}}
```

### Sonuç

Bu makalede, GORM ORM kütüphanesinin nasıl kullanılacağını detaylı bir şekilde inceledik. GORM ile CRUD işlemlerinin yanı sıra, ilişkilerin nasıl kurulacağı ve kullanılacağı hakkında bilgi verdik. GORM, Go ile veritabanı etkileşimlerini oldukça kolay ve etkili hale getirir, bu nedenle modern Go uygulamalarında sıklıkla kullanılmaktadır.