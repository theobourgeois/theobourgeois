.PHONY: run run-gen db-up clean

DB_USER := $(shell grep "DB_USER" .env | cut -d '=' -f2)
DB_NAME := $(shell grep "DB_NAME" .env | cut -d '=' -f2)
DB_PASSWORD := $(shell grep "DB_PASSWORD" .env | cut -d '=' -f2)
DB_PORT := $(shell grep "DB_PORT" .env | cut -d '=' -f2)
DB_HOST := $(shell grep "DB_HOST" .env | cut -d '=' -f2)

db-up:
	@for file in database/migrations/*.sql ; do \
		mysql -u$(DB_USER) -p$(DB_PASSWORD) -h ${DB_HOST} --port ${DB_PORT} $(DB_NAME) < $$file; \
		echo "Executed $$file"; \
	done

gen:
	npx tailwindcss -i ./styles/input.css -o ./static/output.css
	templ generate

run-gen:
	npx tailwindcss -i ./styles/input.css -o ./static/output.css
	templ generate
	go run .

run: 
	go run .

build:
	go build -o out . 

clean: 
	go tidy
	go clean -cache