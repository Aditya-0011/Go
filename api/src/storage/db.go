package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	query := `
	   create table if not exists todos (
        id integer primary key autoincrement,
        title text not null,
        done boolean not null check (done in (0,1))
    );
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
