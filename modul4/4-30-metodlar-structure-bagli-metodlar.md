Metotları ve yapıların (struct) nasıl kullanılacağını detaylandıracağız. Ayrıca, pointer ve değer üzerinden metotlar tanımlamanın önemini de vurgulayacağız.

### 1. Metotlar

Metotlar, bir nesnenin davranışlarını tanımlayan işlevlerdir. C# gibi dillerde, metotlar genellikle bir sınıf (class) veya yapı (struct) içinde tanımlanır. Metotlar, veri üzerinde işlem yapma yeteneği sağlar ve tekrar kullanılabilirliği artırır.

#### 1.1 Structlara Bağlı Metotlar Tanımlama

Struct, C# gibi dillerde veri gruplarını temsil eden bir türdür. Struct'lar, bir nesne gibi davranabilen fakat değer tipi olan veri yapılarıdır. Struct'lar için metot tanımlamak mümkündür ve bu metotlar struct'ın davranışlarını belirler.

**Örnek: Struct ile Metot Tanımlama**

```csharp
struct Rectangle
{
    public double Width;
    public double Height;

    // Alan hesaplama metodu
    public double Area()
    {
        return Width * Height;
    }

    // Çevre hesaplama metodu
    public double Perimeter()
    {
        return 2 * (Width + Height);
    }
}

// Kullanım
Rectangle rect = new Rectangle { Width = 5, Height = 10 };
Console.WriteLine("Alan: " + rect.Area());
Console.WriteLine("Çevre: " + rect.Perimeter());
```

Bu örnekte, `Rectangle` adında bir yapı oluşturduk ve bu yapıya `Area` ve `Perimeter` adında iki metot tanımladık.

#### 1.2 Pointer ve Değer Üzerinden Metotlar

Metotlar, değer tipi ve referans tipi değişkenlerle çalışabilir. Değer tipi değişkenler (örneğin, `struct` ve `int`) değerlerini kopyalayarak geçerken, referans tipi değişkenler (örneğin, `class` ve `array`) bellek adresini geçirir.

**Değer Üzerinden Metotlar:**

Değer tipleri doğrudan kopyalanır. Bu, metot çağrısında orijinal değişkenin etkilenmemesi anlamına gelir.

**Örnek:**

```csharp
void UpdateValue(int number)
{
    number = 20; // Bu değişiklik dışarıda etkili olmayacak
}

int value = 10;
UpdateValue(value);
Console.WriteLine(value); // 10 yazdırır
```

**Pointer Üzerinden Metotlar:**

Pointer veya referans tipleri kullanıldığında, orijinal değişken üzerinde değişiklik yapılabilir.

**Örnek:**

```csharp
void UpdateRectangle(ref Rectangle rect)
{
    rect.Width = 20; // Orijinal rect değişecek
}

Rectangle myRect = new Rectangle { Width = 5, Height = 10 };
UpdateRectangle(ref myRect);
Console.WriteLine("Yeni Genişlik: " + myRect.Width); // 20 yazdırır
```

### Özet

- **Struct'lar**, verileri gruplamak için kullanılır ve metotlar tanımlamak mümkündür.
- **Değer tipleri** metotlara geçerken kopyalanır; bu, orijinal değişkenin etkilenmediği anlamına gelir.
- **Referans tipleri** veya pointer kullanıldığında, orijinal değişken üzerinde değişiklik yapılabilir.

Bu modül, arayüzler ve metotların nasıl kullanılacağını anlamanıza yardımcı olacak temel bilgileri sağlamaktadır. İlerleyen konularda arayüzlerin nasıl tanımlanıp kullanıldığını inceleyeceğiz.