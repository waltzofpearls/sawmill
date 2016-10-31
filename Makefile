.PHONY: all deps

PKG = $$(glide novendor)
PKG_COVER = $$(go list ./... | grep -v /vendor/)

all: build

include docker/docker.mk

deps:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install

build: *.go
	go build -race -o sawmill ./

test:
	go vet $(PKG)
	go test -race $(PKG)

testv:
	go vet -v $(PKG)
	go test -race -v -cover $(PKG)

clean:
	go clean $(PKG)

cover:
	@echo "mode: count" > c.out
	@for pkg in $(PKG_COVER); do \
		go test -coverprofile c.out.tmp $$pkg; \
		tail -n +2 c.out.tmp >> c.out; \
	done
	go tool cover -html=c.out
