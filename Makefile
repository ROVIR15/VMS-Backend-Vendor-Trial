ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.DEFAULT_GOAL := help

init-install-tools: ## install tools
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest && \
	go get -u -t github.com/vattle/sqlboiler && \
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrate-create: ## database migration: add new migration file need to pass name="migration name" e.g. make migrateadd name="create_users_table"
	goose create $(name) sql

migrate-status: ## database migration: status
	goose status

migrate-up: ## migrate database up to the most recent version available
	goose validate -v
	goose up -v

migrate-down: ## migrate database down roll back the version by 1
	goose validate -v
	goose down -v

test: ## go test
	go test -v ./...
	
test-coverage: ## coverage test
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
	open coverage.html

renew-migration: ## renew migration file. use param="migration name pattern"
	@find migrations -type f -name "*$(param).sql" | sort > result.txt
	DOWN_CONTENT=""; \
	UP_CONTENT=""; \
	while IFS= read -r file; do \
		if [ $${file: -9} == ".down.sql" ]; then \
			DOWN_CONTENT=$$(cat "$$file"); \
			rm "$$file"; \
			continue; \
		else \
			UP_CONTENT=$$(cat "$$file"); \
			rm "$$file"; \
		fi; \
		new_file=$$(echo "$$file" | cut -d'_' -f2- | sed 's/\.up\.sql$$//'); \
		printf "renewing $$new_file\n"; \
		if [ -z "$$(migrate status -database "postgresql://local:local@localhost:5434/hotel?sslmode=disable" -path migrations | grep "$$new_file")" ]; then \
			$(MAKE) migrateadd name="$$new_file"; \
			get_down_file=$$(find migrations -type f -name "*$$new_file.down.sql" | sort); \
			get_up_file=$$(find migrations -type f -name "*$$new_file.up.sql" | sort); \
			if [ ! -z "$$get_down_file" ]; then \
				echo "$$DOWN_CONTENT" > "$$get_down_file"; \
			fi; \
			if [ ! -z "$$get_up_file" ]; then \
				echo "$$UP_CONTENT" > "$$get_up_file"; \
			fi; \
		fi; \
		sleep 1; \
	done < result.txt
	@rm result.txt

.PHONY: install-tools migrateup migratedown compose-exec seed

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
