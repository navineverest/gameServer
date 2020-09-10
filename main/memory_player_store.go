package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return i.store[name], nil
}

func (i *InMemoryPlayerStore) GetLeaguePlayers() []Player {
	players := []Player{}
	for name, wins := range i.store {
		players = append(players, Player{Name: name, Wins: wins})
	}
	return players

}
