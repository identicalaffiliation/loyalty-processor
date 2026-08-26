FROM golang:1.26 AS builder

WORKDIR /app

COPY gen/ ./gen/
COPY inventory/go.mod inventory/go.sum ./inventory/

WORKDIR /app/inventory

RUN go mod download

COPY inventory/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o seed ./cmd/seed/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/inventory/seed .

ENTRYPOINT ["./seed"]
CMD ["-config", "config.yml"]