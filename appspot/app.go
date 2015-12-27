package moulasaservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/moul/showcase"
)

func init() {
	for name := range moulshowcase.Actions() {
		http.HandleFunc(fmt.Sprintf("/%s", name), actionHandler)
	}
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	if fn, found := moulshowcase.Actions()[path]; found {
		args := []string{}
		ret, err := fn(args)

		if err != nil {
			http.Error(w, fmt.Sprintf("service error: %v\n", err), 500)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			enc := json.NewEncoder(w)
			if err := enc.Encode(&ret); err != nil {
				http.Error(w, fmt.Sprintf("json encode error: %v\n", err), 500)
			}
		}
	} else {
		http.NotFound(w, r)
	}
}
