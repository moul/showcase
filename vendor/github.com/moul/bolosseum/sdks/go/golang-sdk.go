package bolosseumbot

import "github.com/moul/bolosseum/bots"

type BolosseumBot interface {
	Init(bots.QuestionMessage) *bots.ReplyMessage
	PlayTurn(bots.QuestionMessage) *bots.ReplyMessage
}
