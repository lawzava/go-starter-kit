build-docker:
	docker build -f deployments/Dockerfile -t goapp:latest .

run-docker:
	docker run -p 8080:8080 --env LISTEN_ADDRESS=':8080' goapp $(APP)

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
