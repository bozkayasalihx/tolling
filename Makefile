obu: 
		@go build -o bin/obu ./obu
		@bin/obu

recv: 
	@go build -o bin/recv ./recv
	@bin/recv

calc: 
	@go build -o bin/calculator ./distance-calculator
	@bin/calculator

aggr: 
	@go build -o bin/aggr ./aggr
	@bin/aggr

.PHONY: obu recv calc aggr

