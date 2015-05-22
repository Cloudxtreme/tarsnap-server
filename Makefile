#!/usr/bin/make -f

SHELL=/bin/bash
bin=bin
name=tarsnapsrv

all: build

build:
	go build -v -o bin/$(name)

clean:
	go clean -x

remove:
	go clean -i