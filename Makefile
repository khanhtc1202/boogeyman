# application meta info
NAME := boogeyman
VERSION= 1.2.7
REVISION := $(shell git rev-parse --short HEAD)
BUILDDATE := $(shell date '+%Y/%m/%d %H:%M:%S %Z')
GOVERSION := $(shell go version)
LDFLAGS := -X 'main.revision=$(REVISION)' -X 'main.version=$(VERSION)' -X 'main.buildDate=$(BUILDDATE)' -X 'main.goVersion=$(GOVERSION)'
CLI_ENTRYPOINT := cmd/boogeyman/main.go
REST_ENTRYPOINT := web/boogeyman/app.go
REQUIRETESTPKG := internal

all: dep cli rest

dep:
	dep ensure

cli: gobindata-production build-cli-production test-production

rest: gobindata-production build-rest-production test-production

# development mode: make [development | develop | dev | d]
development develop dev d: dep gobindata-development build-cli-development build-rest-development

# buid
build-cli-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-linux-64 ./$(CLI_ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-darwin-64 ./$(CLI_ENTRYPOINT)

build-rest-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-rest-linux-64 ./$(REST_ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-rest-darwin-64 ./$(REST_ENTRYPOINT)

# test
test-%:
	go test -tags="$* netgo" ./$(REQUIRETESTPKG)...

update:
	dep ensure -update

# Generate Go file
gobindata-%:
	go-bindata -pkg config -o config/config_tml.go config.$*.tml
