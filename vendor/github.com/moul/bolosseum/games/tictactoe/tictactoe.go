package tictactoe

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

var pieces = []string{"X", "O"}

type TictactoeGame struct {
	games.BotsBasedGame

	board map[string]string `json:"board,omitempty"`
}

func NewGame() (*TictactoeGame, error) {
	game := TictactoeGame{}
	game.Bots = make([]bots.Bot, 0)
	game.board = make(map[string]string, 9)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			game.board[fmt.Sprintf("%d-%d", y, x)] = ""
		}
	}

	return &game, nil
}

func (g *TictactoeGame) GetAsciiOutput() []byte {
	str := ""
	str += "+---+---+---+\n"
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			str += fmt.Sprintf("+ %1s ", g.board[fmt.Sprintf("%d-%d", y, x)])
		}
		str += "+\n"
		str += "+---+---+---+\n"
	}
	return []byte(str)
}

func (g *TictactoeGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *TictactoeGame) checkBoard() (bots.Bot, error) {
	// check if the board is invalid
	if len(g.board) != 9 {
		return nil, fmt.Errorf("Invalid board: %d cases", len(g.board))
	}

	// check if there is a winner
	for idx, piece := range pieces {

		// check for horizontal match
		for y := 0; y < 3; y++ {
			valid := true
			for x := 0; x < 3; x++ {
				if g.board[fmt.Sprintf("%d-%d", y, x)] != piece {
					valid = false
					break
				}
			}
			if valid {
				return g.Bots[idx], nil
			}
		}

		// check for vertical match
		for x := 0; x < 3; x++ {
			valid := true
			for y := 0; y < 3; y++ {
				if g.board[fmt.Sprintf("%d-%d", y, x)] != piece {
					valid = false
					break
				}
			}
			if valid {
				return g.Bots[idx], nil
			}
		}

		// diagonals
		if g.board["0-0"] == piece && g.board["1-1"] == piece && g.board["2-2"] == piece {
			return g.Bots[idx], nil
		}
		if g.board["0-2"] == piece && g.board["1-1"] == piece && g.board["2-0"] == piece {
			return g.Bots[idx], nil
		}
	}

	return nil, nil
}

func (g *TictactoeGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	for turn := 0; turn < 9; turn++ {
		idx := turn % 2
		bot := g.Bots[idx]
		piece := pieces[idx]

		copyBoard := make(map[string]string)
		for k, v := range g.board {
			copyBoard[k] = v
		}
		question := bots.QuestionMessage{
			GameID:      gameID,
			Game:        g.Name(),
			Action:      "play-turn",
			Board:       copyBoard,
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
		g.board[reply.Play.(string)] = piece

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

func (g *TictactoeGame) Name() string {
	return "tictactoe"
}
