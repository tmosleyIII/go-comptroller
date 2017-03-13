package main

import (
	"database/sql"
	"log"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	db, err = sql.Open("sqlite3", path.Join(databaseDir, "data.db"))
	if err != nil {
		log.Fatal(err)
	}

}

func addRow() {
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
