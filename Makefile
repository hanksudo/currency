.DEFAULT_GOAL := start
.PHONY: start

## Start server
start:
	go run . -web