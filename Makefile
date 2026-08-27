.PHONY: clean
clean:
	docker compose down -v

.PHONY: infrastructure
infrastructure:
	docker compose up -d --build inventory-postgres
	docker compose up -d --build orders-postgres
	docker compose up -d --build payments-postgres
	docker compose up -d --build loyalty-zookeeper
	docker compose up -d --build loyalty-kafka
	docker compose up -d --build kafka-ui

.PHONY: lint
lint:
	cd migrator && golangci-lint run ./...
	cd topic-creator && golangci-lint run ./...
	cd inventory && golangci-lint run ./...
	cd orders && golangci-lint run ./...

.PHONY: create_topics
create_topics:
	docker compose up -d --build topic-creator

.PHONY: migrate_up
migrate_up:
	docker compose up -d --build inventory-migrator
	docker compose up -d --build orders-migrator

.PHONY: up
up:
	docker compose up -d --build inventory-service
	docker compose up -d --build orders-service

.PHONY: proto
proto:
	protoc \
      --go_out=. \
      --go_opt=module=github.com/identicalaffiliation/loyalty-processor \
      --go-grpc_out=. \
      --go-grpc_opt=module=github.com/identicalaffiliation/loyalty-processor \
      proto/inventory/inventory.proto

.PHONY: seed
seed:
	docker compose up -d --build seed