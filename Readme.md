# Go E-Commerce REST API

![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)
![Docker](https://img.shields.io/badge/Docker-20.10+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

Ini adalah proyek RESTful API untuk aplikasi E-Commerce sederhana, dibangun menggunakan Go (Golang) dengan mengikuti prinsip **Clean Architecture**. Proyek ini bertujuan untuk menyediakan _backend_ yang tangguh dan terstruktur untuk mengelola akun, produk, dan transaksi.

Proyek ini sepenuhnya dikontainerisasi menggunakan **Docker** dan **Docker Compose**, memungkinkan proses instalasi dan deployment yang konsisten di berbagai lingkungan hanya dengan beberapa perintah.

## 🏛️ Arsitektur

Aplikasi ini dirancang dengan pendekatan _Clean Architecture_ untuk memastikan pemisahan yang jelas antara lapisan-lapisan aplikasi (_domain_, _use case_, _interface_), sehingga lebih mudah untuk dipelihara, diuji, dan dikembangkan lebih lanjut.

![golang clean architecture](./docs/architecture.png)

## ✨ Fitur Utama

- ✅ **Manajemen Akun**: Registrasi, login, dan pengelolaan data pengguna.
- ✅ **Manajemen Produk & Kategori**: Operasi CRUD (Create, Read, Update, Delete) untuk produk dan kategori.
- ✅ **Manajemen Transaksi**: Simulasi alur proses transaksi jual beli.
- ✅ **Middleware & Otentikasi**: Menggunakan API Key untuk mengamankan _endpoint_.
- ✅ **Database Migration**: Skema database dikelola menggunakan `golang-migrate`.

## 🛠️ Tumpukan Teknologi (Tech Stack)

- **Bahasa**: Go (Golang)
- **Database**: MySQL 8.0
- **Kontainerisasi**: Docker & Docker Compose
- **Router**: `julienschmidt/httprouter`
- **Migrasi DB**: `golang-migrate`
- **Validasi**: `go-playground/validator`
- **Dependency Injection**: `google/wire`

## 🚀 Memulai (Getting Started)

Proyek ini dirancang untuk berjalan di dalam lingkungan Docker. Pastikan Anda telah menginstal prasyarat yang dibutuhkan.

### Prasyarat

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Langkah-langkah Instalasi

1.  **Clone Repository**

    ```bash
    git clone [https://github.com/assidik12/go-restfull-api.git](https://github.com/assidik12/go-restfull-api.git)
    cd go-restfull-api
    ```

2.  **Buat File Environment (`.env`)**
    Buat file `.env` di direktori utama proyek. Anda bisa menyalin dari contoh di bawah ini. File ini berisi semua konfigurasi penting.

    ```bash
    touch .env
    ```

    Isi file `.env` dengan konfigurasi berikut (sesuaikan jika perlu):

    ```env
    # Konfigurasi Aplikasi
    APP_PORT=3000

    # Konfigurasi Database (digunakan oleh Docker Compose & aplikasi)
    DB_HOST=db
    DB_PORT=3306
    MYSQL_USER=user         # Ganti dengan username Anda
    MYSQL_PASSWORD=secret   # Ganti dengan password yang kuat
    MYSQL_DATABASE=go_rest_api   # Ganti dengan nama database Anda
    MYSQL_ROOT_PASSWORD=secret   # Gunakan password yang sama atau berbeda

    # URL Lengkap untuk koneksi (digunakan oleh golang-migrate & aplikasi Go)
    DB_URL=mysql://user:secret@tcp(db:3306)/go_rest_api?multiStatements=true
    ```

    **PENTING**: Pastikan untuk menambahkan `.env` ke dalam file `.gitignore` Anda!

3.  **Jalankan Aplikasi dengan Docker Compose**
    Perintah ini akan membangun _image_ Docker, membuat kontainer untuk aplikasi Go dan MySQL, menjalankan migrasi database, dan memulai server aplikasi Anda secara otomatis.
    ```bash
    docker-compose up --build
    ```
    Setelah proses selesai, API Anda akan berjalan dan siap diakses.

## usage Penggunaan

- **URL Utama**: `http://localhost:3001`
- **Dokumentasi API**: `http://localhost:3001/api/docs`

## 🧪 Menjalankan Pengujian

Untuk menjalankan _unit test_ yang ada, gunakan perintah Go standar:

```bash
go test -v ./...
```

## 🎉 Kontribusi

Jika Anda ingin berkontribusi, silakan ikuti [petunjuk kontribusi](./docs/CONTRIBUTING.md).

## 🌟 Terima Kasih

Terima kasih telah berkunjung ke Go E-Commerce REST API. Kami sangat menghargai masukan dan umpan balik Anda. Silakan [hubungi kami](https://www.linkedin.com/in/ahmad-sofi-sidik/) untuk memberikan saran atau pertanyaan.
