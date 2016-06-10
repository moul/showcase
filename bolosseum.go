package moulshowcase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/sdks/go"
	"github.com/moul/connect-four"
	"github.com/moul/tictactoe/pkg/tictactoebot"
)

func init() {
	RegisterAction("bolosseum-tictactoe", BolosseumTictactoeAction)
	RegisterAction("bolosseum-connectfour", BolosseumConnectfourAction)
}

func BolosseumConnectfourAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	bot := connectfour.NewConnectfourBot()
	return BolosseumGenericAction(bot, qs, stdin)
}

func BolosseumTictactoeAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	bot := tictactoebot.NewTictactoeBot()
	return BolosseumGenericAction(bot, qs, stdin)
}

func BolosseumGenericAction(bot bolosseumbot.BolosseumBot, qs string, stdin io.Reader) (*ActionResponse, error) {
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

	var inputMessage string

	if opts.Message != "" {
		inputMessage = opts.Message
	} else {
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			line := scanner.Text()
			inputMessage += line
		}
	}

	var question bots.QuestionMessage
	if err := json.Unmarshal([]byte(inputMessage), &question); err != nil {
		return nil, err
	}

	// FIXME: validate input

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
