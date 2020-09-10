package main

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (p *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	return p.scores[name], nil
}

func (p *StubPlayerStore) RecordWin(name string) {
	p.winCalls = append(p.winCalls, name)
}

func (p *StubPlayerStore) GetLeaguePlayers() []Player {
	return p.league
}
