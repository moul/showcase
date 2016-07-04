package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moul/tictactoe"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ttt := tictactoe.NewTicTacToe()
		ttt.Set(0, 0, c.Query("0-0"))
		ttt.Set(0, 1, c.Query("0-1"))
		ttt.Set(0, 2, c.Query("0-2"))
		ttt.Set(1, 0, c.Query("1-0"))
		ttt.Set(1, 1, c.Query("1-1"))
		ttt.Set(1, 2, c.Query("1-2"))
		ttt.Set(2, 0, c.Query("2-0"))
		ttt.Set(2, 1, c.Query("2-1"))
		ttt.Set(2, 2, c.Query("2-2"))
		ttt.SetPlayer(c.Query("you"))

		fmt.Println(ttt.ShowMap())
		fmt.Printf("Winner: %q\n", ttt.Winner())

		if c.Query("show-map") == "1" {
			c.String(200, ttt.ShowMap())
			return
		}

		next, err := ttt.Next()
		if err != nil {
			c.String(404, fmt.Sprintf("Error: %v", err))
		} else {
			c.String(200, fmt.Sprintf("%d-%d", next.Y, next.X))
		}
	})
	r.Run(":8080")
}
