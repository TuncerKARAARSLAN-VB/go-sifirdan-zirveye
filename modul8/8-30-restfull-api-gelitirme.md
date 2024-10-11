Go dilinde RESTful API geliştirmek, modern web uygulamaları için oldukça önemli bir beceridir. REST (Representational State Transfer), HTTP protokolü üzerinden istemci-sunucu iletişimi kuran bir mimaridir. Bu yazıda, Go dilinde basit bir RESTful API geliştirerek temel kavramları açıklayacağız. Örneklerimiz kitap yönetimi üzerine olacak.

## RESTful API Nedir?

RESTful API, istemcilerin sunucu ile HTTP yöntemlerini (GET, POST, PUT, DELETE) kullanarak etkileşimde bulunmasına olanak tanır. Bu API'ler, kaynakları temsil eder ve bu kaynaklar üzerinde çeşitli işlemler yapmayı sağlar.

### Temel HTTP Metodları

- **GET**: Belirtilen kaynakları almak için kullanılır.
- **POST**: Yeni kaynak eklemek için kullanılır.
- **PUT**: Mevcut bir kaynağı güncellemek için kullanılır.
- **DELETE**: Belirtilen bir kaynağı silmek için kullanılır.

## Proje Kurulumu

### 1. Proje Dizini Oluşturma

Öncelikle bir proje dizini oluşturalım:

```bash
mkdir go-book-api
cd go-book-api
```

### 2. Go Modülünü Başlatma

Modül oluşturmak için şu komutu çalıştırın:

```bash
go mod init go-book-api
```

### 3. Gerekli Paketleri Ekleyin

Bu proje için `net/http` ve `encoding/json` paketlerini kullanacağız. Go dilinin standart kütüphanelerinde bu paketler mevcut olduğu için ek bir yükleme yapmamıza gerek yoktur.

### 4. Kitap Yapısı Tanımlama

`main.go` dosyasını oluşturun ve aşağıdaki kodu yazın:

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
)

// Book, kitap verilerini tutan bir yapıdır.
type Book struct {
    ID     int    `json:"id"`     // Kitap ID'si
    Title  string `json:"title"`  // Kitap başlığı
    Author string `json:"author"` // Yazar adı
}

// Kitapları saklamak için bir dilim (slice) kullanıyoruz.
var books []Book
var nextID int = 1 // Yeni kitapların ID'si için bir sayaç
var mu sync.Mutex // Eş zamanlı erişim için mutex kullanıyoruz.

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

// UpdateBook, var olan bir kitabı güncelleyen bir handler fonksiyonudur.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        w.WriteHeader(http.StatusMethodNotAllowed) // Yöntem desteklenmiyorsa hata döndür.
        fmt.Fprintln(w, "Sadece PUT yöntemi desteklenmektedir.")
        return
    }

    var updatedBook Book
    if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    mu.Lock() // Eş zamanlı erişimi sağlamak için kilitliyoruz.
    for i, book := range books {
        if book.ID == updatedBook.ID {
            books[i].Title = updatedBook.Title // Kitap başlığını güncelle.
            books[i].Author = updatedBook.Author // Yazar bilgisini güncelle.
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(books[i]) // Güncellenen kitabı yanıt olarak döndürüyoruz.
            mu.Unlock() // Kilidi açıyoruz.
            return
        }
    }
    mu.Unlock() // Kilidi açıyoruz.

    w.WriteHeader(http.StatusNotFound) // Kitap bulunamazsa 404 hatası döndür.
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
    if err := json.NewDecoder(r.Body).Decode(&bookToDelete); err != nil {
        w.WriteHeader(http.StatusBadRequest) // Hatalı istek durumunda hata döndür.
        fmt.Fprintln(w, "Geçersiz istek gövdesi.")
        return
    }

    mu.Lock() // Eş zamanlı erişimi sağlamak için kilitliyoruz.
    for i, book := range books {
        if book.ID == bookToDelete.ID {
            books = append(books[:i], books[i+1:]...) // Kitabı listeden çıkarıyoruz.
            w.WriteHeader(http.StatusNoContent) // Başarıyla silindiği için 204 döndürüyoruz.
            mu.Unlock() // Kilidi açıyoruz.
            return
        }
    }
    mu.Unlock() // Kilidi açıyoruz.

    w.WriteHeader(http.StatusNotFound) // Kitap bulunamazsa 404 hatası döndür.
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

### 5. Kodu Çalıştırma

Aşağıdaki komutu terminalden çalıştırarak sunucuyu başlatabilirsiniz:

```bash
go run main.go
```

### Çıktı

```plaintext
Sunucu 8080 portunda dinliyor...
```

**Açıklama**: Sunucu başarıyla başlatıldı ve 8080 portunda dinlemeye başladı.

### 6. HTTP İsteklerini Gönderme

Artık API'miz hazır olduğuna göre, HTTP isteklerini göndererek test edebiliriz.

#### 6.1. GET İsteği

Tüm kitapları listelemek için şu komutu kullanın:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[]
```

**Açıklama**: Başlangıçta kitap listesi boş olduğu için boş bir JSON dizisi döndürüldü.

#### 6.2. POST İsteği

Yeni bir kitap eklemek için aşağıdaki komutu kullanın:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Go Programming", "author": "John Doe"}' http://localhost:8080/books/add
```

### Çıktı

```json
{"id":1,"title":"Go Programming","author":"John Doe"}
```

**Açıklama**: Sunucu, yeni eklenen kitabın bilgilerini yanıt olarak döndürdü. Kitabın

 ID'si otomatik olarak 1 olarak belirlendi.

#### 6.3. GET İsteği (Kitapları Listeleme)

Kitapları tekrar listelemek için:

```bash
curl http://localhost:8080/books
```

### Çıktı

```json
[{"id":1,"title":"Go Programming","author":"John Doe"}]
```

**Açıklama**: Şimdi kitap listesi, eklediğimiz kitapla birlikte güncellendi.

#### 6.4. PUT İsteği

Kitap bilgisini güncellemek için aşağıdaki komutu kullanın:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Advanced Go Programming", "author": "John Doe"}' http://localhost:8080/books/update
```

### Çıktı

```json
{"id":1,"title":"Advanced Go Programming","author":"John Doe"}
```

**Açıklama**: Kitap güncellenmiş haliyle yanıt olarak döndürüldü.

#### 6.5. DELETE İsteği

Kitabı silmek için şu komutu kullanın:

```bash
curl -X DELETE -H "Content-Type: application/json" -d '{"id": 1}' http://localhost:8080/books/delete
```

### Çıktı

```plaintext
```

**Açıklama**: Kitap başarıyla silindiği için yanıt olarak içerik döndürülmedi. HTTP durumu 204 (No Content) döndürüldü.

#### 6.6. Kitapları Listeleme (Silmeden Sonra)

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

- **RESTful API**: Go dilinde basit bir kitap yönetim API'si geliştirdik.
- **HTTP Metodları**: GET, POST, PUT ve DELETE metodlarını kullanarak veri ekleme, güncelleme, silme ve listeleme işlemlerini gerçekleştirdik.
- **JSON İletişimi**: Sunucu ve istemci arasında JSON formatında veri alışverişi gerçekleştirdik.
- **Concurrency (Eş Zamanlılık)**: Eş zamanlı erişim için mutex kullandık.

Bu bilgilerle daha karmaşık API'ler ve uygulamalar geliştirebilirsiniz. RESTful API tasarımı, modern web uygulamalarında sıklıkla kullanılan bir yaklaşımdır.