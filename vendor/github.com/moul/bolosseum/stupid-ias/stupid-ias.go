package stupidias

import "github.com/moul/bolosseum/bots"

type StupidIA interface {
	Init(bots.QuestionMessage) *bots.ReplyMessage
	PlayTurn(bots.QuestionMessage) *bots.ReplyMessage
}
