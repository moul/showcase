package stupidguessnumber

import (
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

type GuessnumberStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*GuessnumberStupidIA, error) {
	return &GuessnumberStupidIA{}, nil
}

func (ia *GuessnumberStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "Guessnumber StupidIA",
	}
}

func (ia *GuessnumberStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{Play: rand.Intn(101)}
}
