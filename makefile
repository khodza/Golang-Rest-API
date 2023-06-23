include .env

export POSTGRES_URL = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

read:
	@echo $(POSTGRES_URL)

migrate-create:
	migrate create -dir ./internal/db/migrations -seq -ext sql $(name)

migrate-up:
	migrate -path ./internal/db/migrations -database $(POSTGRES_URL) up

migrate-up-single:
	migrate -path ./internal/db/migrations -database $(POSTGRES_URL) up $(file)

migrate-change-version:
	migrate -path ./internal/db/migrations -database $(POSTGRES_URL) force $(v)
compose-up:
	docker compose up -d

compose-down:
	docker compose down 

compose-clean:
	docker compose down -v


to-database:
	psql -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -U $(POSTGRES_USER) -d $(POSTGRES_DB)


run:
	go run cmd/main.go