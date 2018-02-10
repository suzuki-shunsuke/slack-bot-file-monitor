t:
	go test ./... -covermode=atomic
run:
	bash entrypoint.sh
main: main.go
	GOOS=linux GOARCH=amd64 go build main.go
