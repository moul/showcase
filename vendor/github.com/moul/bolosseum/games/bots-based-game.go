package games

import "github.com/moul/bolosseum/bots"

type BotsBasedGame struct {
	Game
	Bots []bots.Bot
}

func (g *BotsBasedGame) RegisterBot(bot bots.Bot) {
	g.Bots = append(g.Bots, bot)
}
