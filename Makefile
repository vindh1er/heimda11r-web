# Go parameters
GOCMD=go
GOBUILD=GO111MODULE=on $(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=GO111MODULE=on $(GOCMD) test
BINARY_NAME=heimda11r-web
VERSION := $(shell git rev-parse HEAD)
COMMIT := $(shell git rev-parse HEAD)

all: test build
build:
	pkger
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -s -w" -v
test:
	pkger
	printf "Linter:\n"
	GO111MODULE=on $(GOCMD) list ./... | xargs -L1 golint | tee golint-report.out
	printf "\n\nTests:\n\n"
	GO111MODULE=on $(GOCMD) test -v --bench --benchmem -coverprofile coverage.txt ./...
	GO111MODULE=on $(GOCMD) vet 2> govet-report.out
	GO111MODULE=on $(GOCMD) tool cover -html=coverage.txt -o cover-report.html
	printf "\nCoverage report available at cover-report.html\n\n"

clean:
	$(GOCLEAN)
	$(GOCMD) fmt ./...
	rm -f $(BINARY_NAME)
	go mod tidy
run-gopher:
	gophor -root $(PWD)/gopher -bind-addr 0.0.0.0 -footer "Heimda11r - Served by Gophor"
