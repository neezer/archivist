SHELL := $(shell which bash)
VERSION := $(shell git tag | tail -n 1 | tail -c +2)
OSARCH := "linux/amd64 darwin/amd64"
ENV = /usr/bin/env
.SHELLFLAGS = -c
.DEFAULT_GOAL := help # Running Make will run the help target

.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.PHONY: all              # All targets are accessible for user

dep: ## Get build dependencies
	go get -v -u github.com/golang/dep/cmd/dep && \
	go get -v -u github.com/mitchellh/gox

build: ## Build the app
	dep ensure && go build

cross-build: ## Build the app for multiple os/arch
	gox -osarch=$(OSARCH) -output="bin/archivist_{{.OS}}_{{.Arch}}" -ldflags="-s -w"
	upx bin/archivist_linux_amd64

debian: ## Build .deb package
	cp bin/archivist_linux_amd64 deb/archivist/usr/local/bin/archivist
	chmod +x deb/archivist/usr/local/bin/archivist
	SIZE=$(shell du -k "./deb/archivist/usr/local/bin/archivist" | cut -f1) perl -pe 's/\{\{(.*?)\}\}/$$ENV{$$1}/' deb/control.template > deb/archivist/DEBIAN/control
	(cd deb && dpkg-deb --build archivist && mv archivist-$(VERSION)_i386.deb)

help: ## Show Help
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
