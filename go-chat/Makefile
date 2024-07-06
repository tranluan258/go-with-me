include .env
run:
	go run main.go

up:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose --dir $(MIGRATION_DIR) up

down:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose --dir $(MIGRATION_DIR) down

reset:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose --dir $(MIGRATION_DIR) reset

db-status:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose --dir $(MIGRATION_DIR) status
