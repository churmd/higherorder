format:
	go1.18beta1 fmt ./...

clean_test_cache:
	go1.18beta1 clean -testcache

test: clean_test_cache
	go1.18beta1 test ./...

coverage:
	go1.18beta1 test ./... -coverprofile=coverage.out
	go1.18beta1 tool cover -html=coverage.out

benchmark:
	go1.18beta1 test -bench=. ./...