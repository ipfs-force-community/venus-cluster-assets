test-all:
	go test -v ./...

build-all:
	go install ./...

check-all:
	golangci-lint run
