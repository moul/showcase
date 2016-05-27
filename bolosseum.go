package moulshowcase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/tictactoe/pkg/tictactoebot"
)

func init() {
	RegisterAction("bolosseum-tictactoe", BolosseumTictactoeAction)
}

func BolosseumTictactoeAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Message string `schema:"message"`
	}
	// FIXME: handle --help

	// Parse query
	m, err := url.ParseQuery(qs)
	if err != nil {
		return nil, err
	}
	if len(m) > 0 {
		// FIXME: if in web mode -> redirect to have options demo in the URL
		decoder := schema.NewDecoder()
		if err := decoder.Decode(&opts, m); err != nil {
			return nil, err
		}
	}

	var question bots.QuestionMessage
	if err := json.Unmarshal([]byte(opts.Message), &question); err != nil {
		return nil, err
	}

	// FIXME: validate input

	fmt.Println(question)
	bot := tictactoebot.NewTictactoeBot()
	reply := &bots.ReplyMessage{}
	switch question.Action {
	case "init":
		reply = bot.Init(question)
	case "play-turn":
		reply = bot.PlayTurn(question)
	default:
		// FIXME: reply message error
		return nil, fmt.Errorf("Unknown action: %q", question.Action)
	}

	return JsonResponse(reply), nil
}
