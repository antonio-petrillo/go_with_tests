package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score, err := p.store.GetPlayerScore(player)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}
}
