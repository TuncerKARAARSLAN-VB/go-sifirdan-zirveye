# Go Sıfırdan Zirveye

Bu eğitimin amacı, go dilini hiç bilmeyen bir kişinin kendi başına sıfırdan zirveye çıkmasını sağlaması için hazırlanmıştır.

Eğitimin ajandası dolu dolu hazırlandı. Bu eğtime başlarken hiç bilmediğinizi varsayacağız. Ancak kısa kısa eğitimler ile hızlı bir eğitim içeriği tamamlayacağız. Eğitim bittiğinde shell arayüzlerinden bulut tabanlı sistemlerde sizlere go dili ile proje geliştirme yetkinliği kazandıracağım. 

Eğitim içeriği her modül için ayrı klasörler hazırlandı. Her klasör içerisinde eğitim içerikleri, uygulamalar ve karşılaşılan hatar yer almaktadır.

Şimdi ajandaya bir göz atalım.

Aşağıda, Go Dili Eğitimi İçerik Ajandası'nın markdown formatında düzenlenmiş hali bulunmaktadır:

# Go Dili Eğitimi İçerik Ajandası

## Modül 1: Go'ya Giriş ve Kurulum
- Go dilinin genel özellikleri ve tarihçesi.
- Go’nun güçlü yönleri ve kullanım alanları.
- Go dilini bilgisayara kurma (Windows, Linux, macOS).
- İlk Go programı yazma ve çalıştırma.
  - Go çalışma zamanı (runtime) nasıl çalışır?
  - Go'nun derleme ve çalıştırma süreci.

## Modül 2: Temel Go Yapıları
- **Değişkenler ve Veri Tipleri**
  - Değişken tanımlama, var, kısa tanımlama (:=) ve veri tipleri.
  - Sayısal, karakter, mantıksal ve diziler.
- **Operatörler ve İfadeler**
  - Matematiksel, karşılaştırma ve mantıksal operatörler.
- **Kontrol Yapıları**
  - if, else, switch, for döngüsü.
  - Go’da switch-case farklılıkları.
- **Fonksiyonlar**
  - Fonksiyon tanımlama ve kullanma.
  - Parametreler ve dönüş değerleri.
  - Çoklu dönüş değeri kullanımı.

## Modül 3: Veri Yapıları ve Gösterim
- **Array ve Slices**
  - Array tanımlama ve kullanımı.
  - Slices ve alt dilimler.
  - Slices kapasitesi ve dinamik yapılar.
- **Maps (Sözlükler)**
  - Map veri yapısı: anahtar-değer ilişkisi.
  - Map işlemleri: ekleme, silme ve arama.
- **Structs**
  - Go’da yapı (struct) tanımlama.
  - İç içe struct kullanımı ve struct ile metot tanımlama.

## Modül 4: Arayüzler (Interfaces) ve Metotlar
- **Arayüz Kavramı**
  - Arayüzlerin Go dilindeki yeri.
  - Arayüz tanımlama ve implementasyonu.
  - Dinamik tipler ve type assertion.
- **Metotlar**
  - Structlara bağlı metotlar tanımlama.
  - Pointer ve değer üzerinden metotlar.

## Modül 5: Eş Zamanlı Programlama (Concurrency)
- **Goroutine'ler**
  - Goroutine kavramı ve Go’da eş zamanlı işlemler.
  - Goroutine’lerin başlatılması ve yönetimi.
- **Kanallar (Channels)**
  - Kanal tanımı, kullanımı ve eş zamanlı işlemler arası veri paylaşımı.
  - select yapısı ile kanalları yönetme.
  - Kanallarda bloklama ve asenkron iletişim.

## Modül 6: Hata Yönetimi
- **Go’da Hatalar**
  - error arayüzü ve hata yönetimi.
  - Özel hata türleri ve hata oluşturma.
  - defer, panic ve recover kullanımı.

## Modül 7: Go Paket Yönetimi
- **Paketler ve Modüller**
  - Paket nedir? Paket oluşturma ve kullanma.
  - Standart paketler (fmt, net/http, os, vb.).
  - Modül sistemi: go mod init, go mod tidy, go get.
