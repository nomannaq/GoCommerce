build:
	@go build -o bin/webshop-api /Users/noumanqureshi/Desktop/go-projects/webshop-api/cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/webshop-api

