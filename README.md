
# 🛒 Go Tokopedia Clone API (Go + Chi)

Proyek ini merupakan kloning backend Tokopedia menggunakan bahasa pemrograman **Go**, dengan framework **Chi**, database **PostgreSQL**, dan dukungan cache menggunakan **Redis**. Email service disiapkan menggunakan SMTP (Amazon SES).

---

## 📦 Fitur Utama

- RESTful API menggunakan [Go-Chi](https://github.com/go-chi/chi)
- PostgreSQL sebagai database utama
- Redis untuk caching
- Ratelimiter untuk membatasi seberapa sering suatu aksi dapat dilakukan dalam periode waktu tertentu.
- Swagger untuk mempermudah pengembangan dan dokumentasi API RESTful. 
- Auto-reload development server via [Air](https://github.com/cosmtrek/air)

---

## 🚀 Cara Menjalankan Proyek

### 1. Clone Repository

```bash
git clone https://github.com/yogaprasetya22/api-gotokopedia.git
cd api-gotokopedia
````

### 2. Install Dependency

Install module Go:

```bash
go mod tidy
```

### 3. Setup Environment

Buat file `.envrc` di root proyek:

```bash
touch .envrc
```

Isi dengan variabel berikut:

```env
export ADDR=":8080"
export DB_ADDR="postgresql://username:password@localhost/database?sslmode=disable"
export ENV="development"
export REDIS_ENABLED="true"
export FROM_EMAIL="your-email@example.com"
export MAIL_HOST="email-smtp.region.amazonaws.com"
export MAIL_PORT=587
export MAIL_USERNAME="SMTP_USERNAME"
export MAIL_PASSWORD="SMTP_PASSWORD"
export MAIL_SENDER="your-email@example.com"
```

Aktifkan `direnv` untuk memuat environment:

```bash
direnv allow .
```

> **Catatan:** Pastikan `direnv` telah terinstal di sistem Anda. Jika belum, kunjungi: [https://direnv.net/](https://direnv.net/)

---

### 4. Jalankan dengan Docker (Opsional)

Jika ingin menjalankan database dan Redis dengan Docker:

```bash
docker compose up --build
```

---

### 5. Jalankan Server

Jika ingin menggunakan hot reload saat develop:

```bash
air
```

Jika ingin menjalankan secara biasa:

```bash
go run main.go
```

---

## 📚 Dokumentasi Tambahan

* API structure disusun modular untuk memudahkan maintenance dan scale.
* Gunakan tools seperti Postman atau Insomnia untuk testing endpoint.

## 🔗 Tautan Terkait

* 🔗 GitHub: [api-gotokopedia](https://github.com/yogaprasetya22/api-gotokopedia)

---
