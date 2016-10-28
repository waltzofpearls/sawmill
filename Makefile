.PHONY: all

PKG = $$(glide novendor)

all: build

include docker/docker.mk

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
	@for pkg in $(PKG); do \
		go test -coverprofile c.out.tmp $$pkg; \
		tail -n +2 c.out.tmp >> c.out; \
	done
	go tool cover -html=c.out
