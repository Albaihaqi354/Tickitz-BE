# Tickitz Backend API

Tickitz adalah platform pemesanan tiket film online yang memungkinkan pengguna untuk mencari jadwal film, kursi yang tersedia, dan melakukan pemesanan tiket dengan mudah. Repository ini berisi kode sumber untuk layanan backend built-in menggunakan Go.

## Teknologi yang Digunakan

| Komponen | Teknologi |
| --- | --- |
| **Bahasa Pemrograman** | [Go (Golang)](https://go.dev/) |
| **Web Framework** | [Gin Gonic](https://gin-gonic.com/) |
| **Database** | [PostgreSQL](https://www.postgresql.org/) |
| **Caching/Session** | [Redis](https://redis.io/) |
| **Authentication** | [JSON Web Token](https://github.com/golang-jwt/jwt) |
| **Security** | [Argon2/Bcrypt Password Hashing](https://github.com/matthewhartstonge/argon2) |
| **Documentation** | [Swagger](https://swagger.io/) |
| **Dev Ops** | [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/) |
| **Hot Reload** | [Air](https://github.com/cosmtrek/air) |

## Fitur Utama

- **Authentication System**: Registrasi, login, dan reset password.
- **Movie Management**: Daftar film yang sedang tayang, detail film, dan manajemen data film (Admin).
- **Booking System**: Pencarian jadwal bioskop berdasarkan kota/tanggal, manajemen kursi real-time, dan pembuatan pesanan tiket.
- **User Profile**: Kelola informasi akun, riwayat pemesanan, dan upload foto profil.
- **Admin Dashboard**: Statistik penjualan, manajemen jadwal tayang, dan manajemen pengguna.
- **API Documentation**: Dokumentasi otomatis menggunakan Swagger UI.

## Instruksi Instalasi

### 1. Prasyarat
- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/get-started/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download/)

### 2. Setup Environment Variable
Buat file `.env` di direktori root backend:

```env
DB_USER=superuser
DB_PASS=yourpassword
DB_HOST=localhost
DB_PORT=5433
DB_NAME=tickitz

JWT_SECRET=yoursecretkey
JWT_ISSUER=tickitz

RDS_USER=yourredisuser
RDS_PASS=yourredispassword
RDS_HOST=localhost
RDS_PORT=6380
```

### 3. Instalasi Dependensi
Jalankan perintah berikut untuk mengunduh semua library yang diperlukan:
```bash
go mod tidy
```

### 4. Database Migration & Seeder
Pastikan database PostgreSQL sudah berjalan, lalu gunakan Makefile untuk migrasi:
```bash
# Migrasi table
make migrate-up

# Seed data awal
make seeder-up
```

## Cara Penggunaan

### Mode Pengembangan (Production/Direct)
```bash
go run cmd/main.go
```

### Mode Pengembangan dengan Hot Reload (Air)
```bash
air
```

### Menggunakan Docker
Jika Anda ingin menjalankan seluruh stack (BE, DB, Redis) menggunakan Docker Compose:
```bash
docker-compose up
```

## Dokumentasi API

Dokumentasi interaktif Swagger dapat diakses melalui browser di:
`http://localhost:8080/swagger/index.html`

## Informasi Tambahan

- **Project Structure**: Mengikuti pola desain `internal/pkg` untuk menjaga modularitas dan keamanan kode.
- **Error Handling**: Implementasi response standar JSON untuk mempermudah integrasi dengan frontend.
- **Middleware**: Termasuk CORS, Logging, dan Auth.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).

## Project Terkait

- **Frontend**: [Tickitz Frontend](https://github.com/Albaihaqi354/tickitz)
