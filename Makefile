build:
	GOPATH="${PWD}" go build .
run:
	./newegg-watcher
build-run:
	GOPATH="${PWD}" go build . && ./newegg-watcher
