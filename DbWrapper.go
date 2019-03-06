package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type DbWrapper struct {
	db *sql.DB
}

func (dbw *DbWrapper) initialize(filename string) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("create table if not exists flight_data(" +
		"id integer primary key autoincrement," +
		"data_type integer not null," +
		"time timestamp default current_timestamp," +
		"data_string text" +
		")")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	dbw.db = db
}

func (dbw *DbWrapper) insert(dataType int, data string) {
	stmt, err := dbw.db.Prepare("insert into flight_data(data_type, data_stirng ) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(dataType, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (dbw *DbWrapper) find(dataType int, date time.Time) {
	//stmt, err := dbw.db.Prepare("select * from ")
}

func (dbw *DbWrapper) close() {
	dbw.db.Close()
}
