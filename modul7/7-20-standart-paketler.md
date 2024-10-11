Go dilinde standart kütüphane (ya da standart paketler), dilin kendisiyle birlikte gelen ve birçok yaygın işlevselliği sağlayan bir dizi pakettir. Bu paketler, temel veri yapılarından, dosya ve ağ işlemlerine kadar geniş bir yelpazeyi kapsar. Aşağıda, Go dilinin en yaygın standart paketlerini detaylı bir şekilde inceleyeceğiz. Her bir paket için örnekler ve açıklamalar da ekleyeceğiz.

## 1. fmt Paketi

`fmt` paketi, biçimlendirilmiş giriş ve çıkış işlemleri için kullanılır. Konsola yazdırmak, kullanıcıdan veri almak gibi işlemleri gerçekleştirir.

### Örnek: fmt Paketi Kullanımı

```go
package main

import (
    "fmt" // fmt paketini içe aktarıyoruz
)

func main() {
    name := "Ahmet"
    age := 25

    // Biçimlendirilmiş bir çıktı yazdırma
    fmt.Printf("Merhaba, benim adım %s ve yaşım %d.\n", name, age)
    // %s string formatı, %d integer formatı
}
```

### Çıktı:

```
Merhaba, benim adım Ahmet ve yaşım 25.
```

**Açıklama**:
- `fmt.Printf` fonksiyonu, formatlı bir çıktı verir. `%s` string ve `%d` tam sayı formatını belirtir. `\n` ise yeni bir satıra geçilmesini sağlar.

## 2. os Paketi

`os` paketi, işletim sistemi ile etkileşim kurmak için kullanılır. Dosya işlemleri, ortam değişkenleri ve işlem yönetimi gibi işlevsellikleri içerir.

### Örnek: os Paketi ile Dosya Oluşturma

```go
package main

import (
    "fmt"
    "os" // os paketini içe aktarıyoruz
)

func main() {
    // Yeni bir dosya oluşturma
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Dosya oluşturulamadı:", err)
        return
    }
    defer file.Close() // İşlem bittiğinde dosyayı kapat

    // Dosyaya yazma
    _, err = file.WriteString("Merhaba, bu bir örnek dosyadır.\n")
    if err != nil {
        fmt.Println("Dosyaya yazılamadı:", err)
    } else {
        fmt.Println("Dosyaya yazma işlemi başarılı.")
    }
}
```

### Çıktı:

```
Dosyaya yazma işlemi başarılı.
```

**Açıklama**:
- `os.Create`: Belirtilen isme sahip yeni bir dosya oluşturur. Eğer dosya zaten varsa, üzerine yazar.
- `defer file.Close()`: Dosya işlemleri tamamlandıktan sonra dosyanın kapatılmasını sağlar.
- `file.WriteString`: Dosyaya belirtilen metni yazar.

## 3. time Paketi

`time` paketi, zamanla ilgili işlemleri yönetmek için kullanılır. Tarih, saat ve zaman dilimleri gibi bilgileri içerir.

### Örnek: time Paketi Kullanımı

```go
package main

import (
    "fmt"
    "time" // time paketini içe aktarıyoruz
)

func main() {
    // Şu anki zamanı al
    currentTime := time.Now()
    
    // Zamanı biçimlendirme
    formattedTime := currentTime.Format("2006-01-02 15:04:05")
    
    fmt.Println("Şu anki zaman:", formattedTime)
}
```

### Çıktı:

```
Şu anki zaman: 2024-10-11 14:30:00
```

**Açıklama**:
- `time.Now()`: Geçerli tarihi ve saati döndürür.
- `currentTime.Format(...)`: Tarihi belirtilen formatta biçimlendirir. Go'da tarih formatında 2006, 01, 02, 15, 04, 05 sabit değerleri kullanılır.

## 4. strings Paketi

`strings` paketi, string (metin) ile ilgili çeşitli işlemler yapmaya yarar. Metin üzerinde kesme, birleştirme, arama gibi işlevler içerir.

### Örnek: strings Paketi Kullanımı

