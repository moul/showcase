package russianbullet

import (
	"fmt"
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type RussianbulletGame struct {
	games.BotsBasedGame
}

func NewGame() (*RussianbulletGame, error) {
	game := RussianbulletGame{}
	game.Bots = make([]bots.Bot, 0)
	return &game, nil
}

func (g *RussianbulletGame) CheckArgs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("You need to specify at least 1 bot")
	}
	return nil
}

func (g *RussianbulletGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	bulletIndex := rand.Intn(6) // 6 slots in the revolver
	for i := 0; i <= bulletIndex; i++ {
		idx := i % len(g.Bots)
		bot := g.Bots[idx]

		question := bots.QuestionMessage{
			GameID:      gameID,
			Players:     len(g.Bots),
			Game:        g.Name(),
			Action:      "play-turn",
			PlayerIndex: idx,
		}
		steps <- games.GameStep{QuestionMessage: &question}
		reply, err := bot.SendMessage(question)
		if err != nil {
			return err
		}
		reply.PlayerIndex = idx
		steps <- games.GameStep{ReplyMessage: reply}

		if reply.Play != "click" {
			err := fmt.Errorf("Invalid bot input: %v", reply.Play)
			steps <- games.GameStep{Error: err}
			return err
		}
		if i == bulletIndex {
			steps <- games.GameStep{Loser: g.Bots[idx]}
			return nil
		}
	}

	steps <- games.GameStep{Draw: true}
	return nil
}

func (g *RussianbulletGame) Name() string {
	return "russianbullet"
}

func (g *RussianbulletGame) GetAsciiOutput() []byte {
	return nil
}
