# socker listener ve rabbitmq 

RabbitMQ kullanarak socket üzerinden gelen tüm verileri asenkron olarak dinleyen ve bu verileri RabbitMQ kuyruğuna yazan bir uygulama geliştirmek için Go dilini kullanabiliriz. Bu örnekte, bir TCP soketi dinleyeceğiz ve gelen verileri RabbitMQ'ya göndereceğiz. RabbitMQ, verilerin güvenli bir şekilde kuyruklandığı ve daha sonra tüketiciler tarafından alındığı bir mesajlaşma aracıdır.

### Proje Planı

1. **TCP Socket Dinleme**: TCP soketini dinleyecek ve gelen verileri alacak bir sunucu oluşturacağız.
2. **RabbitMQ Bağlantısı**: RabbitMQ sunucusuna bağlanacağız ve bir kuyruk oluşturacağız.
3. **Asenkron Veri Gönderimi**: Gelen verileri RabbitMQ kuyruğuna göndereceğiz.
4. **Hata Yönetimi**: Uygulama hataları için uygun hata yönetimi yapacağız.

### Aşama 1: Gerekli Kütüphanelerin Kurulumu

Öncelikle, Go uygulamamız için RabbitMQ ile etkileşim kurmak için gerekli olan `amqp` kütüphanesini yükleyelim.

Aşağıdaki komutu terminalde çalıştırarak kütüphaneyi indirin:

```bash
go get github.com/streadway/amqp
```

### Aşama 2: Uygulamanın Kodunu Yazma

Şimdi, TCP soketini dinleyen ve gelen verileri RabbitMQ'ya gönderen uygulamamızı yazalım. `main.go` dosyasını oluşturun ve aşağıdaki kodu ekleyin:

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "net"

    "github.com/streadway/amqp"
)

// RabbitMQ bağlantısı oluşturma
func connectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return nil, nil, fmt.Errorf("RabbitMQ bağlantısı başarısız: %s", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, nil, fmt.Errorf("Kanal oluşturulamadı: %s", err)
    }

    _, err = ch.QueueDeclare(
        "myQueue", // Kuyruk adı
        true,      // Kalıcı mı?
        false,     // Tüketim sonrası silinsin mi?
        false,     // İlk tüketici oluncaya kadar beklesin mi?
        false,     // Otomatik silinsin mi?
        nil,       // Ek parametreler
    )
    if err != nil {
        return nil, nil, fmt.Errorf("Kuyruk oluşturulamadı: %s", err)
    }

    return conn, ch, nil
}

// TCP soketini dinleme
func startTCPServer(address string, ch *amqp.Channel) {
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("TCP sunucusu başlatılamadı: %s", err)
    }
    defer listener.Close()
    log.Printf("TCP sunucusu %s adresinde dinleniyor...", address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Bağlantı hatası: %s", err)
            continue
        }
        go handleConnection(conn, ch) // Her bağlantı için yeni bir goroutine başlat
    }
}

// Bağlantıyı işleme
func handleConnection(conn net.Conn, ch *amqp.Channel) {
    defer conn.Close()
    log.Printf("Yeni bağlantı kabul edildi: %s", conn.RemoteAddr())

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        message := scanner.Text()
        log.Printf("Gelen mesaj: %s", message)

        // Mesajı RabbitMQ kuyruğuna gönder
        err := ch.Publish(
            "",         // Default exchange
            "myQueue",  // Kuyruk adı
            false,      // Durum gerektirmiyor
            false,      // Durum gerektirmiyor
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(message),
            },
        )
        if err != nil {
            log.Printf("Mesaj gönderim hatası: %s", err)
        } else {
            log.Printf("Mesaj RabbitMQ kuyruğuna gönderildi: %s", message)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Printf("Okuma hatası: %s", err)
    }
}

// Ana işlem
func main() {
    conn, ch, err := connectRabbitMQ()
    if err != nil {
        log.Fatalf("RabbitMQ bağlantısı hatası: %s", err)
    }
    defer conn.Close()
    defer ch.Close()

    address := ":8080" // Dinlenecek TCP adresi
    startTCPServer(address, ch)
}
```

### Açıklamalar

1. **connectRabbitMQ**: RabbitMQ'ya bağlanır ve bir kuyruk oluşturur. Eğer bağlantı veya kuyruk oluşturma işlemi başarısız olursa hata döner.
2. **startTCPServer**: Belirtilen adreste bir TCP sunucusu başlatır. Her yeni bağlantı için bir goroutine başlatarak `handleConnection` fonksiyonunu çağırır.
3. **handleConnection**: Bağlantıdan gelen verileri okur ve her satırda gelen veriyi RabbitMQ kuyruğuna gönderir.
4. **main**: Uygulamanın başlangıç noktasıdır. RabbitMQ'ya bağlanır ve TCP sunucusunu başlatır.

### Aşama 3: Projeyi Çalıştırma

Projeyi çalıştırmak için RabbitMQ sunucusunun çalıştığından emin olun. Ardından terminalde aşağıdaki komutu çalıştırın:

```bash
go run main.go
```

### Aşama 4: Test Etme

Sunucu çalıştığında, `telnet` veya `netcat` (nc) gibi bir araç kullanarak TCP soketine mesaj gönderebilirsiniz.

```bash
telnet localhost 8080
```

Veya:

```bash
nc localhost 8080
```

Bağlantı sağlandığında, mesajlarınızı yazıp gönderdiğinizde, bu mesajların RabbitMQ kuyruğuna gönderildiğini göreceksiniz.

### RabbitMQ Yönetim Arayüzü

RabbitMQ'yu yönetmek için RabbitMQ'nun yönetim arayüzünü kullanabilirsiniz. Web arayüzüne erişmek için tarayıcınızdan `http://localhost:15672` adresine gidin. Varsayılan kullanıcı adı ve şifre `guest` ve `guest` olacaktır. 

### Sonuç

Bu uygulama, RabbitMQ ile entegre bir TCP soketi dinleyicisi olarak çalışır. Gelen veriler asenkron bir şekilde RabbitMQ kuyruğuna gönderilir. Uygulamanızı genişletmek ve daha karmaşık iş akışları oluşturmak için bu temel yapıdan yararlanabilirsiniz.