#DB_URL=mysql://mcputro:welcome1@tcp(localhost:1123)/e_commerce?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true
DB_URL=mysql://mcputro:welcome1@tcp(localhost:1123)/e_commerce

PATH_MIGRATE_FILE=/internal/infrastructure/database/migrations
# Default target
all: menu

# The menu target
menu:
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo " migration-create name={name}  Create migration"
	@echo " migration-up                  Up migrations"
	@echo " migration-down                Down last migration"
	@echo " go-run                        Run Project"
#	@echo " docker-up                     Run with docker"
#	@echo " docker-down                   Stop docker"


migration-create:
	@migrate create -ext sql -dir .$(PATH_MIGRATE_FILE) -seq $(name)

migration-up:
	@migrate -database "$(DB_URL)" -path ./internal/infrastructure/database/migrations up

migration-down:
	@migrate -database "$(DB_URL)" -path ./internal/infrastructure/database/migrations down 1

go-run:
	@go run cmd/main.go