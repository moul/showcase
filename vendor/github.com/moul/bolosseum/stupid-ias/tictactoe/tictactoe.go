package stupidtictactoe

import (
	"fmt"
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

type TictactoeStupidIA struct {
	stupidias.StupidIA
}

func NewIA() (*TictactoeStupidIA, error) {
	return &TictactoeStupidIA{}, nil
}

func (ia *TictactoeStupidIA) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	return &bots.ReplyMessage{
		Name: "Tictactoe StupidIA",
	}
}

func (ia *TictactoeStupidIA) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	availableMoves := []string{}

	board := question.Board.(map[string]string)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			move := fmt.Sprintf("%d-%d", y, x)
			if board[move] == "" {
				availableMoves = append(availableMoves, move)
			}
		}
	}

	return &bots.ReplyMessage{
		Play: availableMoves[rand.Intn(len(availableMoves))],
	}
}