- **Dış Kütüphaneler Kullanımı**
  - Go modülü ile harici kütüphaneleri projeye dahil etme.
  - go get ve bağımlılık yönetimi.

## Modül 8: Go ile Web Uygulamaları Geliştirme
- **HTTP Sunucusu Kurma**
  - net/http paketi ile basit bir web sunucusu.
  - HTTP metotları ve handler fonksiyonları.
- **RESTful API Geliştirme**
  - Temel REST API yapısı.
  - JSON ile veri işleme (encoding/decoding).
  - gorilla/mux kullanarak route yönetimi.
- **Middleware Kavramı**
  - Middleware tanımlama ve kullanma.

## Modül 9: Veritabanı Bağlantısı
- **Go ile Veritabanı Bağlantısı**
  - database/sql paketi ile MySQL/PostgreSQL bağlantısı.
  - CRUD işlemleri (Oluşturma, Okuma, Güncelleme, Silme).
  - SQL sorguları ve işlem (transaction) yönetimi.
- **ORM Kullanımı**
  - GORM ile ORM entegrasyonu ve veri modelleme.

## Modül 10: Go Test Yazımı ve Benchmarking
- **Test Yazma**
  - testing paketi ile unit test yazma.
  - Test fonksiyonları ve test dosyaları oluşturma.
  - Mocklama ve hata yakalama.
- **Benchmark Testleri**
  - Performans testleri ve go test -bench kullanımı.
  - Kodun performansını ölçme ve optimizasyon.

## Modül 11: Go ile CLI Uygulamaları Geliştirme
- **CLI Uygulama Geliştirme**
  - Go ile basit bir komut satırı uygulaması (CLI) geliştirme.
  - flag paketi ile parametre yönetimi.
  - cobra veya urfave/cli gibi frameworkler ile daha gelişmiş CLI uygulamaları.

## Modül 12: Go ile Mikroservis Mimarisi
- **Mikroservis Kavramı**
  - Mikroservis mimarisi nedir ve avantajları nelerdir?
  - Go ile mikroservis geliştirme adımları.
- **RPC ve gRPC**
  - gRPC ile Go'da RPC (Remote Procedure Call) uygulamaları.
  - Protobuf ile veri serileştirme.

## Modül 13: Docker ve Go Uygulamalarını Kapsülleme
- **Go Uygulamalarını Docker ile Kapsülleme**
  - Dockerfile oluşturma ve Go uygulamasını Docker imajı haline getirme.
  - Çok aşamalı (multi-stage) Docker imajı oluşturma.
- **Docker Compose ile Go Uygulamaları**
  - Docker Compose kullanarak bir Go uygulaması ve veritabanı entegrasyonu.

## Modül 14: Go ile Dağıtım ve CI/CD Süreçleri
- **CI/CD ile Otomatik Dağıtım**
  - GitHub Actions, TravisCI veya Jenkins ile Go projelerini otomatik olarak derleme ve test etme.
- **Go Uygulamalarını Bulutta Dağıtma**
  - AWS, GCP, Azure gibi platformlara Go uygulamalarını dağıtma.
  - Kubernetes ile Go uygulamalarını yönetme.

## Modül 15: Büyük Proje: Tamamlayıcı Uygulama
- **Katılımcıların öğrendikleri konuları bir arada kullanarak büyük bir uygulama geliştirme**
  - (örneğin, bir e-ticaret platformu, blog sistemi veya kullanıcı yönetim sistemi).
  - Proje planlaması ve tasarımı.
  - Ekip çalışması ile uygulama geliştirme ve kod inceleme.

## Eğitim Süresi ve Yöntem
Eğitimi şu formatlarda sunabilirsiniz:
- Günlük/Haftalık Dersler: Her modülü bir haftaya yayarak derinlemesine işleyebilirsiniz.
- Atölye Çalışmaları: Katılımcıların pratik yapmasına olanak tanıyan uygulamalı oturumlar düzenleyebilirsiniz.
- Online Video Serisi: Her modülü kısa video dersler halinde hazırlayarak, katılımcıların kendi hızlarında ilerlemelerini sağlayabilirsiniz.
