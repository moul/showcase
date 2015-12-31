package moulshowcase

import (
	"io"
	"net/url"

	"github.com/getmillipede/millipede-go"
	"github.com/gorilla/schema"
)

func init() {
	RegisterAction("millipede", MillipedeAction)
}

func MillipedeAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Define arguments
	var opts struct {
		Size     uint64 `schema:"size"`
		Curve    uint64 `schema:"curve"`
		Width    uint64 `schema:"width"`
		Steps    uint64 `schema:"steps"`
		Skin     string `schema:"skin"`
		Reverse  bool   `schema:"reverse"`
		Zalgo    bool   `schema:"zalgo"`
		Opposite bool   `schema:"opposite"`
	}
	// FIXME: handle --help

	// Set defaults
	opts.Size = 0
	opts.Curve = 0
	opts.Width = 0
	opts.Steps = 0
	opts.Skin = ""
	opts.Reverse = false
	opts.Zalgo = false
	opts.Opposite = false

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

	// Create millipede
	creature := millipede.New()

	if opts.Size > 0 {
		creature.Size = opts.Size
	}
	if opts.Curve > 0 {
		creature.Curve = opts.Curve
	}
	if opts.Width > 0 {
		creature.Width = opts.Width
	}
	if opts.Steps > 0 {
		creature.Steps = opts.Steps
	}
	if opts.Skin != "" {
		creature.Skin = opts.Skin
	}
	creature.Reverse = opts.Reverse
	creature.Zalgo = opts.Zalgo
	creature.Opposite = opts.Opposite

	return PlainResponse(creature.String()), nil
}
