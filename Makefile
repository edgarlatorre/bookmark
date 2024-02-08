.PHONY: all test clean

run:
	go run ./cmd/bookmark

test:
	go test ./test/... -v

get:
	go get ./cmd/bookmark
