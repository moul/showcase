SOURCE :=	$(shell find . -name "*.go")
OWN_PACKAGES := $(shell go list ./... | grep -v vendor)


bolosseum: $(SOURCE)
	go build -o ./bolosseum ./cmd/bolosseum/main.go


.PHONY: docker
docker:
	docker build -t moul/bolosseum .


.PHONY: test
test:
	go test -v $(OWN_PACKAGES)


.PHONY: test-coinflip
test-coinflip: bolosseum
	./bolosseum run coinflip file://./test/coinflip.sh file://./test/coinflip.sh


.PHONY: test-russianbullet
test-russianbullet: bolosseum
	./bolosseum run russianbullet file://./test/russianbullet.sh file://./test/russianbullet.sh


.PHONY: test-tictactoe
test-tictactoe: bolosseum test/tictactoe
	./bolosseum run tictactoe file://./test/tictactoe file://./test/tictactoe


test/tictactoe:
	go get github.com/moul/tictactoe/cmd/tictactoe-bolosseum
	ln -s $(GOPATH)/bin/tictactoe-bolosseum $@


.PHONY: test-connect-four
test-connectfour: bolosseum test/connect-four
	./bolosseum run connectfour file://./test/connect-four file://./test/connect-four


test/connect-four:
	go get github.com/moul/connect-four/cmd/connect-four
	ln -s $(GOPATH)/bin/connect-four $@
