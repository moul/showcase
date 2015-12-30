GOENV ?=	GO15VENDOREXPERIMENT=1
GODEP ?=	$(GOENV) godep
GO ?=		$(GOENV) go
SOURCES :=	$(shell find . -name "*.go")
PORT ?=		8080

all: build

.PHONY: build
build: moul-showcase

.PHONY: test
test:
	$(GODEP) restore
	$(GO) get -t .
	$(GO) test -v .

.PHONY: godep-save
godep-save:
	$(GODEP) save $(shell go list ./... | grep -v /vendor/)

.PHONY: cover
	rm -f profile.out
	$(GO) test -covermode=count -coverpkg=. -coverprofile=profile.out

.PHONY: convey
convey:
	$(GO) get github.com/smartysteets/goconvey
	goconvey -cover -port=9032 -workDir="$(shell realpath .)" -depth=-1

.PHONY: clean
clean:
	rm -rf moul-showcase

moul-showcase: $(SOURCES)
	$(GO) build -o $@ ./cmd/$@

.PHONY: goapp_serve
goapp_serve:
	goapp serve ./appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	goapp deploy -application moul-showcase ./appspot/app.yaml


.PHONY: gin
gin:
	$(GO) get ./...
	$(GO) get github.com/codegangsta/gin
	cd ./cmd/moul-showcase; $(GOENV) gin --immediate --port=$(PORT) server
