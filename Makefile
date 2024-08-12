.PHONY: buildContainer clean serve serveContainer setup shell watch usage

all: setup clean build

build:
	go build -o ./bin/server ./cmd/main.go

buildContainer:
	docker compose build

clean:
	rm -rf ./bin/server

serve:
	go run ./cmd/main.go -b 0.0.0.0

serveContainer:
	docker compose up

setup:
	sh ./scripts/setup.sh

shell:
	docker compose run --service-ports web /bin/sh

watch:
	air

usage:
	@echo ""
	@echo "make all		# Run setup, clean, build, serve commands"
	@echo "make build		# Build project"
	@echo "make buildContainer	# Build Docker image"
	@echo "make clean		# Clean build directory"
	@echo "make serve		# Start server"
	@echo "make serveContainer	# Start server with Docker"
	@echo "make setup		# Run setup script"
	@echo "make shell		# Open shell inside running Docker container"
