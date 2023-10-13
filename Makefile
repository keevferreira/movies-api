build:
	@go build -o ../movies-api

run: build
	@../movies-api

test:
	@go test -v ./...