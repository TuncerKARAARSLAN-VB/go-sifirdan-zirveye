# fmt Fonksiyonları: Fonksiyon Listesi

Go dilindeki `fmt` paketi, girdileri formatlama, çıktı yazdırma ve okuma gibi işlevler için çeşitli fonksiyonlar sunar. Aşağıda `fmt` paketinde yaygın olarak kullanılan bazı fonksiyonları listeleyen bir tablo bulunuyor:

| Fonksiyon          | Açıklama                                                                 |
|--------------------|--------------------------------------------------------------------------|
| **`fmt.Print`**     | Argümanları yazdırır, yeni satıra geçmez.                                |
| **`fmt.Println`**   | Argümanları yazdırır ve yeni satıra geçer.                               |
| **`fmt.Printf`**    | Argümanları belirtilen biçimde (formatlı) yazdırır.                      |
| **`fmt.Sprintf`**   | Argümanları belirtilen biçimde bir string olarak döner.                  |
| **`fmt.Sprint`**    | Argümanları bir string olarak döner.                                     |
| **`fmt.Sprintln`**  | Argümanları bir string olarak döner ve yeni satıra geçer.                |
| **`fmt.Fprint`**    | Argümanları belirtilen `io.Writer` (örneğin bir dosya) üzerine yazar.    |
| **`fmt.Fprintln`**  | Argümanları belirtilen `io.Writer` üzerine yazar ve yeni satıra geçer.   |
| **`fmt.Fprintf`**   | Argümanları belirtilen `io.Writer` üzerine formatlı yazar.               |
| **`fmt.Scan`**      | Kullanıcıdan girdi alır, boşlukla ayrılan değerleri okur.                |
| **`fmt.Scanln`**    | Kullanıcıdan girdi alır, yeni satıra kadar olan değerleri okur.          |
| **`fmt.Scanf`**     | Formatlı bir şekilde kullanıcıdan girdi alır.                            |
| **`fmt.Sscan`**     | Bir string'den boşlukla ayrılan değerleri okur.                          |
| **`fmt.Sscanln`**   | Bir string'den yeni satıra kadar olan değerleri okur.                    |
| **`fmt.Sscanf`**    | Bir string'den formatlı değerleri okur.                                  |
| **`fmt.Fscan`**     | Belirtilen `io.Reader`'dan boşlukla ayrılan değerleri okur.              |
| **`fmt.Fscanln`**   | Belirtilen `io.Reader`'dan yeni satıra kadar olan değerleri okur.        |
| **`fmt.Fscanf`**    | Belirtilen `io.Reader`'dan formatlı değerleri okur.                      |

## Fonksiyon Açıklamaları:

1. **Yazdırma Fonksiyonları**: 
   - `Print`, `Println`, ve `Printf` gibi fonksiyonlar, ekrana yazdırma işlevi görürler.
   - `Fprint`, `Fprintln`, `Fprintf` ise çıktı hedefini belirleyebileceğiniz (`io.Writer` arayüzüne sahip olan) fonksiyonlardır. Örneğin dosya yazımı.
   - `Sprint`, `Sprintln`, ve `Sprintf` ise yazdırılan çıktıyı bir `string` olarak döner.

2. **Girdi Alma Fonksiyonları**:
   - `Scan`, `Scanln`, ve `Scanf`, kullanıcıdan veya standart girdiden (örneğin komut satırından) veri almak için kullanılır.
   - `Sscan`, `Sscanln`, ve `Sscanf` ise bir `string` girdiden veri okur.
   - `Fscan`, `Fscanln`, ve `Fscanf`, belirtilen `io.Reader`'dan veri okumak için kullanılır.
