include .env

export POSTGRES_URL = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

read:
	@echo $(POSTGRES_URL)

migrate-create:
	migrate create -dir ./internal/app/migrations -seq -ext sql $(name)

migrate-up:
	migrate -path ./internal/app/migrations -database $(POSTGRES_URL) up

migrate-up-single:
	migrate -path ./internal/app/migrations -database $(POSTGRES_URL) up $(file)

migrate-down-single:
	migrate -path ./internal/app/migrations -database $(POSTGRES_URL) down $(file)

migrate-change-version:
	migrate -path ./internal/app/migrations -database $(POSTGRES_URL) force $(v)
compose-up:
	docker compose up -d

compose-down:
	docker compose down 

compose-clean:
	docker compose down -v


to-database:
	psql -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -U $(POSTGRES_USER) -d $(POSTGRES_DB)



mockgen-user-repository:
	mockgen -source=internal/app/repositories/user-repository.go -destination=internal/app/services/mocks/mock_user_repository.go -package=mocks

mockgen-user-validator:
	mockgen -source=internal/app/validators/user-validator.go -destination=internal/app/services/mocks/mock_user_validator.go -package=mocks

mockgen-order-repository:
	mockgen -source=internal/app/repositories/order-repository.go -destination=internal/app/services/mocks/mock_order_repository.go -package=mocks

mockgen-product-service:
	mockgen -source=internal/app/services/product-service.go -destination=internal/app/services/mocks/mock_product_service.go -package=mocks

mockgen-transactions:
		mockgen -source=./pkg/db/tx.go -destination=internal/app/services/mocks/mock_transactions.go -package=mocks




test-services:
	go test ./internal/app/services/tests

run:
	go run cmd/main.go