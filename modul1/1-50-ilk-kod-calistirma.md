Go ile "Hello, World!" yazmak, derlemek ve çalıştırmak için Visual Studio Code (VS Code) kullanarak aşağıdaki adımları izleyebilirsiniz:

### 1. Visual Studio Code'u Açın
Öncelikle, Visual Studio Code uygulamasını açın.

### 2. Go'yu Kurun
Eğer Go henüz bilgisayarınıza kurulu değilse, [Go resmi web sitesinden](https://golang.org/dl/) Go'yu indirip kurun. Kurulumdan sonra, terminalde `go version` komutunu çalıştırarak kurulumun başarılı olup olmadığını kontrol edin.

### 3. Yeni Bir Klasör Oluşturun
VS Code'da yeni bir proje klasörü oluşturun. Örneğin, `hello-world` adında bir klasör oluşturabilirsiniz.
- Terminali açın (Ctrl + `) ve aşağıdaki komutu girin:
  ```bash
  mkdir hello-world
  cd hello-world
  ```

### 4. Yeni Bir Go Dosyası Oluşturun
VS Code'da yeni bir dosya oluşturun:
- **File > New File** veya sağ üstteki **New File** ikonuna tıklayın.
- Dosya adını `hello-world.go` olarak belirleyin.

### 5. "Hello, World!" Kodunu Yazın
`hello-world.go` dosyasına aşağıdaki kodu ekleyin:
```go
package hello-world

import "fmt"

func hello-world() {
    fmt.Println("Hello, World!")
}
```

### 6. Dosyayı Kaydedin
Dosyayı kaydetmek için `Ctrl + S` tuş kombinasyonunu kullanın veya **File > Save** seçeneğine tıklayın.

### 7. Terminali Açın
VS Code'da bir terminal açın:
- **Terminal > New Terminal** seçeneğine tıklayın.

### 8. Go Programını Derleyin
Aşağıdaki komutu terminalde çalıştırarak programınızı derleyin:
```bash
go build hello-world.go
```
Bu komut, `hello-world.go` dosyanızı derler ve çalışma dosyası oluşturur.

### 9. Programı Çalıştırın
Derleme işlemi başarılı ise, aşağıdaki komutla programınızı çalıştırabilirsiniz:
- Windows için:
  ```bash
  .\hello-world.exe
  ```
- Linux veya macOS için:
  ```bash
  ./hello-world
  ```

### 10. Çıktıyı Kontrol Edin
Eğer her şey doğru bir şekilde yapıldıysa, terminalde aşağıdaki gibi bir çıktı görmelisiniz:
```
Hello, World!
```

### Özet
Bu adımlar sayesinde Visual Studio Code kullanarak Go ile "Hello, World!" programını yazdınız, derlediniz ve çalıştırdınız. Go dilinin diğer özelliklerini keşfetmek için yeni dosyalar ve projeler oluşturmaya devam edebilirsiniz!