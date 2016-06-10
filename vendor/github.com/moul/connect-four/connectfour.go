package connectfour

import (
	"math"
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
)

var Rows = 6
var Cols = 7
var MaxDeepness = 3

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
	bot := NewConnectFour(question.You.(string))

	doneMoves := 0
	board := question.Board
	for y := 0; y < Rows; y++ {
		row := board.([]interface{})[y]
		for x := 0; x < Cols; x++ {
			val := row.([]interface{})[x]
			if val.(string) != "" {
				bot.Board[y][x] = val.(string)
				doneMoves++
			}
		}
	}

	// first move is random
	if doneMoves == 0 {
		play := rand.Intn(Cols)
		logrus.Warnf("the first move is always random, playing %d", play)
		return &bots.ReplyMessage{
			Play: play,
		}
	}

	// Debug
	logrus.Warnf("bot: %v", bot)
	bot.PrintMap()
	moves := bot.ScoreMovements(bot.Player, 1)
	logrus.Warnf("score-moves: %v", moves)
	bot.PrintMap()

	if len(moves) == 0 {
		return &bots.ReplyMessage{
			Error: "no available movement",
		}
	}

	// take the best score
	maxIdx := 0
	maxScore := moves[0].Score
	for idx, move := range moves {
		if move.Score > maxScore {
			maxScore = move.Score
			maxIdx = idx
		}
	}
	logrus.Warnf("Playing %d with score %f", moves[maxIdx].Play, moves[maxIdx].Score)
	return &bots.ReplyMessage{
		Play: moves[maxIdx].Play,
	}
}

type ConnectFour struct {
	Board  [][]string
	Player string
}

type Movement struct {
	Play  int
	Score float64
}

func (b *ConnectFour) PrintMap() {
	for y := 0; y < Rows; y++ {
		line := "|"
		for x := 0; x < Cols; x++ {
			if b.Board[y][x] != "" {
				line += b.Board[y][x] + "|"
			} else {
				line += " |"
			}
		}
		logrus.Warnf(line)
	}
}

func (b *ConnectFour) Winner() string {
	pieces := []string{"X", "O"}

	// horizontal
	for _, piece := range pieces {
		for y := 0; y < Rows; y++ {
			continuous := 0
			for x := 0; x < Cols; x++ {
				if b.Board[y][x] == piece {
					continuous++
					if continuous == 4 {
						return piece
					}
				} else {
					continuous = 0
				}
			}
		}
	}

	//vertical
	// FIXME

	// diagnoals
	// FIXME

	return ""
}

func (b *ConnectFour) ScoreMovements(currentPlayer string, deepness int) []Movement {
	// check if board is already finished
	if b.Winner() != "" {
		return []Movement{}
	}

	// get available moves
	moves := b.AvailableMovements()

	// useless to go too deep
	if deepness > MaxDeepness {
		return moves
	}

	//size := Cols * Rows
	value := math.Pow(float64(MaxDeepness+1), float64(MaxDeepness-deepness))
	logrus.Warnf("score=%q deepness=%d moves=%v winner=%q value=%f", currentPlayer, deepness, moves, b.Winner(), value)

	for idx, move := range moves {
		b.Play(move.Play, currentPlayer)
		switch b.Winner() {
		case b.Player:
			moves[idx].Score = value
		case b.Opponent():
			moves[idx].Score = -value
		default:
			for _, subMove := range b.ScoreMovements(b.NextPlayer(currentPlayer), deepness+1) {
				moves[idx].Score += subMove.Score
			}
		}
		b.Play(move.Play, "")
	}

	return moves
}

func (b *ConnectFour) Opponent() string {
	return b.NextPlayer(b.Player)
}

func (b *ConnectFour) NextPlayer(currentPlayer string) string {
	switch currentPlayer {
	case "X":
		return "O"
	case "O":
		return "X"
	}
	return ""
}

func (b *ConnectFour) AvailableMovements() []Movement {
	movements := []Movement{}
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if b.Board[y][x] == "" {
				movement := Movement{
					Play:  x,
					Score: 0,
				}
				movements = append(movements, movement)
				break
			}
		}
	}
	return movements
}

func (b *ConnectFour) Play(col int, piece string) {
	if piece != "" { // place a piece
		for y := 0; y < Rows; y++ {
			if b.Board[y][col] == "" {
				b.Board[y][col] = piece
				return
			}
		}
	} else { // remove a piece
		for y := Rows - 1; y >= 0; y-- {
			if b.Board[y][col] != "" {
				b.Board[y][col] = ""
				return
			}
		}
	}
}

func NewConnectFour(player string) ConnectFour {
	cf := ConnectFour{
		Board:  make([][]string, Rows),
		Player: player,
	}
	for i := 0; i < Rows; i++ {
		cf.Board[i] = make([]string, Cols)
	}
	return cf
}
