Go dilinde bağımlılık yönetimi, `go get` komutu ile sağlanır. Bu komut, projelerinize kütüphaneler eklemenizi, güncellemelerinizi ve kaldırmanızı kolaylaştırır. Bu yazıda, `go get` komutunun nasıl kullanılacağını ve bağımlılıkların nasıl yönetileceğini detaylı olarak inceleyeceğiz.

## 1. Go Modülleri ve `go get` Komutu

Go, modüler bir yapıya sahiptir. Bu, projelerinizi ve bağımlılıklarını yönetmek için `go.mod` ve `go.sum` dosyalarını kullanır. `go get` komutu, bu modül dosyalarını güncelleyerek bağımlılıkları indirir ve yönetir.

### 1.1. Proje Oluşturma

Öncelikle yeni bir Go projesi oluşturup modül başlatmamız gerekiyor. Aşağıdaki adımları takip ederek başlayalım.

#### Adım 1: Proje Dizini Oluşturma

```bash
# Yeni bir proje dizini oluştur
mkdir go-get-example
cd go-get-example
```

#### Adım 2: Modül Başlatma

Go modülünü başlatmak için `go mod init` komutunu kullanıyoruz. Bu komut, proje dizininde `go.mod` dosyasını oluşturur.

```bash
# Go modülünü başlat
go mod init github.com/kullanici/go-get-example
```

### 2. `go get` ile Bağımlılık Ekleme

Bir kütüphane eklemek için `go get` komutunu kullanırız. Bu, belirli bir kütüphaneyi indirir ve projenin bağımlılıkları arasında ekler.

#### Örnek: Gorm Kütüphanesini Ekleme

`Gorm` kütüphanesini ekleyerek SQLite veritabanı ile çalışacağız. 

```bash
# Gorm kütüphanesini ve SQLite sürücüsünü ekleyelim
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### Çıktı

```plaintext
go: added gorm.io/gorm v1.23.4
go: added gorm.io/driver/sqlite v1.0.0
```

**Açıklama**:
- `go get gorm.io/gorm`: Gorm kütüphanesini projeye ekler.
- `go get gorm.io/driver/sqlite`: SQLite sürücüsünü projeye ekler.

### 3. `go.mod` ve `go.sum` Dosyaları

`go get` komutunu çalıştırdığınızda, `go.mod` ve `go.sum` dosyaları güncellenir.

#### 3.1. `go.mod` Dosyası

Bu dosya, projenizin bağımlılıklarını ve modül bilgilerini tutar. Örneğin, yukarıdaki `go get` komutlarından sonra `go.mod` dosyanız aşağıdaki gibi görünebilir:

```go
module github.com/kullanici/go-get-example

go 1.20

require (
    gorm.io/driver/sqlite v1.0.0
    gorm.io/gorm v1.23.4
)
```

**Açıklama**:
- `module github.com/kullanici/go-get-example`: Projenizin modül adı.
- `require`: Projeye eklenen bağımlılıklar.

#### 3.2. `go.sum` Dosyası

`go.sum` dosyası, bağımlılıkların sürüm kontrolünü sağlar ve her bağımlılığın checksum'unu içerir. Bu, güvenli bir şekilde bağımlılıkların doğrulanmasını sağlar.

```plaintext
gorm.io/driver/sqlite v1.0.0 h1:XXXX...
gorm.io/gorm v1.23.4 h1:XXXX...
```

**Açıklama**:
- Her bağımlılığın yanında bir `h1:` değeri bulunur. Bu değer, o bağımlılığın checksum'unu ifade eder. Projenizin doğru ve güvenli bir şekilde çalışması için gereklidir.

### 4. Bağımlılık Güncelleme

Bir kütüphanenin güncellenmesi gerekiyorsa, `go get` komutunu kullanarak belirli bir sürümü veya en son sürümü alabilirsiniz.

#### Örnek: Gorm Kütüphanesinin Güncellenmesi

```bash
# Gorm kütüphanesinin en son sürümünü güncelle
go get gorm.io/gorm@latest
```

### Çıktı

```plaintext
go: upgraded gorm.io/gorm v1.23.4 => v1.23.5
```

**Açıklama**:
- `go get gorm.io/gorm@latest`: Gorm kütüphanesinin en son sürümünü indirir ve `go.mod` dosyasını günceller.

### 5. Bağımlılık Silme

Artık kullanmadığınız bir kütüphaneyi silmek istiyorsanız, `go mod tidy` komutunu kullanabilirsiniz. Bu komut, kullanılmayan bağımlılıkları temizler ve `go.mod` dosyasını günceller.

```bash
# Kullanılmayan bağımlılıkları temizle
go mod tidy
```

### Çıktı

```plaintext
go: removing gorm.io/gorm v1.23.5
```

**Açıklama**:
- `go mod tidy`: Projeye ekli olmayan bağımlılıkları kaldırır ve `go.mod` dosyasını temizler.

### 6. Belirli Sürümleri Yükleme

Bağımlılıkların belirli sürümlerini yüklemek için `@` sembolü ile sürüm numarasını belirtebilirsiniz. 

#### Örnek: Belirli Bir Sürümü Ekleme

```bash
# Belirli bir sürüm yükleme
go get gorm.io/gorm@v1.21.0
```

### Çıktı

```plaintext
go: gorm.io/gorm v1.21.0 h1:XXXX...
```

**Açıklama**:
- `go get gorm.io/gorm@v1.21.0`: Gorm kütüphanesinin `v1.21.0` sürümünü indirir.

## Özet

- **`go get` Komutu**: Dış kütüphaneleri projeye eklemek, güncellemek ve kaldırmak için kullanılır.
- **Modül ve Bağımlılıklar**: `go.mod` ve `go.sum` dosyaları, projelerin bağımlılıklarını yönetir.
- **Güncelleme ve Temizlik**: `go get` ile bağımlılıkları güncelleyebilir, `go mod tidy` ile kullanılmayanları temizleyebilirsiniz.
- **Belirli Sürümler**: Dış kütüphanelerin belirli sürümlerini yüklemek için `@` sembolü kullanılır.

Bu şekilde Go dilinde bağımlılık yönetimi, projelerinizi daha düzenli ve yönetilebilir hale getirir. `go get` komutunu kullanarak dış kütüphaneler ile çalışmak, kod geliştirme sürecinizi hızlandırır ve işlevselliği artırır.