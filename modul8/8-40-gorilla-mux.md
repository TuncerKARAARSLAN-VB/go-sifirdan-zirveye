**Gorilla Mux** kütüphanesi, Go dilinde HTTP yönlendirme işlemlerini kolaylaştıran popüler bir paketidir. RESTful API'ler geliştirirken, isteklere göre farklı yönlendirmeler yapmak için kullanışlıdır. Mux, birden fazla HTTP rotası tanımlamanıza ve bu rotalara özel işlem mantıkları eklemenize olanak tanır.

### 1. Gorilla Mux Nedir?

Gorilla Mux, Go dilinde URL yönlendirmeleri (routing) için kullanılan bir kütüphanedir. HTTP isteklerini farklı yollar üzerinden yönlendirebilir, parametreler alabilir ve URL desenleri ile çalışabilirsiniz. Bu özellikler, RESTful API'ler geliştirmek için oldukça önemlidir.

### 2. Kurulum

Gorilla Mux kullanmak için öncelikle kütüphaneyi kurmamız gerekiyor. Aşağıdaki komut ile kütüphaneyi projemize ekleyebiliriz:

```bash
go get -u github.com/gorilla/mux
```

### 3. Örnek Proje

Aşağıda, `gorilla/mux` kütüphanesini kullanarak basit bir kitap yönetimi API'si oluşturacağız.

#### 3.1 Proje Dizini Oluşturma

Proje dizinimizi oluşturalım:

```bash
mkdir gorilla-mux-example
cd gorilla-mux-example
go mod init gorilla-mux-example
```

#### 3.2 Kitap Yapısını Tanımlama

`main.go` dosyasını oluşturun ve aşağıdaki kodu yazın:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux" // Mux kütüphanesini ekliyoruz
    "net/http"
    "sync"
)

// Book yapısı, kitap bilgilerini tutmak için kullanılır.
type Book struct {
    ID     int    `json:"id"`     // Kitap ID'si
    Title  string `json:"title"`  // Kitap başlığı
    Author string `json:"author"` // Yazar adı
}

// Kitapları saklamak için bir dilim (slice) kullanıyoruz.
var books []Book
var nextID int = 1 // Yeni kitapların ID'si için bir sayaç
var mu sync.Mutex // Eş zamanlı erişim için mutex kullanıyoruz

// GetBooks, tüm kitapları döndüren bir handler fonksiyonudur.
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Yanıtın içeriğini JSON formatında ayarlıyoruz.
    mu.Lock() // Eş zamanlı erişimi sağlamak için kilitliyoruz.
    json.NewEncoder(w).Encode(books) // Kitapları JSON formatında yazıyoruz.
    mu.Unlock() // Kilidi açıyoruz.
}

// AddBook, yeni bir kitap ekleyen bir handler fonksiyonudur.
func AddBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed) // Yöntem desteklenmiyorsa hata döndür.
        fmt.Fprintln(w, "Sadece POST yöntemi desteklenmektedir.")
        return
    }

    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    mu.Lock() // Eş zamanlı erişimi sağlamak için kilitliyoruz.
    book.ID = nextID // Kitabın ID'sini atıyoruz.
    nextID++        // Sonraki kitap için ID'yi artırıyoruz.
    books = append(books, book) // Kitabı listeye ekliyoruz.
    mu.Unlock() // Kilidi açıyoruz.

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // Yeni kaynak oluşturulduğunu belirt.
    json.NewEncoder(w).Encode(book)    // Eklenen kitabı yanıt olarak döndürüyoruz.
}

// main fonksiyonu, uygulamanın başlangıç noktasıdır.
func main() {
    // Mux yönlendiricisini oluşturuyoruz.
    r := mux.NewRouter()

    // HTTP rotalarını tanımlıyoruz.
    r.HandleFunc("/books", GetBooks).Methods("GET")      // Tüm kitapları listelemek için GET
    r.HandleFunc("/books/add", AddBook).Methods("POST")  // Yeni kitap eklemek için POST

    // HTTP sunucusunu başlatıyoruz.
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        fmt.Println("Sunucu başlatılamadı:", err)
    }
}
```

### 4. Kodu Çalıştırma

Aşağıdaki komutu terminalden çalıştırarak sunucuyu başlatabilirsiniz:

```bash
go run main.go
```

### Çıktı

```plaintext
Sunucu 8080 portunda dinliyor...
```

**Açıklama**: Sunucu başarıyla başlatıldı ve 8080 portunda dinlemeye başladı.

### 5. HTTP İsteklerini Gönderme

Artık API'miz hazır olduğuna göre, HTTP isteklerini göndererek test edebiliriz.

#### 5.1. GET İsteği

Tüm kitapları listelemek için şu komutu kullanın:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[]
```

**Açıklama**: Başlangıçta kitap listesi boş olduğu için boş bir JSON dizisi döndürüldü.

#### 5.2. POST İsteği

Yeni bir kitap eklemek için aşağıdaki komutu kullanın:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Go Programming", "author": "John Doe"}' http://localhost:8080/books/add
```

### Çıktı

```json
{"id":1,"title":"Go Programming","author":"John Doe"}
```

**Açıklama**: Sunucu, yeni eklenen kitabın bilgilerini yanıt olarak döndürdü. Kitabın ID'si otomatik olarak 1 olarak belirlendi.

#### 5.3. GET İsteği (Kitapları Listeleme)

Kitapları tekrar listelemek için:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[{"id":1,"title":"Go Programming","author":"John Doe"}]
```

**Açıklama**: Şimdi kitap listesi, eklediğimiz kitapla birlikte güncellendi.

## 6. Özet

- **Gorilla Mux**: Go dilinde HTTP yönlendirme işlemlerini kolaylaştıran bir kütüphanedir.
- **HTTP Metodları**: GET ve POST metodları kullanarak kitap ekleme ve listeleme işlemleri gerçekleştirdik.
- **JSON İletişimi**: Sunucu ve istemci arasında JSON formatında veri alışverişi sağladık.
- **Eş Zamanlılık**: Eş zamanlı erişim için mutex kullandık.

Bu bilgilerle daha karmaşık API'ler ve uygulamalar geliştirebilirsiniz. `Gorilla Mux`, Go dilinde RESTful API'ler oluştururken oldukça güçlü bir araçtır.
