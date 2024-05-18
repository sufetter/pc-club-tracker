FROM golang:1.22.3-alpine3.19

WORKDIR /app

COPY . .

CMD ["go", "run", "cmd/main.go", "storage/club.txt"]