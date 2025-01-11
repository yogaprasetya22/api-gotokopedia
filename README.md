# GO Tokopedia (Go-Chi)

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

