package connectfour

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

var pieces = []string{"X", "O"}
var Rows = 6
var Cols = 7

type ConnectfourGame struct {
	games.BotsBasedGame

	board [][]string `json:"board,omitempty"`
}

func NewGame() (*ConnectfourGame, error) {
	game := ConnectfourGame{}
	game.Bots = make([]bots.Bot, 0)
	game.board = make([][]string, Rows)
	for i := 0; i < Rows; i++ {
		game.board[i] = make([]string, Cols)
	}
	return &game, nil
}

func (g *ConnectfourGame) GetAsciiOutput() []byte {
	sep := ""
	for x := 0; x < Cols; x++ {
		sep += "+---"
	}
	sep += "+\n"

	str := sep
	for y := Rows - 1; y >= 0; y-- {
		for x := 0; x < Cols; x++ {
			str += fmt.Sprintf("+ %1s ", g.board[y][x])
		}
		str += "+\n"
		str += sep
	}
	return []byte(str)
}
func (g *ConnectfourGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *ConnectfourGame) checkBoard() (bots.Bot, error) {
	for idx, piece := range pieces {
		// horizontal
		for y := 0; y < Rows; y++ {
			continuous := 0
			for x := 0; x < Cols; x++ {
				if g.board[y][x] == piece {
					continuous++
					if continuous == 4 {
						return g.Bots[idx], nil
					}
				} else {
					continuous = 0
				}
			}
		}

		// vertical
		for x := 0; x < Cols; x++ {
			continuous := 0
			for y := 0; y < Rows; y++ {
				if g.board[y][x] == piece {
					continuous++
					if continuous == 4 {
						return g.Bots[idx], nil
					}
				} else {
					continuous = 0
				}
			}
		}

		// diagonals
		for y := 0; y < Rows-4; y++ {
			for x := 0; x < Cols-4; x++ {
				continuous := 0
				for i := 0; i < 4; i++ {
					if g.board[y+i][x+i] == piece {
						continuous++
						if continuous == 4 {
							return g.Bots[idx], nil
						}
					} else {
						continuous = 0
					}
				}
			}
		}
	}
	return nil, nil
}

func (g *ConnectfourGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	for turn := 0; ; turn++ {
		if turn == Cols*Rows {
			steps <- games.GameStep{Draw: true}
			return nil
		}
		idx := turn % 2
		bot := g.Bots[idx]
		piece := pieces[idx]

		question := bots.QuestionMessage{
			GameID:      gameID,
			Game:        g.Name(),
			Action:      "play-turn",
			Board:       g.board,
			You:         piece,
			PlayerIndex: idx,
		}
		steps <- games.GameStep{QuestionMessage: &question}
		reply, err := bot.SendMessage(question)
		if err != nil {
			return err
		}
		reply.PlayerIndex = idx
		steps <- games.GameStep{ReplyMessage: reply}

		if reply.Error != nil {
			steps <- games.GameStep{Error: fmt.Errorf("%v", reply.Error)}
			return fmt.Errorf("%v", reply.Error)
		}

		x := int(reply.Play.(float64))
		placed := false
		for y := 0; y < Rows; y++ {
			if g.board[y][x] == "" {
				g.board[y][x] = piece
				placed = true
				break
			}
			if placed {
				break
			}
		}
		if !placed {
			return fmt.Errorf("Invalid location")
		}

		// check board
		winner, err := g.checkBoard()
		if err != nil {
			return err
		}
		if winner != nil {
			steps <- games.GameStep{Winner: winner}
			return nil
		}
	}

	steps <- games.GameStep{Draw: true}
	return nil
}

func (g *ConnectfourGame) Name() string {
	return "connectfour"
}
