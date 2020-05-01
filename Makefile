test:
	cd ./v2 && go test `go list ./... | grep -v example` -race -coverprofile=../coverage.txt -covermode=atomic

codecov:
	bash <(curl -s https://codecov.io/bash)
