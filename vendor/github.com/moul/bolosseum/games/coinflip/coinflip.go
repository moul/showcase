package coinflip

import (
	"fmt"
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type CoinflipGame struct {
	games.BotsBasedGame

	board []string
}

func NewGame() (*CoinflipGame, error) {
	game := CoinflipGame{
		board: make([]string, 2),
	}
	game.Bots = make([]bots.Bot, 0)
	return &game, nil
}

func (g *CoinflipGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *CoinflipGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	for idx, bot := range g.Bots {
		question := bots.QuestionMessage{
			GameID:      gameID,
			Game:        g.Name(),
			Action:      "play-turn",
			Board:       g.board,
			PlayerIndex: idx,
		}
		steps <- games.GameStep{QuestionMessage: &question}
		reply, err := bot.SendMessage(question)
		if err != nil {
			return err
		}
		reply.PlayerIndex = idx
		steps <- games.GameStep{ReplyMessage: reply}

		// parse reply
		if reply.Play != "ship" && reply.Play != "head" {
			err := fmt.Errorf("player %d sent an invalid response: %q", idx, reply.Play)
			steps <- games.GameStep{Error: err}
			return err
		}
		g.board[idx] = reply.Play.(string)
	}

	// check if players chose the same value
	if g.board[0] == g.board[1] {
		err := fmt.Errorf("the second player played the same value")
		steps <- games.GameStep{Error: err}
		return err
	}

	// trigger
	expectedValue := []string{"ship", "head"}[rand.Intn(2)]
	steps <- games.GameStep{Message: fmt.Sprintf("Expected value: %q", expectedValue)}

	for idx, value := range g.board {
		if value == expectedValue {
			steps <- games.GameStep{Winner: g.Bots[idx]}
			//steps <- games.GameStep{Message: fmt.Sprintf("Player %d (%s) won", idx, g.Bots[idx].Name())}
			return nil
		}
	}

	return nil
}

func (g *CoinflipGame) Name() string {
	return "coinflip"
}

func (g *CoinflipGame) GetAsciiOutput() []byte {
	return nil
}
