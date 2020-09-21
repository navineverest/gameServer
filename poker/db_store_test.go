package poker

import (
	"testing"
)

func TestDBStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {

		store := DBPlayerStore{}

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {

		store := DBPlayerStore{}

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {

		store, _ := NewDBStore()

		store.RecordWin("Navin")
		store.RecordWin("Navin")

		got := store.GetPlayerScore("Navin")
		want := 2
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {

		store := DBPlayerStore{}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("league sorted", func(t *testing.T) {

		store := DBPlayerStore{}

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

}
