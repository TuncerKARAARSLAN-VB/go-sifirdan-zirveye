# Chat App

Elbette! Aşağıda, Go dilinde bir WebSocket sunucusu ve HTML ile jQuery kullanarak istemci tarafında bir mesajlaşma uygulaması geliştireceğiz. Uygulama, kullanıcıların birden fazla kişi ile mesajlaşabileceği bir mesajlaşma odası oluşturur. 

### Proje Planı

1. **Server Tarafı**:
   - Go dilinde WebSocket sunucusu oluşturacağız.
   - Kullanıcıların odaya katılmasını ve mesaj göndermesini sağlayacağız.

2. **Client Tarafı**:
   - HTML ve jQuery kullanarak kullanıcı arayüzü oluşturacağız.
   - Socket.io ile sunucuya bağlanarak mesajlaşma işlevselliğini sağlayacağız.

### Aşama 1: Go Sunucusunu Oluşturma

Öncelikle, Go ile WebSocket sunucusunu oluşturalım.

#### 1. Gerekli Kütüphanelerin Kurulumu

Aşağıdaki komutları kullanarak gerekli kütüphaneleri yükleyin:

```bash
go mod init chat-app
go get -u github.com/gorilla/websocket
```

#### 2. Sunucu Kodunu Yazma

`main.go` dosyasını oluşturun ve aşağıdaki kodu ekleyin:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

// WebSocket bağlantısı için yapı
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var clients = make(map[*Client]bool) // Bağlı istemcileri saklamak için
var broadcast = make(chan []byte)     // Mesajları yaymak için

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // CORS ayarlarını burada yapabilirsiniz
	},
}

func main() {
	// WebSocket için HTTP endpoint
	http.HandleFunc("/ws", handleConnections)

	go handleMessages() // Mesajları yönetmek için bir goroutine başlat

	fmt.Println("Sunucu 8080 portunda çalışıyor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Sunucu başlatılamadı: " + err.Error())
	}
}

// WebSocket bağlantılarını yönetme
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket bağlantısı sağlanamadı:", err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn, send: make(chan []byte)}
	clients[client] = true

	// Mesaj alma döngüsü
	go client.readMessages()

	// Mesajları yayma döngüsü
	for {
		message := <-client.send
		broadcast <- message
	}
}

// Mesajları dinleme
func (c *Client) readMessages() {
	for {
		var msg []byte
		err := c.conn.ReadMessage(&msg)
		if err != nil {
			fmt.Println("Mesaj okuma hatası:", err)
			delete(clients, c)
			break
		}
		broadcast <- msg
	}
}

// Mesajları yayma
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
	}
}
```

### Açıklamalar

- **Client Yapısı**: Her bir istemci için bağlantı ve mesaj göndermek için kullanılan bir yapı.
- **Clients Map**: Bağlı istemcileri saklamak için bir harita.
- **Broadcast Channel**: Tüm istemcilere gönderilecek mesajları saklar.
- **handleConnections**: WebSocket bağlantılarını yönetir, yeni bağlantılar ekler ve mesajları dinler.
- **handleMessages**: Gelen mesajları alır ve tüm bağlı istemcilere gönderir.

### Aşama 2: HTML ve jQuery İstemcisini Oluşturma

Şimdi, istemci tarafını oluşturacağız. `index.html` adlı bir dosya oluşturun ve aşağıdaki kodu ekleyin:

```html
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mesajlaşma Uygulaması</title>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }
        #messages {
            border: 1px solid #ccc;
            padding: 10px;
            height: 300px;
            overflow-y: scroll;
            margin-bottom: 10px;
        }
        input[type="text"] {
            width: calc(100% - 100px);
            padding: 10px;
        }
        button {
            padding: 10px;
        }
    </style>
</head>
<body>
    <h1>Mesajlaşma Uygulaması</h1>
    <div id="messages"></div>
    <input type="text" id="message" placeholder="Mesajınızı yazın...">
    <button id="send-btn">Gönder</button>

    <script>
        $(document).ready(function() {
            // WebSocket bağlantısı kurma
            const conn = new WebSocket('ws://localhost:8080/ws');

            // Mesaj alındığında
            conn.onmessage = function(e) {
                const messagesDiv = $('#messages');
                messagesDiv.append('<div>' + e.data + '</div>');
                messagesDiv.scrollTop(messagesDiv[0].scrollHeight); // En alta kaydır
            };

            // Mesaj gönderme
            $('#send-btn').click(function() {
                const message = $('#message').val();
                if (message) {
                    conn.send(message);
                    $('#message').val(''); // Giriş alanını temizle
                }
            });

            // Enter tuşuna basıldığında mesaj gönder
            $('#message').keypress(function(e) {
                if (e.which === 13) {
                    $('#send-btn').click();
                }
            });
        });
    </script>
</body>
</html>
```

### Açıklamalar

- **WebSocket Bağlantısı**: JavaScript ile sunucuya bağlanmak için bir WebSocket bağlantısı kurar.
- **Mesaj Alma**: Sunucudan gelen her mesajı dinler ve ekrana ekler.
- **Mesaj Gönderme**: Kullanıcı, mesaj giriş alanına bir mesaj yazıp "Gönder" butonuna bastığında mesajı sunucuya gönderir.
- **Enter Tuşu Desteği**: Kullanıcı, mesajı girmek için "Enter" tuşuna bastığında mesaj gönderme işlemi gerçekleştirilir.

### Aşama 3: Projeyi Çalıştırma

1. **Sunucu Çalıştırma**: Terminalden aşağıdaki komutu çalıştırarak Go sunucusunu başlatın:

   ```bash
   go run main.go
   ```

2. **HTML Dosyasını Açma**: `index.html` dosyasını bir tarayıcıda açın.

3. **Mesajlaşma**: Birden fazla tarayıcı penceresi açarak aynı odaya bağlanabilir ve mesaj gönderebilirsiniz.

### Sonuç

Bu proje, Go dilinde yazılmış bir WebSocket sunucusu ile HTML ve jQuery kullanarak yapılmış bir istemci oluşturur. Kullanıcılar, birden fazla kişi ile mesajlaşabilir ve mesajlar anlık olarak görüntülenir. Projeyi genişletmek ve geliştirmek için yeni özellikler ekleyebilirsiniz, örneğin kullanıcı isimleri, odalar, mesaj geçmişi gibi.