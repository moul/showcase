GOENV ?=	GO15VENDOREXPERIMENT=1
GO ?=		$(GOENV) go
SOURCES :=	$(shell find . -name "*.go")
PORT ?=		8080

all: build

run: up

up:
	docker-compose up -d --force-recreate --remove-orphans

.PHONY: build
build: moul-showcase

.PHONY: test
test:
	$(GO) get -t .
	$(GO) test -v .

.PHONY: cover
cover:
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


heroku_deploy:
	git remote add heroku https://git.heroku.com/moul-showcase.git || true
	git push heroku master


dokku_deploy:
	git remote add dokku dokku@dokku.m.42.am:moul-showcase || true
	git push dokku master


.PHONY: gin
gin:
	$(GO) get ./...
	$(GO) get github.com/codegangsta/gin
	cd ./cmd/moul-showcase; $(GOENV) gin --immediate --port=$(PORT) server
