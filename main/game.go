package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(player,w)
	case http.MethodGet:
		p.showScore(player, w)
	}
}

func (p *PlayerServer) processWin(player string, w http.ResponseWriter) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
	return
}

func (p *PlayerServer) showScore(player string, w http.ResponseWriter) {
	score, _ := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, strconv.Itoa(score))
}


func main() {

	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	if err := http.ListenAndServe(":5000", &server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}