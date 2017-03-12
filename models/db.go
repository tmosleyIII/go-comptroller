package models

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBCon *sql.DB
)

func InitDB(dataSourceName string) {
	os.Remove(path.Join(dataSourceName, "data.db"))

	db, err := sql.Open("sqlite3", path.Join(dataSourceName, "data.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("1", "TEST")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
