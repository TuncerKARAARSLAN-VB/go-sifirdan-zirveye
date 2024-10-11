# Const

Sabit değerler için `const` anahtar kelimesi kullanılır. `const` ile tanımlanan değişkenlerin değeri, program boyunca değiştirilemez. Yani sabit bir değeri temsil ederler. Sabitler genellikle sayısal değerler, string'ler veya boolean türünde olabilirler.

**Örnek: Go'da `const` kullanımı**

```go
package main

import "fmt"

const Pi = 3.14
const Greeting = "Merhaba"

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println(Greeting)
}
```

Bu örnekte `Pi` ve `Greeting` sabit olarak tanımlanmıştır ve bu değerler programın herhangi bir yerinde değiştirilemez.