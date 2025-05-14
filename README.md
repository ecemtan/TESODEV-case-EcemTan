# Tesodev Product API

Bu proje, Go dili ile yazılmış bir ürün yönetim API'sidir. MongoDB kullanılarak ürün verileri saklanır ve RESTful prensiplere uygun CRUD işlemleri ile ürünler yönetilebilir. Ayrıca, ürünler üzerinde arama yapma imkanı da sunulmaktadır.

## Proje Yapısı

tesodev-product-api/
├── main.go // Uygulama başlangıç noktası
├── config/
│ └── db.go // MongoDB bağlantısı
├── models/
│ └── product.go // Product veri modeli
├── handlers/
│ └── product_handler.go // Tüm CRUD + arama işlemleri
├── routes/
│ └── routes.go // API endpoint tanımları
├── middleware/
│ └── logger.go // Loglama middleware


## Özellikler

- ✅ Ürün oluşturma
- ✅ Ürün listeleme
- ✅ Ürün güncelleme
- ✅ Ürün silme
- ✅ Ürün arama (isim veya açıklamaya göre)
- ✅ Basit loglama sistemi (istekleri loglar)

## Kullanılan Teknolojiler

- **Go** – Backend dili
- **MongoDB** – NoSQL veritabanı
- **Gin** – HTTP web framework
- **mgo / official MongoDB driver** – Mongo bağlantısı için

## Başlangıç

Projeyi çalıştırmak için aşağıdaki adımları izleyin:

### 1. Bağımlılıkları yükleyin

```bash
go mod tidy
Yöntem	Endpoint	Açıklama
GET	/products	Tüm ürünleri listeler
GET	/products/:id	ID ile ürün getirir
POST	/products	Yeni ürün oluşturur
PUT	/products/:id	Ürünü günceller
DELETE	/products/:id	Ürünü siler
GET	/search?q=kelime	Arama yapar
