# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
DIRBUILD=build

.PHONY: build

all: test build

test:
	$(GOTEST) -v ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine -v
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine_x64.exe -v

clean:
	$(GOCLEAN)
	rm -rf ${DIRBUILD}/*