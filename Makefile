.PHONY: mocks

mocks:
	mockgen -destination mocks/ports.go -package mocks -source ./internal/ports/ports.go
tests: mocks
	go test -v ./...