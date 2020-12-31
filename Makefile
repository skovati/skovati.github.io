# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=kripto
DIR=/usr/local/bin

all: test build
build: 
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./src/...
test: 
	$(GOTEST) -v ./src/...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./src/...
	./bin/$(BINARY_NAME)
install:
	install -D ./bin/kripto $(DIR)
uninstall:
	rm -f $(DIR)kripto
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v ./src/...
