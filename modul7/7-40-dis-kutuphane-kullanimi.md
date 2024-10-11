Go dilinde dış kütüphane (veya modül) kullanmak, projelerinizi daha işlevsel hale getirir ve kod tekrarını azaltarak geliştirme sürecini hızlandırır. Go, `go get` komutu ile kütüphaneleri kolayca indirmenizi ve projelerinize eklemenizi sağlar. Bu bölümde, dış kütüphaneleri nasıl kullanabileceğinizi, bağımlılıkları nasıl yöneteceğinizi ve örneklerle açıklayacağım.

## Dış Kütüphanelerin Kullanımı

### 1. Dış Kütüphane Kurulumu

Bir dış kütüphane kullanmak için öncelikle `go mod init` komutunu kullanarak bir modül oluşturmalısınız. Ardından, `go get` komutu ile kütüphaneyi projeye ekleyebilirsiniz.

#### Örnek: Gorm Kütüphanesini Kullanma

**Gorm**, Go için popüler bir ORM (Object Relational Mapping) kütüphanesidir. Şimdi, Gorm'u projemize nasıl ekleyeceğimizi ve kullanacağımızı görelim.

### Adım 1: Proje Oluşturma ve Modül Başlatma

Öncelikle yeni bir proje dizini oluşturup modül başlatmalıyız.

```bash
# Yeni bir proje dizini oluştur
mkdir gorm-example
cd gorm-example

# Go modülünü başlat
go mod init github.com/kullanici/gorm-example
```

### Adım 2: Gorm Kütüphanesini Ekleme

Gorm kütüphanesini projeye eklemek için `go get` komutunu kullanıyoruz.

```bash
# Gorm kütüphanesini indir
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### Adım 3: Proje Kodu Yazma

`main.go` dosyası oluşturarak Gorm kullanarak bir veritabanı bağlantısı kurabilir ve basit bir işlem yapabiliriz.

```bash
# main.go dosyasını oluştur
touch main.go
```

Aşağıdaki kodu `main.go` dosyasına yazalım:

```go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // SQLite sürücüsü
	"gorm.io/gorm"          // Gorm kütüphanesi
)

// User modeli
type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// SQLite veritabanı bağlantısı
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}

	// User modelini veritabanında oluştur
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Model oluşturma hatası:", err)
	}

	// Yeni bir kullanıcı ekleyelim
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := db.Create(&user) // Veritabanına kullanıcıyı ekle
	if result.Error != nil {
		log.Fatal("Kullanıcı ekleme hatası:", result.Error)
	}

	// Kullanıcıyı ekledikten sonra yazdıralım
	fmt.Println("Kullanıcı eklendi:", user)
}
```

### Çıktı Açıklaması

```bash
Kullanıcı eklendi: {1 John Doe john@example.com}
```

**Açıklama**:
- `gorm.io/driver/sqlite`: SQLite veritabanı sürücüsü. Gorm ile SQLite kullanmak için bu sürücüyü ekliyoruz.
- `gorm.io/gorm`: Gorm ORM kütüphanesi. Veritabanı işlemlerimizi bu kütüphane ile gerçekleştiriyoruz.
- `User` modeli: Veritabanında bir kullanıcıyı temsil eden bir yapı.
- `gorm.Open(...)`: Veritabanına bağlanır ve bağlantı hatalarını kontrol eder.
- `db.AutoMigrate(&User{})`: Veritabanında `User` modeline karşılık gelen tabloyu otomatik olarak oluşturur.
- `db.Create(&user)`: Yeni bir kullanıcıyı veritabanına ekler.

### 2. Kütüphane Kullanımı ve Hatalar

Gorm gibi dış kütüphaneleri kullanırken, bağlantı hataları ve sorgu hataları gibi durumları kontrol etmek önemlidir. Hataları ele almak için `if err != nil` kontrolünü yapıyoruz.

### Adım 4: Kütüphaneyi Kullanma

Kütüphaneyi kullanarak daha fazla işlem yapabiliriz. Örneğin, kullanıcıları listeleyelim.

Güncellenmiş `main.go` dosyası:

```go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User modeli
type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// SQLite veritabanı bağlantısı
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}

	// User modelini veritabanında oluştur
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Model oluşturma hatası:", err)
	}

	// Yeni bir kullanıcı ekleyelim
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Kullanıcı ekleme hatası:", result.Error)
	}

	// Kullanıcıyı ekledikten sonra yazdıralım
	fmt.Println("Kullanıcı eklendi:", user)

	// Tüm kullanıcıları listeleme
	var users []User
	db.Find(&users)
	fmt.Println("Kullanıcılar:")
	for _, u := range users {
		fmt.Printf("- %s (%s)\n", u.Name, u.Email)
	}
}
```

### Çıktı Açıklaması

```bash
Kullanıcı eklendi: {1 John Doe john@example.com}
Kullanıcılar:
- John Doe (john@example.com)
```

**Açıklama**:
- `db.Find(&users)`: Veritabanındaki tüm kullanıcıları alır ve `users` dilimine kaydeder.
- Kullanıcıları listeleyerek eklediğimiz kullanıcıları görüntüledik.

### 3. Kütüphanelerin Güncellenmesi

Bir dış kütüphanenin güncellenmesi gerekiyorsa, `go get` komutunu kullanarak belirli bir sürümü veya en son sürümü alabilirsiniz.

#### Örnek: Gorm Kütüphanesinin Güncellenmesi

```bash
# Gorm kütüphanesinin en son sürümünü güncelle
go get gorm.io/gorm@latest
```

### Çıktı:

```
go: upgraded gorm.io/gorm v1.21.8 => v1.23.4
```

**Açıklama**:
- `go get gorm.io/gorm@latest`: Gorm kütüphanesinin en son sürümünü indirir ve `go.mod` dosyasını günceller.

### 4. Kütüphanelerin Silinmesi

Artık kullanmadığınız bir kütüphaneyi silmek istiyorsanız, `go mod tidy` komutunu kullanabilirsiniz.

```bash
# Kullanılmayan bağımlılıkları temizle
go mod tidy
```

### Çıktı:

```
go: removing gorm.io/gorm v1.23.4
```

**Açıklama**:
- `go mod tidy`: Projeye ekli olmayan bağımlılıkları kaldırır ve `go.mod` dosyasını temizler.

## Özet

- **Dış Kütüphane Kullanımı**: Go dilinde dış kütüphaneleri kullanarak projelerinizi daha işlevsel hale getirebilirsiniz.
- **Bağımlılık Ekleme**: `go get` komutu ile dış kütüphaneleri projeye ekleyebilirsiniz.
- **Bağımlılık Güncelleme**: `go get` ile belirli bir sürümü veya en son sürümü alabilirsiniz.
- **Bağımlılık Temizleme**: `go mod tidy` ile kullanılmayan bağımlılıkları temizleyebilirsiniz.

Go dilindeki dış kütüphane kullanımı, projelerinizi daha verimli hale getirir ve topluluk tarafından sağlanan zengin kaynaklardan yararlanmanıza olanak tanır. Gorm gibi popüler kütüphaneler, veritabanı işlemleri gibi sık karşılaşılan senaryoları kolaylaştırır.