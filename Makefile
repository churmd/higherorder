format:
	go fmt ./...

clean_test_cache:
	go clean -testcache

test: clean_test_cache
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

benchmark:
	go test -bench=. ./...

run-example:
	go run examples/example.go