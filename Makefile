test:
	go test `go list ./... | grep -v example` -race -coverprofile=coverage.txt -covermode=atomic
