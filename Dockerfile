FROM golang:1.19 AS builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o my-go-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/my-go-app .

EXPOSE 8080

CMD ["./my-go-app"]
