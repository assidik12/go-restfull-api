### Back-end Apps online store with GOLANG

hai teman-teman, nama saya <a href="https://www.instagram.com/invites/contact/?i=19oc9ovkmoscw&utm_content=ni3uyco">sidik</a> seorang mahasiswa jurusan teknologi informasi di universitas bina sarana informatika. di sini saya ingin berbagi tentang pembuatan aplikasi toko online sederhana yang didevelop dengan bahasa pemograman GOLANG yang dibuat oleh GOOGLE.

aplikasi ini bertujuan untuk mengatur alur antara seller dan customer dalam kegiatan jual beli secara online.

## Architecture Diagram

![golang clean architecture](architecture.png)

## Features

- account management
- product & category management
- transaction management

## Clone Repository

```bash
# Clone into your workspace
$ git clone https://github.com/assidik12/go-restfull-api.git

# Change directory
$ cd go-restfull-api

# create env file
$ touch .env
```

## Database Migration

- buat database dengan perintah :

```bash
$ migrate create -ext sql -dir db/migrations create_table_xxx
```

- run migrasi database dengan perintah :

```bash
$ migrate -database 'mysql://root:@tcp(localhost:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local' -path db/migrations up
```

## Run Application

- run aplikasi dengan perintah :

```bash
$ go run cmd/web/main.go
```

- run test aplikasi dengan perintah :

```bash
$ go test -v ./test/
```

## View API Documentation

- tambahkan api-key di header pada setiap request api

```bash
 X-API-KEY=RAHASIA DONG BRO
```

- api documentation dengan perintah :

```bash
$ http://localhost:3000/api
```

- main URL :

```bash
$ http://localhost:3000
```

alasan saya membuat aplikasi online-store ini adalah untuk menguji kemampuan saya dalam mengimplementasikan bahasa pemograman GOLANG yang telah saya pelajari pada 1 bulan yang lalu kedalam projek nyata dalam hal ini adalah pembuatan online-store sederhana.

bagi teman-teman yang sudah menyempatkan untuk berkunjung kesini, saya ucapkan banyak terimakasih🤞.

dan tak lupa pula, saya ingin meminta feedback dari teman-teman untuk meluaskan potensi untuk berkembang di masa mendatang.
