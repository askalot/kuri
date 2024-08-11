ENTRY_FILE=./cmd/main.go
OUTPUT_DIRECTORY=./bin
OUTPUT_EXECUTABLE=server

.PHONY: buildContainer clean debugContainer serve serveContainer setup

all: setup clean build

build:
	go build -o $(OUTPUT_DIRECTORY)/$(OUTPUT_EXECUTABLE) $(ENTRY_FILE)

buildContainer:
	docker build --tag kuri_image .

clean:
	rm -rf $(OUTPUT_DIRECTORY)/$(OUTPUT_EXECUTABLE)

debugContainer:
	docker exec --interactive --tty kuri_container /bin/sh

serve:
	go run $(ENTRY_FILE)

serveContainer:
	docker run --interactive --tty --rm --publish 3000:3000 --name kuri_container kuri_image

setup:
	sh ./scripts/setup.sh

watch:
	air --build.cmd "make build" --build.bin "./bin/server"

usage:
	@echo ""
	@echo "make all		# Run setup, clean, build, serve commands"
	@echo "make build		# Build project"
	@echo "make buildContainer	# Build Docker image"
	@echo "make clean		# Clean build directory"
	@echo "make debugContainer	# Open shell inside running Docker container"
	@echo "make serve		# Start server"
	@echo "make serveContainer	# Start server with Docker"
	@echo "make setup		# Run setup script"
