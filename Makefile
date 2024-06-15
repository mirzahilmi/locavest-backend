.PHONY: run
run:
	@go run cmd/api/main.go
.PHONY: build
build:
	@CGO_ENABLED=0 go build -o tmp/app cmd/api/main.go