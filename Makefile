.PHONY: start
start:
	go run main.go

.PHONY: release
release: main.go
	GOOS=linux GOARCH=amd64 go build -o main --tags=release main.go &&\
	zip -q function.zip main
