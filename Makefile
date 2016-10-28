.PHONY: all

all: build

include docker/docker.mk

build: *.go
		go build -race -o sawmill ./
