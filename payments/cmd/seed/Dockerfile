FROM golang:1.26 AS builder

WORKDIR /app

COPY payments/go.mod payments/go.sum ./payments/

WORKDIR /app/payments

RUN go mod download

COPY payments/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o seed ./cmd/seed/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/payments/seed .

ENTRYPOINT ["./seed"]
CMD ["-config", "config.yml"]