.DEFAULT_GOAL := run

prepare:
	go generate

run
	go build -tags dev && ./vue-go-starter -debug=true -loglevel=debug

build-windows: prepare
	GOOS=windows GOARCH=amd64 go build

build-linux: prepare
	GOOS=linux GOARCH=amd64 go build

build-mac: prepare
	GOOS=darwin GOARCH=amd64 go build

