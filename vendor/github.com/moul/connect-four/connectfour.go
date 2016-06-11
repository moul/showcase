package connectfour

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
	"github.com/robfig/go-cache"
)

var Rows = 6
var Cols = 7
var MaxDeepness = 6

var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 30*time.Second)
}

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

	// get movements
	moves := bot.BestMovements()
	if len(moves) == 0 {
		return &bots.ReplyMessage{
			Error: "no available movement",
		}
	}

	// pick one
	picked := moves[rand.Intn(len(moves))]

	logrus.Warnf("Playing %d with score %f, %d best moves", picked.Play, picked.Score, len(moves))
	return &bots.ReplyMessage{
		Play: picked.Play,
	}
}

func (b *ConnectFour) Hash(currentPlayer string) string {
	hash := ""
	hash += fmt.Sprintf("%d", MaxDeepness)
	for y := 0; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if b.Board[y][x] != "" {
				hash += b.Board[y][x]
			} else {
				hash += "."
			}
		}
	}

	hash += currentPlayer
	return hash
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
	for _, piece := range pieces {
		for x := 0; x < Cols; x++ {
			continuous := 0
			for y := 0; y < Rows; y++ {
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

	// diagnoals
	for _, piece := range pieces {
		for x := 0; x < Cols-4; x++ {
			for y := 0; y < Rows-4; y++ {
				continuous := 0
				for i := 0; i < 4; i++ {
					if b.Board[y+i][x+i] == piece {
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
	}

	return ""
}

func (b *ConnectFour) BestMovements() []Movement {
	hash := b.Hash(b.Player)
	if cachedMoves, found := c.Get(hash); found {
		return cachedMoves.([]Movement)
	}

	logrus.Warnf("bot: %v", b)
	moves := b.ScoreMovements(b.Player, 1)
	logrus.Warnf("score-moves: %v", moves)
	b.PrintMap()

	if len(moves) == 0 {
		return moves
	}

	// take the best score
	maxScore := moves[0].Score
	for _, move := range moves {
		if move.Score > maxScore {
			maxScore = move.Score
		}
	}
	bestMoves := []Movement{}
	for _, move := range moves {
		if move.Score == maxScore {
			bestMoves = append(bestMoves, move)
		}
	}

	c.Set(hash, bestMoves, -1)
	return bestMoves
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
	if deepness == 1 {
		logrus.Warnf("score=%q deepness=%d moves=%v winner=%q value=%f", currentPlayer, deepness, moves, b.Winner(), value)
	}

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
