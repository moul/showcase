package stupidcoinflip

import (
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

type CoinflipStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*CoinflipStupidIA, error) {
	return &CoinflipStupidIA{}, nil
}

func (ia *CoinflipStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "Coinflip StupidIA",
	}
}

func (ia *CoinflipStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	availableMoves := []string{}
	board := question.Board.([]string)

	for _, move := range []string{"ship", "head"} {
		if board[0] != move && board[1] != move {
			availableMoves = append(availableMoves, move)
		}
	}

	return &bots.ReplyMessage{
		Play: availableMoves[rand.Intn(len(availableMoves))],
	}
}
