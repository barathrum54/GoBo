# GoBo

GoBo, Go programlama dilinde yazılmış modüler ve ölçeklenebilir bir backend boilerplate'tir. Fiber framework, GORM ORM ve Zap logging gibi modern araçlar kullanılarak geliştirilmiştir. Bu proje, hızlı API geliştirme ve kolay genişletilebilirlik için tasarlanmıştır.

---

## 🚀 Özellikler

- **Fiber Framework**: Hızlı ve esnek HTTP sunucusu.
- **GORM**: Veritabanı ORM desteği ile kolay modelleme ve migration.
- **Zap Logging**: Performanslı ve yapılandırılabilir loglama.
- **Modüler Mimari**: API genişletilebilir yapıya sahip.
- **Yüksek Kod Kalitesi**: `golangci-lint` ile linter ve statik analiz entegrasyonu.
- **Test Desteği**: `testify` kullanarak birim testler için yapılandırılmış test altyapısı.

---

## 🛠️ Kurulum ve Çalıştırma

### 1. **Depoyu Klonlayın**

```bash
git clone https://github.com/username/gobo.git
cd gobo
```

### 2. **Bağımlılıkları Yükleyin**

```bash
go mod tidy
```

### 3. **.env Dosyasını Oluşturun**

`DATABASE_URL` ortam değişkenini içeren bir `.env` dosyası oluşturun:

```env
DATABASE_URL=postgres://username:password@localhost:5432/dbname
```

### 4. **Veritabanı Migration İşlemi**

Veritabanı tablolarını oluşturmak için proje başlatılırken migration işlemleri otomatik olarak yapılır.

### 5. **Sunucuyu Başlatın**

```bash
go run main.go
```

Sunucu, `http://localhost:3000` adresinde çalışır.

---

## 📂 Proje Yapısı

```
gobo/
├── internal/
│   ├── app/           # Fiber uygulaması ve yapılandırma
│   ├── db/            # Veritabanı bağlantıları
│   ├── logger/        # Zap logger yapılandırması
│   ├── cache/         # Redis bağlantısı ve yardımcı fonksiyonlar
│   ├── models/        # GORM modelleri
│   ├── routes/        # API rotaları
├── .env               # Ortam değişkenleri
├── .golangci-lint.yaml # Linter yapılandırması
├── main.go            # Uygulamanın giriş noktası
├── README.md          # Proje dökümantasyonu
```

---

## 📋 Kullanılan Teknolojiler

- [Go](https://go.dev/) - Programlama dili
- [Fiber](https://gofiber.io/) - HTTP framework
- [GORM](https://gorm.io/) - ORM kütüphanesi
- [Zap](https://github.com/uber-go/zap) - Loglama
- [Redis](https://redis.io/) - Önbellekleme
- [PostgreSQL](https://www.postgresql.org/) - Veritabanı
- [GolangCI-Lint](https://golangci-lint.run/) - Kod analizi ve linter

---

## ✅ Testler

### Testleri Çalıştırmak

Projede bulunan testleri çalıştırmak için aşağıdaki komutu kullanabilirsiniz:

```bash
go test ./... -v
```

Testler, veritabanını sıfırlayıp yeni tablolar oluşturur ve CRUD işlemlerini doğrular.

---

## 🔧 Linter

Projenizde statik kod analizi ve linter kontrolü yapmak için:

```bash
golangci-lint run
```

---

## 🔧 Redis Önbelleği

Proje, Redis ile önbellekleme desteğine sahiptir. Redis bağlantısı `internal/cache` modülünde yönetilir ve API rotalarında kullanılabilir.

### Örnek Kullanım:

Aşağıdaki örnek, bir veriyi Redis önbelleğine kaydetme ve alma işlemini gösterir:

```go
import "gobo/internal/cache"

// Veriyi Redis'e kaydet
cache.Set("key", "value", 60*time.Second)

// Redis'ten veri al
value, err := cache.Get("key")
if err != nil {
    log.Println("Cache miss")
} else {
    log.Printf("Cache hit: %s", value)
}
```

---

## 🔥 Loglama

Proje, **Zap** kullanılarak performanslı ve yapılandırılabilir bir loglama altyapısına sahiptir. Loglama yapılandırması `internal/logger` dizininde bulunur.

### Örnek Kullanım:

```go
import "gobo/internal/logger"

func Example() {
    logger.Log.Info("Example log message", zap.String("key", "value"))
}
```

Loglama yapılandırmasını değiştirmek için `InitLogger` fonksiyonunu kullanabilirsiniz.

---

## 🤝 Katkıda Bulunma

1. Depoyu fork'layın.
2. Kendi dalınızı oluşturun: `git checkout -b my-new-feature`
3. Değişikliklerinizi commit edin: `git commit -m 'Add some feature'`
4. Dalınızı push'layın: `git push origin my-new-feature`
5. Bir PR (Pull Request) oluşturun.

---
