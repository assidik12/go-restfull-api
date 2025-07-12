FROM golang:1.23-alpine AS builder

# Memberikan label untuk maintainer image
LABEL maintainer="Ahmad Sofi Sidik"

# Menginstall paket-paket yang dibutuhkan untuk build
# git: diperlukan oleh 'go mod' untuk beberapa package.
# build-base: berisi compiler C/C++ yang mungkin dibutuhkan oleh beberapa library Go (CGO).
RUN apk add --no-cache git build-base

# Menentukan direktori kerja di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum terlebih dahulu untuk memanfaatkan cache Docker.
# Jika file-file ini tidak berubah, Docker tidak akan men-download ulang dependencies.
COPY go.mod go.sum ./

# Men-download semua dependencies yang terdaftar di go.mod
RUN go mod download

# Menginstall golang-migrate/migrate.
# Ini akan meng-compile dan meletakkan binary 'migrate' di /go/bin/
# Tag 'mysql' diperlukan agar driver database MySQL ikut ter-compile.
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy seluruh sisa source code proyek ke dalam container
COPY . .

# Melakukan build aplikasi Go.
# CGO_ENABLED=0: Membuat static binary tanpa ketergantungan pada library C.
# -a: Memaksa rebuild semua package.
# -installsuffix cgo: Untuk menghindari konflik package dari build sebelumnya.
# -o /app/main: Menentukan nama dan lokasi file output hasil build.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# =================
# Tahap 2: Final Image
# =================
# Menggunakan base image alpine:latest yang sangat ringan untuk production.
FROM alpine:latest

# Menentukan direktori kerja
WORKDIR /app

# Copy binary 'migrate' yang sudah di-build dari tahap 'builder'
COPY --from=builder /go/bin/migrate /usr/local/bin/

# Copy binary aplikasi utama yang sudah di-build dari tahap 'builder'
COPY --from=builder /app/main .

# Copy folder migrasi yang berisi file .sql Anda
COPY ./db/migrations ./db/migrations

# Copy script entrypoint
COPY ./entrypoint.sh .

COPY ./docs ./docs 

# Memberikan izin eksekusi pada script entrypoint
RUN chmod +x ./entrypoint.sh

# Memberi tahu Docker bahwa container akan listen pada port 3000
EXPOSE 3000

# Menjalankan script entrypoint saat container dimulai.
# Ini akan menjalankan migrasi lalu aplikasi Go Anda.
ENTRYPOINT ["/app/entrypoint.sh"]
