FROM golang:1.22-alpine

WORKDIR /app

# Install dependencies
RUN apk add --no-cache bash git

# Copy hanya file go.mod dan go.sum untuk cache dependency
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3000

# Bind mount akan digunakan untuk source code, jadi tidak perlu COPY semua kode
CMD ["sh", "-c", "go run cmd/web/main.go"]
