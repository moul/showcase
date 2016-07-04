package shikaku

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
	"github.com/moul/shikaku"
)

type ShikakuGame struct {
	games.BotsBasedGame

	shikaku *shikaku.ShikakuMap
}

var (
	width  = 10
	height = 10
	blocks = 10
)

func NewGame() (*ShikakuGame, error) {
	game := ShikakuGame{}
	game.Bots = make([]bots.Bot, 0)
	game.shikaku = shikaku.NewShikakuMap(width, height, 0, 0)
	if err := game.shikaku.GenerateBlocks(blocks); err != nil {
		return nil, err
	}
	return &game, nil
}

func (g *ShikakuGame) CheckArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("You need to specify exactly 1 bot")
	}
	return nil
}

func (g *ShikakuGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	bot := g.Bots[0]

	question := bots.QuestionMessage{
		GameID:      gameID,
		Players:     1,
		Game:        g.Name(),
		Action:      "play-turn",
		PlayerIndex: 0,
		Board:       g.shikaku.String(),
	}
	steps <- games.GameStep{QuestionMessage: &question}
	reply, err := bot.SendMessage(question)
	if err != nil {
		return err
	}
	reply.PlayerIndex = 0
	steps <- games.GameStep{ReplyMessage: reply}

	// FIXME: check answer
	if reply.Play == g.shikaku.DrawSolution() {
		steps <- games.GameStep{Winner: bot}
	} else {
		steps <- games.GameStep{Loser: bot}
	}
	return nil
}

func (g *ShikakuGame) Name() string {
	return "shikaku"
}

func (g *ShikakuGame) GetAsciiOutput() []byte {
	return []byte(g.shikaku.DrawMap())
}