```go
package main

import (
    "fmt"
    "strings" // strings paketini içe aktarıyoruz
)

func main() {
    original := "Go programlama dilini öğreniyorum."
    
    // String'i küçük harfe çevirme
    lower := strings.ToLower(original)
    fmt.Println("Küçük harfler:", lower)

    // String'in belirli bir kelime ile başlayıp başlamadığını kontrol etme
    startsWithGo := strings.HasPrefix(original, "Go")
    fmt.Println("'Go' ile başlıyor mu?", startsWithGo) // true

    // String'i bölme
    parts := strings.Split(original, " ")
    fmt.Println("Bölünmüş string:", parts) // Kelimeleri içeren bir dilim döner
}
```

### Çıktı:

```
Küçük harfler: go programlama dilini öğreniyorum.
'Go' ile başlıyor mu? true
Bölünmüş string: [Go programlama dilini öğreniyorum.]
```

**Açıklama**:
- `strings.ToLower`: String'i küçük harflere çevirir.
- `strings.HasPrefix`: String'in belirtilen kelime ile başlayıp başlamadığını kontrol eder.
- `strings.Split`: String'i belirtilen ayırıcı ile bölerek bir dilim (slice) döner.

## 5. net/http Paketi

`net/http` paketi, HTTP istemcisi ve sunucusu oluşturmak için kullanılır. Web uygulamaları geliştirmede temel bir rol oynar.

### Örnek: Basit bir HTTP Sunucusu

```go
package main

import (
    "fmt"
    "net/http" // net/http paketini içe aktarıyoruz
)

// HTTP isteklerine yanıt veren basit bir handler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba, HTTP sunucusuna hoş geldiniz!")
}

func main() {
    // "/hello" yoluna istek geldiğinde handler fonksiyonunu çalıştır
    http.HandleFunc("/hello", handler)

    // 8080 portunda HTTP sunucusunu başlat
    fmt.Println("Sunucu başlatılıyor: http://localhost:8080/hello")
    http.ListenAndServe(":8080", nil)
}
```

### Çıktı:

```
Sunucu başlatılıyor: http://localhost:8080/hello
```

**Açıklama**:
- `http.HandleFunc`: Belirtilen yol için bir handler (işleyici) fonksiyonu atar.
- `http.ListenAndServe`: Belirtilen portta HTTP sunucusunu başlatır.
- Tarayıcıda `http://localhost:8080/hello` adresine gidildiğinde, "Merhaba, HTTP sunucusuna hoş geldiniz!" mesajı gösterilir.

## 6. encoding/json Paketi

`encoding/json` paketi, JSON formatında veri serileştirme ve serileştirmeden geri alma işlemleri için kullanılır.

### Örnek: JSON Serileştirme ve Serileştirmeden Geri Alma

```go
package main

import (
    "encoding/json" // encoding/json paketini içe aktarıyoruz
    "fmt"
)

// Kullanıcı bilgilerini temsil eden bir struct
type User struct {
    Name  string `json:"name"`  // JSON'da "name" olarak gösterilecek
    Age   int    `json:"age"`   // JSON'da "age" olarak gösterilecek
}

func main() {
    // Kullanıcı nesnesi oluşturma
    user := User{Name: "Ali", Age: 30}

    // Kullanıcıyı JSON formatına çevirme
    userJSON, err := json.Marshal(user)
    if err != nil {
        fmt.Println("JSON'a çevirirken hata:", err)
        return
    }

    fmt.Println("JSON formatındaki kullanıcı:", string(userJSON))

    // JSON'dan yeni bir kullanıcı nesnesi oluşturma
    var newUser User
    err = json.Unmarshal(userJSON, &newUser)
    if err != nil {
        fmt.Println("JSON'dan çevirirken hata:", err)
        return
    }

    fmt.Printf("Yeni kullanıcı: %+v\n", newUser)
}
```

### Çıktı:

```
JSON formatındaki kullanıcı: {"name":"Ali","age":30}
Yeni kullanıcı: &{Name:Ali Age:30}
```

**Açıklama**:
- `json.Marshal`: Belirtilen nesneyi JSON formatına çevirir.
- `json.Unmarshal`: JSON formatındaki veriyi belirtilen nesneye geri çevirir.

## 7. strconv Paketi

`strconv` paketi, string ve diğer temel veri tipleri arasında dönüşüm yapar.

### Örnek: Dönüşüm İşlemleri

```go
package main

import (
    "fmt"
    "strconv" // strconv paketini içe aktarıyoruz