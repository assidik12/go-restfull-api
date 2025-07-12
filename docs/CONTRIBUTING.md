### 2. File `CONTRIBUTING.md` (Panduan Kontribusi)

File ini sangat penting untuk menyambut developer lain dan memberi mereka panduan yang jelas tentang cara berpartisipasi dalam proyek Anda.

````markdown
# Panduan Kontribusi untuk Go E-Commerce REST API

Terima kasih banyak atas minat Anda untuk berkontribusi pada proyek ini! Kami sangat menghargai waktu dan usaha yang Anda luangkan. Setiap kontribusi, baik itu laporan _bug_, permintaan fitur baru, perbaikan kode, atau perbaikan dokumentasi, sangat berarti.

Dokumen ini berisi panduan dan alur kerja untuk berkontribusi pada proyek ini.

## 🤔 Bagaimana Saya Bisa Membantu?

- **Melaporkan Bug**: Jika Anda menemukan sesuatu yang tidak berfungsi seperti seharusnya.
- **Mengusulkan Fitur Baru**: Jika Anda punya ide untuk fungsionalitas baru atau perbaikan.
- **Menulis Kode**: Memperbaiki _bug_ atau mengimplementasikan fitur baru.
- **Memperbaiki Dokumentasi**: Memperjelas `README.md`, dokumentasi API, atau komentar kode.

## 🐞 Melaporkan Bug

- **Periksa _Issues_ yang Sudah Ada**: Sebelum membuat laporan baru, mohon periksa daftar [Issues](https://github.com/assidik12/go-restfull-api/issues) untuk memastikan _bug_ tersebut belum pernah dilaporkan.
- **Buat _Issue_ Baru**: Jika belum ada, silakan buat _issue_ baru. Berikan judul yang jelas dan deskripsi yang detail, termasuk langkah-langkah untuk mereproduksi _bug_.

## ✨ Mengusulkan Fitur Baru

Gunakan [GitHub Issues](https://github.com/assidik12/go-restfull-api/issues) untuk mengusulkan fitur baru. Jelaskan ide Anda secara detail, masalah apa yang coba Anda selesaikan, dan bagaimana fitur tersebut akan bekerja.

## 🚀 Alur Kerja Pull Request (PR)

1.  **Fork & Clone**: Buat _fork_ dari repositori ini, lalu _clone_ ke komputer lokal Anda.
2.  **Buat Branch Baru**: Selalu buat _branch_ baru untuk setiap perubahan. Beri nama yang deskriptif (misal: `feat/tambah-fitur-pembayaran`).
    ```bash
    git checkout -b nama-branch-anda
    ```
3.  **Lakukan Perubahan**: Tulis kode Anda dan pastikan semuanya berjalan dengan baik menggunakan `docker-compose up --build`.
4.  **Commit Perubahan**: Buat _commit_ dengan pesan yang jelas mengikuti standar [Conventional Commits](https://www.conventionalcommits.org/).
    ```bash
    git commit -m "feat: Menambahkan endpoint untuk mendapatkan detail transaksi"
    ```
5.  **Push ke Fork Anda**: Kirim perubahan Anda ke repositori _fork_ Anda.
    ```bash
    git push origin nama-branch-anda
    ```
6.  **Buat Pull Request**: Buka halaman repositori Anda di GitHub dan buat _Pull Request_ baru. Jelaskan perubahan yang Anda buat di deskripsi PR.

Setelah PR dibuat, kami akan meninjaunya sesegera mungkin. Sekali lagi, terima kasih!
````
