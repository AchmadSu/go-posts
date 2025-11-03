# go-posts
ini adalah repository yang dibuat khusus untuk kebutuhan test di PT Sharing Vision Indonesia

# Requirement
Go version go1.23.2 windows/amd64 +
MySQL berisi database dengan nama svi
(Saya menggunakan XAMPP untuk kebutuhan database nya)

# SetUp
Silakan jalankan command:
```cmd
go mod tidy
```
Berfungsi untuk menginstall dependency dari project.

Lalu, jalankan command:
```cmd
go run migration/main.go
```
Untuk memigrasi database

Lalu running command:
```cmd
go build
```

Jika proses building success, maka jalankan command:
```cmd
go run main
```

Selamat anda dapat menggunakan API posts melalui tautan: http://localhost:8080 !
Untuk dokumentasi Postman silakan akses tautan berikut: 
https://www.postman.com/warped-shuttle-585736/workspace/svi-golang/collection/16178191-59a350ac-cfe3-4da8-9513-167955f90c66?action=share&creator=16178191

Jika ada kendala atau pertanyaan lebih lanjut silakan tanyakan kepada saya melalui Email: ecepentis@gmail.com atau What'sApp: 0896 5842 0438
Terima kasih banyak!

Regards,
Ecep Achmad Sutisna
