Go dilinde CLI (Komut Satırı Arayüzü) uygulamaları geliştirmek için kullanılan popüler iki paket **Cobra** ve **urfave/cli**'dir. Bu paketler, karmaşık komut satırı uygulamaları oluşturmayı kolaylaştırır. Aşağıda bu iki paketin detaylı bir şekilde açıklaması, örneklerle birlikte verilmiştir.

## 1. Cobra

**Cobra**, Go dilinde güçlü ve esnek bir komut satırı arayüzü oluşturmak için kullanılan bir kütüphanedir. Kullanıcı dostu ve yapılandırılabilir bir komut yapısı sağlar.

### 1.1 Cobra Paketinin Kurulumu

Öncelikle, projenizde Cobra paketini kurmalısınız. Aşağıdaki komut ile kurabilirsiniz:

```bash
go get -u github.com/spf13/cobra@latest
```

### 1.2 Basit Cobra Uygulaması

Aşağıda, Cobra kullanarak basit bir CLI uygulaması oluşturan örnek verilmiştir. Bu uygulama bir selamlaşma komutu içermektedir.

```go
// main.go
package main

import (
	"fmt"
	"github.com/spf13/cobra" // Cobra paketini içe aktarma
	"os"
)

func main() {
	var rootCmd = &cobra.Command{ // Ana komut tanımlama
		Use:   "greet", // Komut adı
		Short: "Greet is a simple greeting application", // Kısa açıklama
		Long:  `This application allows you to greet someone with a custom message`, // Uzun açıklama
		Run: func(cmd *cobra.Command, args []string) { // Ana komutun çalıştırılacak fonksiyonu
			fmt.Println("Merhaba, Dünya!") // Varsayılan mesaj
		},
	}

	var name string // Kullanıcıdan alınacak isim değişkeni
	var greetCmd = &cobra.Command{ // greet komutu
		Use:   "sayhi", // Komut adı
		Short: "Says hi to someone", // Kısa açıklama
		Run: func(cmd *cobra.Command, args []string) { // Komutun çalıştırılacak fonksiyonu
			if name != "" {
				fmt.Printf("Merhaba, %s!\n", name) // Kullanıcıdan alınan isim ile mesaj
			} else {
				fmt.Println("Merhaba, misafir!") // Varsayılan mesaj
			}
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the person to greet") // Bayrak tanımlama
	rootCmd.AddCommand(greetCmd) // greet komutunu ana komuta ekleme

	if err := rootCmd.Execute(); err != nil { // Ana komutun çalıştırılması
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### 1.3 Çalıştırma

Yukarıdaki kodu bir dosyaya kaydettikten sonra terminalde çalıştırabilirsiniz:

```bash
go run main.go greet sayhi --name=Ali
```

### 1.4 Çıktı Açıklaması

```
Merhaba, Ali!
```

Eğer isim bayrağını vermezseniz:

```bash
go run main.go greet sayhi
```

```
Merhaba, misafir!
```

### 1.5 Cobra'nın Avantajları

- Hiyerarşik komut yapısı oluşturma
- Bayrak ve argümanları kolayca yönetme
- Yardım mesajlarının otomatik oluşturulması

## 2. Urfave/cli

**Urfave/cli** (eski adıyla `codegangsta/cli`), Go dilinde CLI uygulamaları oluşturmak için kullanılan bir başka popüler pakettir. Bu paket, uygulamanızı yapılandırmak ve komutları tanımlamak için basit bir arayüz sunar.

### 2.1 Urfave/cli Paketinin Kurulumu

Öncelikle, projenizde `urfave/cli` paketini kurmalısınız. Aşağıdaki komut ile kurabilirsiniz:

```bash
go get github.com/urfave/cli/v2
```

### 2.2 Basit Urfave/cli Uygulaması

Aşağıda, Urfave/cli kullanarak benzer bir CLI uygulaması oluşturan örnek verilmiştir.

```go
// main.go
package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2" // Urfave/cli paketini içe aktarma
)

func main() {
	app := cli.NewApp() // Yeni bir uygulama oluşturma
	app.Name = "greet" // Uygulama adı
	app.Usage = "Greet someone with a custom message" // Uygulama açıklaması

	app.Commands = []*cli.Command{ // Komutları tanımlama
		{
			Name:    "sayhi", // Komut adı
			Aliases: []string{"sh"}, // Alternatif isim
			Usage:   "Says hi to someone", // Kısa açıklama
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name", // Bayrak adı
					Aliases: []string{"n"}, // Alternatif isim
					Usage:   "The name of the person to greet", // Bayrak açıklaması
				},
			},
			Action: func(c *cli.Context) error { // Komutun çalıştırılacak fonksiyonu
				name := c.String("name") // Bayraktan isim alma
				if name != "" {
					fmt.Printf("Merhaba, %s!\n", name) // Kullanıcıdan alınan isim ile mesaj
				} else {
					fmt.Println("Merhaba, misafir!") // Varsayılan mesaj
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args) // Uygulamanın çalıştırılması
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### 2.3 Çalıştırma

Yukarıdaki kodu bir dosyaya kaydettikten sonra terminalde çalıştırabilirsiniz:

```bash
go run main.go sayhi --name=Ali
```

### 2.4 Çıktı Açıklaması

```
Merhaba, Ali!
```

Eğer isim bayrağını vermezseniz:

```bash
go run main.go sayhi
```

```
Merhaba, misafir!
```

### 2.5 Urfave/cli'nin Avantajları

- Basit ve anlaşılır API
- Bayrakların ve argümanların kolay yönetimi
- Yardım metinleri otomatik olarak oluşturulur

## 3. Cobra vs Urfave/cli

- **Cobra**: Hiyerarşik komut yapıları oluşturmayı destekler. Daha karmaşık uygulamalar için idealdir.
- **Urfave/cli**: Basit ve anlaşılır bir yapı sunar. Küçük ve orta ölçekli CLI uygulamaları için uygundur.

Her iki paket de komut satırı uygulamaları geliştirirken esneklik ve kolaylık sağlar. İhtiyacınıza göre her iki paketi de projelerinizde kullanabilirsiniz.