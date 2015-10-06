build: all

cur-dir := $(shell basename `pwd`)

all:
	go build -ldflags "-X main.VERSION=`date +\"%Y%m%d%H%M\"`@`git rev-parse --verify --short HEAD`" main.go
	mv main ${cur-dir}.bin

clean:
	rm -fr *.bin

check:
	go vet ./... && golint ./...

test:
	go test -cover -parallel 2 -v ./...

run:
	go run main.go