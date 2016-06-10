package connectfour

import (
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
)

var Rows = 6
var Cols = 7

func NewConnectfourBot() *ConnectfourBot {
	return &ConnectfourBot{}
}

type ConnectfourBot struct{}

func (b *ConnectfourBot) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	// FIXME: init ttt here
	return &bots.ReplyMessage{
		Name: "moul-connectfour",
	}
}

func (b *ConnectfourBot) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	bot := NewConnectFour()

	board := question.Board
	for y := 0; y < Rows; y++ {
		row := board.([]interface{})[y]
		for x := 0; x < Cols; x++ {
			val := row.([]interface{})[x]
			if val.(string) != "" {
				bot.Board[y][x] = val.(string)
			}
		}
	}

	logrus.Warnf("bot: %v", bot)

	return &bots.ReplyMessage{
		Play: rand.Intn(Cols),
	}
}

type ConnectFour struct {
	Board [][]string
}

func NewConnectFour() ConnectFour {
	cf := ConnectFour{
		Board: make([][]string, Rows),
	}
	for i := 0; i < Rows; i++ {
		cf.Board[i] = make([]string, Cols)
	}
	return cf
}
