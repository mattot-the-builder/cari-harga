run:
	@go run cmd/app/main.go

build:
	@go build -o dist/ cmd/app/main.go
