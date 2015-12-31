package moulasaservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/moul/showcase"
)

func init() {
	http.HandleFunc("/", indexHandler)

	for name := range moulshowcase.Actions() {
		http.HandleFunc(fmt.Sprintf("/%s", name), actionHandler)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var payload struct {
		Services []string `json:"services"`
	}
	payload.Services = make([]string, 0)
	for action := range moulshowcase.Actions() {
		payload.Services = append(payload.Services, fmt.Sprintf("/%s", action))
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(payload); err != nil {
		http.Error(w, fmt.Sprintf("json encode error: %v\n", err), 500)
	}
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	if fn, found := moulshowcase.Actions()[path]; found {
		// parse CLI arguments
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse url %q: %v", r.URL.String(), err), 500)
		}

		// call action
		ret, err := fn(u.RawQuery, r.Body)

		// render result
		if err != nil {
			http.Error(w, fmt.Sprintf("service error: %v\n", err), 500)
		} else {
			switch ret.ContentType {
			case "application/json":
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				enc := json.NewEncoder(w)
				if err := enc.Encode(&(ret.Body)); err != nil {
					http.Error(w, fmt.Sprintf("json encode error: %v\n", err), 500)
				}
			case "text/plain":
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				fmt.Fprintf(w, "%s", ret.Body)
			}
		}
	} else {
		http.NotFound(w, r)
	}
}
