package main

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (p *StubPlayerStore) GetPlayerScore(name string) int {
	return p.scores[name]
}

func (p *StubPlayerStore) RecordWin(name string) {
	p.winCalls = append(p.winCalls, name)
}

func (p *StubPlayerStore) GetLeague() League {
	return p.league
}
