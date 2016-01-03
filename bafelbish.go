package moulshowcase

import (
	"bytes"
	"io"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/moul/bafelbish"
)

func init() {
	RegisterAction("bafelbish", BafelbishAction)
}

func BafelbishAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		InputFormat  string `schema:"input-format"`
		OutputFormat string `schema:"output-format"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.InputFormat = "json"
	opts.OutputFormat = "json"

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

	// Create bafelbish
	encoder := bafelbish.NewFish()

	encoder.SetInputFormat(opts.InputFormat)
	encoder.SetOutputFormat(opts.OutputFormat)

	var b bytes.Buffer
	if err := encoder.Parse(stdin, &b); err != nil {
		return nil, err
	}

	return PlainResponse(b.String()), nil
}
