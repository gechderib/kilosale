FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o kilosale_main_service cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/kilosale_main_service .

EXPOSE 8083

CMD ["./kilosale_main_service"]