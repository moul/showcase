package bots

type Bot interface {
	Name() string
	SetName(string)
	Path() string
	Start() error
	SendMessage(QuestionMessage) (*ReplyMessage, error)
}

type QuestionMessage struct {
	GameID      string      `json:"game-id,omitempty" binding:"required"`
	Action      string      `json:"action,omitempty" binding:"required"`
	Game        string      `json:"game,omitempty" binding:"required"`
	Players     int         `json:"players,omitempty"`
	Board       interface{} `json:"board,omitempty"`
	You         interface{} `json:"you,omitempty"`
	PlayerIndex int         `json:"player-index,omitempty" binding:"required"`
}

type ReplyMessage struct {
	Name        string      `json:"name,omitempty"`
	Play        interface{} `json:"play,omitempty"`
	Error       interface{} `json:"error,omitempty"`
	Comment     interface{} `json:"comment,omitempty"`
	PlayerIndex int         `json:"player-index,omitempty" binding:"required"`
}

// InitTurnBasedBots is an helper that starts and discovers connected bots
func InitTurnBasedBots(bots []Bot, gameName, gameID string) error {
	// start bots
	for _, bot := range bots {
		if err := bot.Start(); err != nil {
			return err
		}
	}

	// send init message to bots
	for idx, bot := range bots {
		reply, err := bot.SendMessage(QuestionMessage{
			GameID:      gameID,
			Action:      "init",
			Game:        gameName,
			Players:     len(bots),
			PlayerIndex: idx,
		})
		if err != nil {
			return err
		}

		// parse reply
		if reply.Name != "" {
			bot.SetName(reply.Name)
		}
	}
	return nil
}
