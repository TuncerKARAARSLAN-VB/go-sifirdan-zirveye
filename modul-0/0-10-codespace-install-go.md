# GitHub Codespaces

## 1. Yeni Bir Codespace Oluşturun

1. GitHub'daki depoya gidin.
2. Yeşil **Code** butonuna tıklayın.
3. **Open with Codespaces** seçeneğini ardından **New codespace**'i seçin.

## 2. Terminali Açın

Codespace oluşturulduktan ve açıldıktan sonra, bir terminale erişiminiz olacak. Terminali açmak için üst menüden **Terminal** seçeneğini ve ardından **New Terminal**'ı seçin.

## 3. Go'yu Kurun

Go'yu kurmak için aşağıdaki komutları kullanın:

### Debian tabanlı dağıtımlar (örneğin, Ubuntu) için:

```bash
sudo apt update
sudo apt install golang-go
```

### Diğer dağıtımlar için

En son sürümü istiyorsanız veya paket yöneticisi aracılığıyla mevcut değilse, Go'yu doğrudan resmi web sitesinden indirebilirsiniz:

```bash
# En son Go sürümünü indirin (en son sürüm için https://golang.org/dl/ adresini kontrol edin)
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz

# Önceki kurulumları kaldırın
sudo rm -rf /usr/local/go

# İndirilen arşivi çıkarın
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Go'yu PATH'inize ekleyin
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
```

## 4. Kurulumu Doğrulayın

Go'nun doğru bir şekilde kurulduğunu doğrulamak için aşağıdaki komutu çalıştırın:

```bash
go version
```

Bu komut, kurulu olan Go sürümünü görüntülemelidir.