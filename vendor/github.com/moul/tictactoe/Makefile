SOURCES :=	$(shell find . -name "*.go")


tictactoe-bolosseum: $(SOURCES)
	go build -o ./tictactoe-bolosseum ./cmd/tictactoe-bolosseum


.PHONY: goapp_serve
goapp_serve:
	goapp serve ./cmd/appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	goapp deploy -application tictactoe-as-a-service ./cmd/appspot/app.yaml
