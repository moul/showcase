package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/moul/showcase"
	"github.com/stvp/rollbar"
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

	// parse CLI arguments
	args := []string{}
	for _, arg := range c.Args() {
		if arg[:2] == "--" {
			args = append(args, arg[2:])

		} else {
			args = append(args, fmt.Sprintf("arg=%s", arg))
		}
	}
	qs := strings.Join(args, "&")

	// call action
	ret, err := moulshowcase.Actions()[action](qs, os.Stdin)
	if err != nil {
		logrus.Fatalf("Failed to execute %q: %v", action, err)
	}

	// render result
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

	rollbar.Token = os.Getenv("ROLLBAR_TOKEN")
	rollbar.Environment = "production"

	rollbar.Message("info", "Starting daemon")

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
		func(action string, fn moulshowcase.Action) {
			callback := func(c *gin.Context) {
				u, err := url.Parse(c.Request.URL.String())
				if err != nil {
					rollbar.Error(rollbar.ERR, err)
					rollbar.Wait()
					c.String(500, fmt.Sprintf("failed to parse url %q: %v", c.Request.URL.String(), err))
					return
				}

				ret, err := fn(u.RawQuery, c.Request.Body)
				if err != nil {
					rollbar.Error(rollbar.ERR, err)
					rollbar.Wait()
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
					err := fmt.Errorf("Unhandled Content-Type: %q", ret.ContentType)
					rollbar.Error(rollbar.ERR, err)
					rollbar.Wait()
					logrus.Fatal(err)
				}
			}
			r.GET(fmt.Sprintf("/%s", action), callback)
			r.POST(fmt.Sprintf("/%s", action), callback)
			//r.PUT(fmt.Sprintf("/%s", action), callback)
			//r.PATCH(fmt.Sprintf("/%s", action), callback)
			//r.DELETE(fmt.Sprintf("/%s", action), callback)
		}(action, fn)
	}
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(fmt.Sprintf(":%s", port))
}
