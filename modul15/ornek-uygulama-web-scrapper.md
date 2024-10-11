# Web Scraper

Go ile bir web tarayıcı (web scraper) geliştirmek için, bir URL'den başlayarak alt bağlantıları tespit edip bu bağlantılara erişip içeriklerini indirecek bir uygulama oluşturabiliriz. Bu tür bir projede, Go'nun `net/http` ve `golang.org/x/net/html` paketlerini kullanacağız. Bu paketler, HTTP istekleri göndermek ve HTML belgelerini analiz etmek için kullanılabilir.

### Proje Planı

1. **HTTP İstekleri Gönderme**: Belirtilen URL'ye HTTP isteği gönderilecek.
2. **HTML Analizi**: HTML içeriği, alt bağlantıları tespit etmek için analiz edilecek.
3. **Alt Bağlantılara Erişim**: Tespit edilen her alt bağlantıya gidilecek ve içeriği indirilecektir.
4. **İçeriklerin Kaydedilmesi**: İndirilen içerikler yerel bir dosyaya kaydedilecektir.

### Aşama 1: Gerekli Kütüphanelerin Kurulumu

Go ile web scrapper geliştirmek için herhangi bir harici kütüphane yüklememize gerek yok, yalnızca standart kütüphaneleri kullanacağız. Ancak `golang.org/x/net/html` paketini kullanmak için onu indirmeniz gerekecek.

Aşağıdaki komutu terminalde çalıştırarak kütüphaneyi indirin:

```bash
go get golang.org/x/net/html
```

### Aşama 2: Scraper Uygulaması Yazma

Şimdi Go uygulamamızı oluşturalım. `main.go` dosyasını oluşturun ve aşağıdaki kodu ekleyin:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// URL'den HTML içeriği indir
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP hata: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Alt bağlantıları bul
func extractLinks(n *html.Node, baseURL string) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link := attr.Val
				// Bağlantının tam URL'ye dönüştürülmesi
				if !strings.HasPrefix(link, "http") {
					link = baseURL + link
				}
				links = append(links, link)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c, baseURL)...)
	}

	return links
}

// İçeriği kaydet
func saveContent(url, content string) {
	fileName := strings.ReplaceAll(url, "https://", "")
	fileName = strings.ReplaceAll(fileName, "http://", "")
	fileName = strings.ReplaceAll(fileName, "/", "_") + ".html"

	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Dosya kaydedilemedi: %s\n", err)
		return
	}
	fmt.Printf("İçerik kaydedildi: %s\n", fileName)
}

// URL'den içerik indir ve kaydet
func downloadContent(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("İçerik indirilemedi: %s\n", err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("İçerik okunamadı: %s\n", err)
		return
	}

	saveContent(url, string(content))
}

// Ana işlem
func main() {
	url := "https://example.com" // Değiştirilmesi gereken URL
	fmt.Printf("Tarayıcı başlatılıyor: %s\n", url)

	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Printf("Hata: %s\n", err)
		return
	}

	links := extractLinks(doc, url)
	fmt.Printf("Bulunan bağlantılar: %v\n", links)

	// Her alt bağlantıyı indir
	for _, link := range links {
		downloadContent(link)
	}
}
```

### Açıklamalar

1. **fetchHTML**: Belirtilen URL'ye bir HTTP GET isteği gönderir ve HTML içeriği döner.
2. **extractLinks**: HTML ağacını gezerek tüm `<a>` etiketlerini bulur ve `href` özniteliklerini toplar. Eğer bir bağlantı tam bir URL değilse, `baseURL` ile birleştirir.
3. **saveContent**: İndirilen içeriği, bağlantının URL'sini dosya adı olarak kullanarak kaydeder.
4. **downloadContent**: Belirtilen URL'den içerik indirir ve `saveContent` fonksiyonunu çağırarak kaydeder.
5. **main**: Programın başlangıç noktasıdır. Kullanıcıdan URL alır, içeriği indirir ve alt bağlantıları bulur.

### Aşama 3: Projeyi Çalıştırma

Projeyi çalıştırmak için terminalde aşağıdaki komutu çalıştırın:

```bash
go run main.go
```

`main.go` dosyasında bulunan URL'yi değiştirmeyi unutmayın. Örneğin, `https://example.com` yerine başka bir URL yazabilirsiniz. Program çalıştığında, belirtilen URL'den tüm alt bağlantıları bulacak ve bu bağlantılardaki içerikleri yerel dosyalara kaydedecektir.

### Dikkat Edilmesi Gerekenler

1. **Robots.txt**: Tarayıcı uygulamanızın etik kurallara uygun olduğundan emin olun. `robots.txt` dosyasını kontrol ederek web sitesinin taranmasına izin verip vermediğini öğrenin.
2. **Hedef Web Sitesinin Politikasını Kontrol Edin**: Herhangi bir web sitesini taramadan önce, o web sitesinin kullanım şartlarını ve veri çekme politikasını kontrol edin.
3. **Hız Sınırlaması**: Çok fazla istek göndermemek için isteklerinizi yavaşlatın veya bir gecikme ekleyin. Bu, sunucuya aşırı yüklenmeyi önler ve IP adresinizin engellenme olasılığını azaltır.

Bu temel scraper, daha karmaşık projelerde geliştirilmek üzere bir başlangıç sağlar. Daha fazla özellik eklemek (örneğin, derinlik sınırlaması, daha iyi hata işleme) projenin işlevselliğini artırabilir.