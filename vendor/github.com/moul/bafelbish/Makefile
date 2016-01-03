all: build


.PHONY: build
build:
	go get ./...
	go build -o bafelbish ./cmd/bafelbish


.PHONY: test
test:
	go get -t ./...
	go test -v .


.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	go get -t ./...
	goconvey -cover -port=9045 -workDir="$(realpath .)" -depth=0
