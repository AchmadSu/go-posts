# go-posts

Repository ini dibuat untuk keperluan technical test di PT Sharing Vision Indonesia.

## 📌 Requirements
- Go version **1.23.2** atau lebih baru
- MySQL dengan database bernama **svi**
  - (Pada contoh ini menggunakan XAMPP sebagai server database)

## ⚙️ Setup Project

1. Install dependencies
   ```bash
   go mod tidy


2. Jalankan migration untuk membuat tabel database
   ```bash
   go run migration/main.go

3. Build project
   ```bash
   go build


4. Jalankan aplikasi
   ```bash
   go run main.go

✅ API Endpoint

Setelah server berjalan, API dapat diakses melalui:

http://localhost:8080

📄 Dokumentasi API (Postman)

Silakan akses dokumentasi Postman pada tautan berikut:

👉 https://www.postman.com/warped-shuttle-585736/workspace/svi-golang/collection/16178191-59a350ac-cfe3-4da8-9513-167955f90c66?action=share&creator=16178191

❓ Bantuan / Kontak

Jika terdapat kendala atau pertanyaan lebih lanjut, silakan hubungi:

📧 Email: ecepentis@gmail.com

📱 WhatsApp: 0896-5842-0438

Terima kasih!

Regards,
Ecep Achmad Sutisna
