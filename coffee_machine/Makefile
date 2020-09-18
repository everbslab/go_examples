# Go parameters
GOCMD=go
GOFMT=$(GOCMD) fmt
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
DIRBUILD=build

.PHONY: build

all: fmt test build

fmt:
	$(GOFMT) ./...

test:
	$(GOTEST) -v ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine -v
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine_x64.exe -v

clean:
	$(GOCLEAN)
	rm -rf ${DIRBUILD}/*