BINDIR=${CURDIR}/bin
APPNAME=js_css_versioner

.PHONY: build
build:
	go build -o ${BINDIR}/${APPNAME} -v ./cmd/${APPNAME}/main.go
	$(eval NEW_VER:=$(shell cat version | cut -d '_' -f 2 ))
	mv ${BINDIR}/${APPNAME} ${BINDIR}/${APPNAME}.$(NEW_VER)

test:
	$(eval NEW_VER:=$(shell cat version | cut -d '_' -f 2 ))
	echo $(NEW_VER)


.DEFAULT_GOAL := build
