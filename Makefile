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

.PHONY: create_topics
create_topics:
	docker compose up -d --build topic-creator

.PHONY: migrate_up
migrate_up:
	docker compose up -d --build inventory-migrator