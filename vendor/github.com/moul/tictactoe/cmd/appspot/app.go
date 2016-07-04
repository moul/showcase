package tttapp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/moul/tictactoe"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// extract query from url
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "URL error: %v:\n", err)
		return
	}

	// parse query
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "URL query error: %v:\n", err)
		return
	}

	ttt := tictactoe.NewTicTacToe()
	ttt.Set(0, 0, m.Get("0-0"))
	ttt.Set(0, 1, m.Get("0-1"))
	ttt.Set(0, 2, m.Get("0-2"))
	ttt.Set(1, 0, m.Get("1-0"))
	ttt.Set(1, 1, m.Get("1-1"))
	ttt.Set(1, 2, m.Get("1-2"))
	ttt.Set(2, 0, m.Get("2-0"))
	ttt.Set(2, 1, m.Get("2-1"))
	ttt.Set(2, 2, m.Get("2-2"))
	ttt.SetPlayer(m.Get("you"))

	next, err := ttt.Next()
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
	} else {
		fmt.Fprintf(w, "%d-%d", next.Y, next.X)
	}
}
