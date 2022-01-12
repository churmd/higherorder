format:
	go1.18beta1 fmt ./...

test:
	go1.18beta1 test ./...

coverage:
	go1.18beta1 test ./... -coverprofile=coverage.out
	go1.18beta1 tool cover -html=coverage.out

benchmark:
	go1.18beta1 test -bench=. ./...