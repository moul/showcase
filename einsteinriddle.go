package moulshowcase

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/moul/einstein-riddle-generator"
)

func init() {
	RegisterAction("einstein-riddle", EinsteinriddleAction)
}

func EinsteinriddleAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Size int `schema:"size"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.Size = 0

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

	// Create einsteinriddle
	/*
		if opts.Srand > 0 {
			rand.Seed(opts.Srand)
		}
	*/
	options := einsteinriddle.Options{
		Size: opts.Size,
	}
	riddle := einsteinriddle.NewGenerator(options)

	if err := riddle.Shazam(); err != nil {
		return nil, err
	}

	lines := []string{}

	lines = append(lines, "Facts:")
	for _, group := range riddle.Pickeds {
		lines = append(lines, fmt.Sprintf("- %s", riddle.GroupString(group)))
	}

	lines = append(lines, "Questions:")
	for _, item := range riddle.Missings() {
		lines = append(lines, fmt.Sprintf("- Where is %s ?", item.Name()))
	}

	return PlainResponse(strings.Join(lines, "\n")), nil
}
