# application meta info
NAME := boogeyman
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X \'main.Revision=$(REVISION)\'
ENTRYPOINT := main.go

dep:
	dep ensure

# production mode: make [production | pro | p]
production pro p: install build-production test-production

# development mode: make [development | develop | dev | d]
development develop dev d: dep build-development

# buid
build-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o bin/$(NAME) ./$(ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o bin/$(NAME)-darwin ./$(ENTRYPOINT)

# test
test:
	go test ./...

update:
	dep ensure -update