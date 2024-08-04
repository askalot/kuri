ENTRY_FILE=./cmd/main.go
OUTPUT_DIRECTORY=./bin
OUTPUT_EXECUTABLE=server

.PHONY: clean serve setup

all: setup clean build serve

build:
	go build -o $(OUTPUT_DIRECTORY)/$(OUTPUT_EXECUTABLE) $(ENTRY_FILE)

clean:
	rm -rf $(OUTPUT_DIRECTORY)/$(OUTPUT_EXECUTABLE)

serve:
	go run $(ENTRY_FILE)

setup:
	sh ./scripts/setup.sh
