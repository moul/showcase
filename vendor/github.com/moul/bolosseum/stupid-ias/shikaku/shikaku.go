package stupidshikaku

import (
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

var stupidResponse = `+---+---+---+---+---+---+---+---+---+---+
|   |       |                           |
+   +   +   +   +   +   +   +   +   +   +
|   |       |                           |
+---+   +   +---+---+---+---+---+---+---+
|   |       |                           |
+   +   +   +---+---+---+---+---+---+---+
|   |       |                           |
+   +   +   +   +   +   +   +   +   +   +
|   |       |                           |
+---+---+---+---+---+---+---+---+---+---+
|           |                           |
+   +   +   +   +   +   +   +   +   +   +
|           |                           |
+---+---+---+---+---+---+---+---+---+---+
|                           |           |
+   +   +   +   +   +   +   +   +   +   +
|                           |           |
+   +   +   +   +   +   +   +   +   +   +
|                           |           |
+---+---+---+---+---+---+---+---+---+---+`

type ShikakuStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*ShikakuStupidIA, error) {
	return &ShikakuStupidIA{}, nil
}

func (ia *ShikakuStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "Shikaku StupidIA",
	}
}

func (ia *ShikakuStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{Play: stupidResponse}
}
