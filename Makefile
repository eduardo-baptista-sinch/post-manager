run.test:
	go test -v -covermode=atomic -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out
