obu: 
		@go build -o bin/obu ./obu
		@bin/obu

recv: 
	@go build -o bin/recv ./recv
	@bin/recv

.PHONY: obu recv

