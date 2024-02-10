run:
	@go run cmd/app/main.go

build:
	@go build -o bin/ cmd/app/main.go

watch-asset:
	@npm run watch
