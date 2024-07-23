.PHONY: generate
generate:
	@echo "Generating code..."
	@go run ./gen/ > ./attack.go
	@go fmt ./attack.go
