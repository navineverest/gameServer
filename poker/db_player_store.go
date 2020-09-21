package poker

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "game"
)

func NewDBStore() (*DBPlayerStore, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return &DBPlayerStore{*db}, err
}

type DBPlayerStore struct {
	db sql.DB
}

func (d *DBPlayerStore) RecordWin(name string) {
	sqlStatement := `SELECT player, score FROM scores WHERE player=$1;`
	var player string
	var score int
	row := d.db.QueryRow(sqlStatement, name)
	switch err := row.Scan(&player, &score); err {
	case sql.ErrNoRows:
		sqlStatement = `INSERT INTO scores (player, score)VALUES ($1, $2)`
		_, err = d.db.Exec(sqlStatement, name, 1)
		if err != nil {
			panic(err)
		}
	case nil:
		sqlStatement := `UPDATE scores SET score = score+1 WHERE player = $1;`
		_, err := d.db.Exec(sqlStatement, name)
		if err != nil {
			panic(err)
		}
	default:
		panic(err)
	}
}

func (d *DBPlayerStore) GetPlayerScore(name string) int {
	sqlStatement := `SELECT player, score FROM scores WHERE player=$1;`
	var player string
	var score int
	row := d.db.QueryRow(sqlStatement, name)
	switch err := row.Scan(&player, &score); err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return score
	default:
		panic(err)
	}
}

func (d *DBPlayerStore) GetLeague() League {
	return League{}
}
