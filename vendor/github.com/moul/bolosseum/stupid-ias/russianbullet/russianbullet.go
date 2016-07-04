package stupidrussianbullet

import (
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

type RussianbulletStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*RussianbulletStupidIA, error) {
	return &RussianbulletStupidIA{}, nil
}

func (ia *RussianbulletStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "Russianbullet StupidIA",
	}
}

func (ia *RussianbulletStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{Play: "click"}
}
