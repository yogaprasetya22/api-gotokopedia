# GO Tokopedia (Go-Chi)

## Konfigurasi Environment

Berikut adalah daftar variabel lingkungan yang digunakan dalam proyek ini:

### Server Configuration
- `ADDR`: Alamat dan port di mana server akan berjalan.
  - **Contoh:** `:8080`

### Database Configuration
- `DB_ADDR`: URL koneksi ke database PostgreSQL.
  - **Contoh:** `postgresql://username:password@localhost/database?sslmode=disable`

### Application Environment
- `ENV`: Mode lingkungan aplikasi (misalnya, `development`, `production`).
  - **Contoh:** `development`

### Redis Configuration
- `REDIS_ENABLED`: Mengaktifkan atau menonaktifkan penggunaan Redis.
  - **Contoh:** `true`

### Email Configuration
- `FROM_EMAIL`: Alamat email pengirim.
  - **Contoh:** `your-email@example.com`
- `MAIL_HOST`: Host server SMTP untuk pengiriman email.
  - **Contoh:** `email-smtp.ap-southeast-1.amazonaws.com`
- `MAIL_PORT`: Port server SMTP.
  - **Contoh:** `587`
- `MAIL_USERNAME`: Nama pengguna untuk autentikasi SMTP.
  - **Contoh:** `SMTP_USERNAME`
- `MAIL_PASSWORD`: Kata sandi untuk autentikasi SMTP.
  - **Contoh:** `SMTP_PASSWORD`
- `MAIL_SENDER`: Alamat email pengirim.
  - **Contoh:** `your-email@example.com`

## Cara Mengatur Environment

1. Buat file `.env` di root proyek Anda.
2. Salin variabel di bawah ini ke dalam file `.env` Anda:
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



Pastikan mengganti placeholder seperti `username`, `password`, `your-email@example.com`, dan lainnya dengan nilai sebenarnya sebelum menggunakan konfigurasi tersebut.

