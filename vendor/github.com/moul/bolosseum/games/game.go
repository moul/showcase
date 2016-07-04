package games

import "github.com/moul/bolosseum/bots"

type GameStep struct {
	Error           error
	QuestionMessage *bots.QuestionMessage
	ReplyMessage    *bots.ReplyMessage
	Message         string
	Winner          bots.Bot
	Loser           bots.Bot
	Draw            bool
}

type Game interface {
	Run(string, chan GameStep) error
	Name() string
	CheckArgs([]string) error
	RegisterBot(bots.Bot)
	GetAsciiOutput() []byte
}
