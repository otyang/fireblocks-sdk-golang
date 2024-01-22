default: test test-coverage test-integration

.PHONY: tidy
tidy:
	go mod tidy
	

.PHONY: test
test:
	go test -v  ./...

# runs coverage tests and generates the coverage report
test-coverage:
	go test ./... -v -coverpkg=./...

# runs integration tests
test-integration:
	go test ./... -tags=integration ./...