## *********************
## * Makefile commands *
## *********************
##


.DEFAULT_GOAL := help
SHELL := /bin/bash


.PHONY: help
help:   ## show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)


.PHONY: gen
gen:    ## run code generation
	go generate -x ./...

.PHONY: tools
tools:  ## build tools
	go generate -x -tags tools ./tools/

.PHONY: tests
tests:  ## run tests
	go test --cover ./...
