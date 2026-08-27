FROM golang:1.26 AS builder

WORKDIR /app

COPY gen/ ./gen/
COPY orders/go.mod orders/go.sum ./orders/

WORKDIR /app/orders

RUN go mod download

COPY orders/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o orders-service ./cmd/orders/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/orders/orders-service .

ENTRYPOINT ["./orders-service"]
CMD ["-config", "config.yml"]