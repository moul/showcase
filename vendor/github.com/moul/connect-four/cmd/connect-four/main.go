package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/connect-four"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	gin.DisableBindValidation()
}

func main() {
	if len(os.Args) == 1 {
		// web mode
		logrus.Warnf("You ran this program without argument, it will then start a web server")
		logrus.Warnf("usage: ")
		logrus.Warnf("- %s            # web mode", os.Args[0])
		logrus.Warnf("- %s some-json  # cli mode", os.Args[0])

		r := gin.Default()

		r.POST("/", func(c *gin.Context) {
			var question bots.QuestionMessage
			if err := c.BindJSON(&question); err != nil {
				fmt.Println(err)
				c.JSON(404, fmt.Errorf("Invalid POST data: %v", err))
				return
			}

			bot := connectfour.NewConnectfourBot()
			reply := &bots.ReplyMessage{}
			switch question.Action {
			case "init":
				reply = bot.Init(question)
			case "play-turn":
				reply = bot.PlayTurn(question)
			default:
				// FIXME: reply message error
				c.JSON(500, gin.H{"Error": fmt.Errorf("Unknown action: %q", question.Action)})
				return
			}

			c.JSON(200, reply)
		})

		r.GET("/", func(c *gin.Context) {
			if message := c.Query("message"); message != "" {
				bot := connectfour.NewConnectfourBot()

				logrus.Warnf("<< %s", message)
				var question bots.QuestionMessage
				if err := json.Unmarshal([]byte(message), &question); err != nil {
					c.JSON(500, gin.H{"Error": err})
					return
				}

				reply := &bots.ReplyMessage{}
				switch question.Action {
				case "init":
					reply = bot.Init(question)
				case "play-turn":
					reply = bot.PlayTurn(question)
				default:
					// FIXME: reply message error
					c.JSON(500, gin.H{"Error": fmt.Errorf("Unknown action: %q", question.Action)})
					return
				}

				c.JSON(200, reply)
			} else {
				c.String(404, "This server is a bot for bolosseum.")
			}
		})
		r.Run(":8080")
	} else {
		// cli mode
		logrus.Warnf("%s << %v", os.Args[0], os.Args[1])

		var question bots.QuestionMessage
		if err := json.Unmarshal([]byte(os.Args[1]), &question); err != nil {
			logrus.Fatalf("%s XX err: %v", err)
		}

		bot := connectfour.NewConnectfourBot()

		reply := &bots.ReplyMessage{}
		switch question.Action {
		case "init":
			reply = bot.Init(question)
		case "play-turn":
			reply = bot.PlayTurn(question)
		default:
			// FIXME: reply message error
			logrus.Fatalf("Unknown action: %q", question.Action)
		}

		jsonString, err := json.Marshal(reply)
		if err != nil {
			logrus.Fatalf("Failed to marshal json: %v", err)
		}

		fmt.Println(string(jsonString))
	}
}
