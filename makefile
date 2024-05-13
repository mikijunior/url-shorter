GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=url_shortener
BINARY_UNIX=$(BINARY_NAME)_unix

all: build

run:
	$(GOCMD) run cmd/main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

# Cross-compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) cmd/main.go

help:
	@echo "Makefile commands:"
	@echo "  all        - Build the project"
	@echo "  run        - Run the application"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  deps       - Install dependencies"
	@echo "  build-linux - Cross-compile for Linux"
	@echo "  install    - Install the application (Optional)"
	@echo "  help       - Display this help message"

.PHONY: all run build clean test deps build-linux install help
