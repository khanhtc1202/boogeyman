# application meta info
NAME := boogeyman
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X \'main.Revision=$(REVISION)\'
ENTRYPOINT := main.go

dep:
	dep ensure

# production mode: make [production | pro | p]
production pro p: dep gobindata-production build-production test-production

# development mode: make [development | develop | dev | d]
development develop dev d: dep gobindata-development build-development
dev-test: dep build-local test-development

# buid
build-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o bin/$(NAME) ./$(ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o bin/$(NAME)-darwin ./$(ENTRYPOINT)

# test
test-%:
	go test -tags="$* netgo" ./...

update:
	dep ensure -update

# Generate Go file
gobindata-%:
	go-bindata -pkg config -o config/config_tml.go config.$*.tml
