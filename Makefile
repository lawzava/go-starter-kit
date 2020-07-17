# Testing
test:
	go test -race ./...

# Linting
lint: gofumpt gofumports golangci-lint

gofumpt:
	gofumpt -w -s ./..

gofumports:
	gofumports -w ./..

golangci-lint:
	golangci-lint run
