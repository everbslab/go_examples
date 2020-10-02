VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Go parameters
GOCMD=go
GOFMT=$(GOCMD) fmt
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
DIRBUILD=build

.PHONY: all
all: fmt test build

fmt:
	$(GOFMT) ./...

.PHONY: test
test:
	$(GOCMD) vet ./...
	$(GOTEST) -v ./...
	$(GOTEST) -cover ./...

.PHONY: run
run:
	$(GORUN) ./

.PHONY: build
build:
	$(shell mkdir -p $(DIRBUILD))
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine -v
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(DIRBUILD)/coffee_machine_x64.exe -v

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf ${DIRBUILD}/*

.PHONY: docker-build
docker-build:
	@export DOCKER_CONTENT_TRUST=1 && docker build -f Dockerfile -t coffeemachine:1.0 .

.PHONY: docker-run
docker-run:
	@docker run --rm -it coffeemachine:1.0 /app/coffee_machine