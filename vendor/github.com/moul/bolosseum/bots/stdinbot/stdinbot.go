package stdinbot

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
)

type StdinBot struct {
	name string
}

func (b *StdinBot) Name() string {
	return b.name
}

func (b *StdinBot) SetName(name string) {
	b.name = name
}

func (b *StdinBot) Path() string {
	return "-"
}

func (b *StdinBot) SendMessage(msg bots.QuestionMessage) (*bots.ReplyMessage, error) {
	switch msg.Action {
	case "init":
		return &bots.ReplyMessage{
			Name: b.Name(),
		}, nil
	case "play-turn":

		// execute script
		fmt.Println(msg)
		var userInput string
		if _, err := fmt.Scanln(&userInput); err != nil {
			return nil, err
		}

		// parse json
		var response bots.ReplyMessage
		response.Play = userInput
		return &response, nil
	default:
		return nil, fmt.Errorf("Unknown action %q", msg.Action)
	}
}

func (b *StdinBot) Start() error {
	return nil
}

func NewBot() (*StdinBot, error) {
	return &StdinBot{}, nil
}
