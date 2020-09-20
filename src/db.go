package main

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
)

var conn *dbr.Connection

func dbInit() {

	// create a connection (e.g. "postgres", "mysql", or "sqlite3")
	conn, err := dbr.Open("sqlite3", "./test.sqlite", nil)
	if err != nil {
		fmt.Println("Error connecting: ", err)
	}
	conn.SetMaxOpenConns(10)

	// create a session for each business unit of execution (e.g. a web request or goworkers job)
	sess := conn.NewSession(nil)

	// create a tx from sessions
	tx, err := sess.Begin()
	if err != nil {
		fmt.Printf("Error creating sessoin")
		return
	}
	// columns are mapped by tag then by field
	type SearchEntry struct {
		ID        int64 // id, will be autoloaded by last insert id
		Title     string
		Body      string
		CreatedAt time.Time
		secret    string `db:"-"` // ignored
	}
	// sugg := &SearchEntry{
	// 	Title:     "gopher",
	// 	CreatedAt: time.Now(),
	// }
	// actually do db ops
	tx.Commit()
}
