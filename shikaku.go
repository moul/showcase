package moulshowcase

import (
	"io"
	"math/rand"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/moul/shikaku"
)

func init() {
	RegisterAction("shikaku", ShikakuAction)
}

func ShikakuAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Width  int `schema:"width"`
		Height int `schema:"height"`
		Blocks int `schema:"blocks"`

		DrawMap         bool  `schema:"draw-map"`
		DrawSolution    bool  `schema:"draw-solution"`
		NoMachineOutput bool  `schema:"no-machine-output"`
		Srand           int64 `schema:"srand"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.Width = 8
	opts.Height = 8
	opts.Blocks = 10
	opts.DrawMap = true
	opts.DrawSolution = true
	opts.NoMachineOutput = false
	opts.Srand = 0

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

	// Create shikaku
	if opts.Srand > 0 {
		rand.Seed(opts.Srand)
	}

	shikakuMap := shikaku.NewShikakuMap(opts.Width, opts.Height, 0, 0)
	if err := shikakuMap.GenerateBlocks(opts.Blocks); err != nil {
		return nil, err
	}

	outputs := []string{}
	if !opts.NoMachineOutput {
		outputs = append(outputs, shikakuMap.String())
	}
	if opts.DrawMap {
		outputs = append(outputs, shikakuMap.DrawMap())
	}
	if opts.DrawSolution {
		outputs = append(outputs, shikakuMap.DrawSolution())
	}

	return PlainResponse(strings.Join(outputs, "\n\n")), nil
}
