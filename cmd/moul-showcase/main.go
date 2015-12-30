package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/moul/showcase"
)

func main() {
	app := cli.NewApp()
	app.Name = "moul-showcase"
	app.Usage = "moul's showcase"
	app.Commands = []cli.Command{}

	for action := range moulshowcase.Actions() {
		command := cli.Command{
			Name:   action,
			Action: CliActionCallback,
		}
		app.Commands = append(app.Commands, command)
	}

	app.Commands = append(app.Commands, cli.Command{
		Name:        "server",
		Description: "Run as a webserver",
		Action:      Daemon,
	})

	app.Run(os.Args)
}

func CliActionCallback(c *cli.Context) {
	action := c.Command.Name
	ret, err := moulshowcase.Actions()[action](c.Args())
	if err != nil {
		logrus.Fatalf("Failed to execute %q: %v", action, err)
	}

	switch ret.ContentType {
	case "application/json":
		out, err := json.MarshalIndent(ret.Body, "", "  ")
		if err != nil {
			logrus.Fatalf("Failed to marshal json: %v", err)
		}
		fmt.Printf("%s\n", out)
		return
	case "text/plain":
		fmt.Printf("%s", ret.Body)
	default:
		logrus.Fatalf("Unhandled Content-Type: %q", ret.ContentType)
	}
}

func Daemon(c *cli.Context) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		services := []string{}
		for action := range moulshowcase.Actions() {
			services = append(services, fmt.Sprintf("/%s", action))
		}
		c.JSON(200, gin.H{
			"services": services,
		})
	})
	for action, fn := range moulshowcase.Actions() {
		r.GET(fmt.Sprintf("/%s", action), func(c *gin.Context) {
			ret, err := fn(nil)
			if err != nil {
				c.JSON(500, gin.H{
					"err": err,
				})
				return
			}
			switch ret.ContentType {
			case "application/json":
				c.JSON(200, ret.Body)
				return
			case "text/plain":
				c.String(200, fmt.Sprintf("%s", ret.Body))
				return
			default:
				logrus.Fatalf("Unhandled Content-Type: %q", ret.ContentType)
			}
		})
	}
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(fmt.Sprintf(":%s", port))
}
