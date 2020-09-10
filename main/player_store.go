package main

type PlayerStore interface {
	GetPlayerScore(name string) (int,error)
	RecordWin(name string)
}

type PlayerStubStore struct {
	scores map[string]int
	winCalls []string

}

func (p *PlayerStubStore) GetPlayerScore(name string) (int,error) {
	return p.scores[name], nil
}

func (p *PlayerStubStore) RecordWin(name string) {
	p.winCalls = append(p.winCalls, name)
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int,error) {
	return i.store[name], nil
}


