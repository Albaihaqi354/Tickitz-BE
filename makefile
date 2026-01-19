include ./.env

MIGRATION_PATH_MIGRATION=./db/migration
MIGRATION_PATH_SEEDER=./db/seeder
DB_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
DB_URL_SEEDER=$(DB_URL)&x-migrations-table=schema_seeders

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH_MIGRATION) -seq create_$(NAME)_table

migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH_MIGRATION) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH_MIGRATION) down

seeder-create:
	migrate create -ext sql -dir $(MIGRATION_PATH_SEEDER) -seq seed_$(NAME)

seeder-up:
	migrate -database "$(DB_URL_SEEDER)" -path $(MIGRATION_PATH_SEEDER) up

seeder-down:
	migrate -database "$(DB_URL_SEEDER)" -path $(MIGRATION_PATH_SEEDER) down