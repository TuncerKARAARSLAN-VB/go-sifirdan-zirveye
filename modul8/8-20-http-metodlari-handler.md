Go dilinde HTTP metodları ve bunların nasıl kullanılacağı konusunu detaylı bir şekilde ele alacağız. HTTP, istemci ve sunucu arasında veri alışverişi için kullanılan bir protokoldür ve çeşitli yöntemler (GET, POST, PUT, DELETE vb.) içerir. Bu yöntemler, sunucuya hangi tür işlem yapılacağını belirtir.

## HTTP Metodları

### 1. GET

GET isteği, sunucudan veri almak için kullanılır. Genellikle bir kaynağın durumu hakkında bilgi almak için kullanılır.

### 2. POST

POST isteği, sunucuya yeni veri göndermek için kullanılır. Örneğin, bir formun gönderilmesi veya yeni bir kaydın oluşturulması için kullanılır.

### 3. PUT

PUT isteği, mevcut bir kaynağı güncellemek için kullanılır. Sunucuya, belirtilen kaynak üzerinde değişiklik yapması için yeni veri gönderir.

### 4. DELETE

DELETE isteği, sunucudan belirtilen kaynağı silmesini ister.

### 5. PATCH

PATCH isteği, mevcut bir kaynağın kısmi olarak güncellenmesi için kullanılır.

## HTTP Handler'lar

Go dilinde, HTTP isteklerini işlemek için handler fonksiyonları kullanılır. Bu fonksiyonlar, gelen isteklere göre yanıt döner. Aşağıda, her bir HTTP yöntemini kullanan basit bir uygulama geliştireceğiz.

### Örnek Proje

Bu örnekte, bir kitap yönetim uygulaması oluşturacağız. Bu uygulama, kitapları listeleme, ekleme, güncelleme ve silme işlemlerini gerçekleştirecek.

### 1. Proje Dizini Oluşturma

Öncelikle bir proje dizini oluşturalım ve bu dizine geçelim:

```bash
mkdir go-book-api
cd go-book-api
```

### 2. Go Modülünü Başlatma

Modül oluşturmak için şu komutu çalıştırın:

```bash
go mod init go-book-api
```

### 3. Kitap Yapısı

`main.go` dosyasını oluşturun ve aşağıdaki kodu yazın:

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Book, kitap verilerini tutan bir yapıdır.
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Kitapları saklamak için bir dilim (slice) kullanıyoruz.
var books []Book
var nextID int = 1 // Yeni kitapların ID'si için bir sayaç

// GetBooks, tüm kitapları döndüren bir handler fonksiyonudur.
func GetBooks(w http.ResponseWriter, r *http.Request) {
    // Yanıtın içeriğini JSON formatında ayarlıyoruz.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books) // Kitapları JSON formatında yazıyoruz.
}

// AddBook, yeni bir kitap ekleyen bir handler fonksiyonudur.
func AddBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed) // Yöntem desteklenmiyorsa hata döndür.
        fmt.Fprintln(w, "Sadece POST yöntemi desteklenmektedir.")
        return
    }

    var book Book
    // JSON gövdesini çözümleyerek yeni bir kitap oluşturuyoruz.
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    book.ID = nextID // Kitabın ID'sini atıyoruz.
    nextID++        // Sonraki kitap için ID'yi artırıyoruz.
    books = append(books, book) // Kitabı listeye ekliyoruz.

    // Yanıtın içeriğini JSON formatında ayarlıyoruz.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // Yeni kaynak oluşturulduğunu belirt.
    json.NewEncoder(w).Encode(book)    // Eklenen kitabı yanıt olarak döndürüyoruz.
}

// UpdateBook, var olan bir kitabı güncelleyen bir handler fonksiyonudur.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        w.WriteHeader(http.StatusMethodNotAllowed) // Yöntem desteklenmiyorsa hata döndür.
        fmt.Fprintln(w, "Sadece PUT yöntemi desteklenmektedir.")
        return
    }

    var updatedBook Book
    // JSON gövdesini çözümleyerek güncellenen kitabı oluşturuyoruz.
    if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    // Kitap listesini güncellemek için ID ile eşleşen kitabı arıyoruz.
    for i, book := range books {
        if book.ID == updatedBook.ID {
            books[i].Title = updatedBook.Title // Kitap başlığını güncelle.
            books[i].Author = updatedBook.Author // Yazar bilgisini güncelle.
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(books[i]) // Güncellenen kitabı yanıt olarak döndürüyoruz.
            return
        }
    }

    // Kitap bulunamazsa 404 hatası döndür.
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintln(w, "Kitap bulunamadı.")
}

