package stupidconnectfour

import (
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

var Rows = 6
var Cols = 7

type ConnectfourStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*ConnectfourStupidIA, error) {
	return &ConnectfourStupidIA{}, nil
}

func (ia *ConnectfourStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "ConnectFour StupidIA",
	}
}

func (ia *ConnectfourStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	availableMoves := []int{}
	board := question.Board.([][]string)
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if board[y][x] == "" {
				availableMoves = append(availableMoves, x)
				break
			}
		}
	}

	return &bots.ReplyMessage{
		Play: float64(availableMoves[rand.Intn(len(availableMoves))]),
	}
}
