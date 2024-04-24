run: build
	@./bin/app
build:
	@mkdir -p bin
	@go build -o bin/app cmd/app/main.go