// DeleteBook, belirtilen bir kitabı silen bir handler fonksiyonudur.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        w.WriteHeader(http.StatusMethodNotAllowed) // Yöntem desteklenmiyorsa hata döndür.
        fmt.Fprintln(w, "Sadece DELETE yöntemi desteklenmektedir.")
        return
    }

    var bookToDelete Book
    // JSON gövdesinden silinecek kitabın ID'sini alıyoruz.
    if err := json.NewDecoder(r.Body).Decode(&bookToDelete); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    // Kitap listesinden belirtilen ID'ye sahip olan kitabı bulup siliyoruz.
    for i, book := range books {
        if book.ID == bookToDelete.ID {
            books = append(books[:i], books[i+1:]...) // Kitabı listeden çıkarıyoruz.
            w.WriteHeader(http.StatusNoContent) // Başarıyla silindiği için 204 döndürüyoruz.
            return
        }
    }

    // Kitap bulunamazsa 404 hatası döndür.
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintln(w, "Kitap bulunamadı.")
}

func main() {
    // HTTP sunucusunu gerekli yolları dinleyecek şekilde ayarlıyoruz.
    http.HandleFunc("/books", GetBooks)      // Tüm kitapları listelemek için GET
    http.HandleFunc("/books/add", AddBook)   // Yeni kitap eklemek için POST
    http.HandleFunc("/books/update", UpdateBook) // Var olan kitabı güncellemek için PUT
    http.HandleFunc("/books/delete", DeleteBook) // Kitabı silmek için DELETE

    // Sunucu 8080 portunda dinlemeye başlıyor.
    fmt.Println("Sunucu 8080 portunda dinliyor...")
    err := http.ListenAndServe(":8080", nil)
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

**Açıklama**: Sunucu, yeni eklenen kitabın bilgilerini yanıt olarak döndürdü. Kitabın ID'si 1 olarak atandı.

#### 5

.3. GET İsteği (Kitapları Listeleme)

Kitapları tekrar listelemek için:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[{"id":1,"title":"Go Programming","author":"John Doe"}]
```

**Açıklama**: Şimdi kitap listesi, eklediğimiz kitapla birlikte güncellendi.

#### 5.4. PUT İsteği

Kitap bilgisini güncellemek için aşağıdaki komutu kullanın:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Advanced Go Programming", "author": "John Doe"}' http://localhost:8080/books/update
```

### Çıktı

```json
{"id":1,"title":"Advanced Go Programming","author":"John Doe"}
```

**Açıklama**: Kitap güncellenmiş haliyle yanıt olarak döndürüldü.

#### 5.5. DELETE İsteği

Kitabı silmek için şu komutu kullanın:

```bash
curl -X DELETE -H "Content-Type: application/json" -d '{"id": 1}' http://localhost:8080/books/delete
```

### Çıktı

```plaintext
```

**Açıklama**: Kitap başarıyla silindiği için yanıt olarak içerik döndürülmedi. HTTP durumu 204 (No Content) döndürüldü.

#### 5.6. Kitapları Listeleme (Silmeden Sonra)

Kitapları tekrar listelemek için:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[]
```

**Açıklama**: Kitap silindiği için liste artık boştur.

## Özet

- **HTTP Metodları**: GET, POST, PUT ve DELETE metodlarının nasıl kullanıldığını öğrendik.
- **Handler Fonksiyonları**: Her bir HTTP isteği için ayrı handler fonksiyonları oluşturduk.
- **JSON İletişimi**: Sunucu ve istemci arasında JSON formatında veri alışverişi gerçekleştirdik.

Go dilinde HTTP sunucusu oluşturmak ve HTTP metodları ile nasıl çalıştığını anlamak, web uygulamaları geliştirmek için temel bir adımdır. Bu bilgilerle daha karmaşık API'ler ve uygulamalar geliştirebilirsiniz.