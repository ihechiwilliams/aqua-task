
local:
	go run cmd/server/*.go

# Run go generate locally without docker container
generate-local:
	go run github.com/vektra/mockery/v2@v2.43.0
	go generate ./...

env:
	cp -v .env.example .env

env-test:
	cp -v .env.test .env

build-app:
	docker compose build app-start

# Start the app with seeding
seed:
	docker compose up --build app-seed

# Start the app without seeding
start:
	docker compose up --build app-start

start-notification:
	docker compose up --build notification

clean:
	docker compose down --remove-orphans --volumes


create-migration:
	docker compose run --rm app-start "db/scripts/create_migration.sh $(name)"

migrate:
	docker compose run --rm app-start "db/scripts/migrate.sh"

schema-dump:
	docker compose run --rm app-start "db/scripts/dump.sh > db/schema.sql"

migrate-up:
	/app/db/scripts/migrate.sh
