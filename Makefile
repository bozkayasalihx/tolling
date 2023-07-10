obu: 
		@go build -o bin/obu obu/main.go
		@bin/obu

recv: 
	@go build -o bin/recv recv/main.go
	@bin/recv

.PHONY: obu recv

