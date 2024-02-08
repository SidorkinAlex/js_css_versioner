.PHONY: build
build:
	go build -v ./cmd/js_css_versioner.go
	$(eval NEW_VER:=$(shell cat version | cut -d '_' -f 2 ))
	mv js_css_versioner js_css_versioner.$(NEW_VER)

test:
	$(eval NEW_VER:=$(shell cat version | cut -d '_' -f 2 ))
	echo $(NEW_VER)


.DEFAULT_GOAL := build
