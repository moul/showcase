package moulshowcase

import (
	"bufio"
	"io"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/moul/anonuuid"
)

func init() {
	RegisterAction("anonuuid", AnonuuidAction)
}

func AnonuuidAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Hexspeak      bool   `schema:"hexspeak"`
		Random        bool   `schema:"random"`
		Prefix        string `schema:"prefix"`
		Suffix        string `schema:"suffix"`
		KeepBeginning bool   `schema:"keep-beginning"`
		KeepEnd       bool   `schema:"keep-end"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.Random = true

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

	// Create anonuuid
	encoder := anonuuid.New()
	encoder.Hexspeak = opts.Hexspeak
	encoder.Random = opts.Random
	encoder.Prefix = opts.Prefix
	encoder.Suffix = opts.Suffix
	encoder.KeepBeginning = opts.KeepBeginning
	encoder.KeepEnd = opts.KeepEnd

	scanner := bufio.NewScanner(stdin)
	output := ""
	for scanner.Scan() {
		line := scanner.Text()
		output += encoder.Sanitize(line) + "\n"
	}

	return PlainResponse(output), nil
}
