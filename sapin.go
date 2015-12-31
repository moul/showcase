package moulshowcase

import (
	"fmt"
	"io"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/moul/sapin"
)

func init() {
	RegisterAction("sapin", SapinAction)
}

func SapinAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Size     int  `schema:"size"`
		Balls    int  `schema:"balls"`
		Garlands int  `schema:"garlands"`
		Star     bool `schema:"star"`
		Emoji    bool `schema:"emoji"`
		Color    bool `schema:"color"`
		Presents bool `schema:"presents"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.Size = 5
	opts.Color = true
	opts.Balls = 4
	opts.Star = true
	opts.Emoji = false
	opts.Presents = true
	opts.Garlands = 5

	// Parse query
	m, err := url.ParseQuery(qs)
	if err != nil {
		return nil, err
	}
	if len(m) > 0 {
		// FIXME: if in web mode -> redirect to have options demo in the URL
		decoder := schema.NewDecoder()
		if err := decoder.Decode(&opts, m); err != nil {
			return nil, err
		}
	}

	// FIXME: validate input (max size etc)

	// Create sapin
	sapin := sapin.NewSapin(opts.Size)

	// Apply options
	if opts.Star {
		sapin.AddStar()
	}
	sapin.AddBalls(opts.Balls)
	sapin.AddGarlands(opts.Garlands)
	if opts.Emoji {
		sapin.Emojize()
	}
	if opts.Color {
		// FIXME: handle HTML - see old sapin for example
		sapin.Colorize()
	}
	if opts.Presents {
		sapin.AddPresents()
	}

	return PlainResponse(fmt.Sprintf("%s", sapin)), nil
}
