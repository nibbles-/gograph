package main

import (
	"fmt"
	"log"
)

// Database is a type for storing stats in
type Database struct {
	Name      string
	Intervall int64
	Rows      int
	ticks     []tick
}

// Info prints information about the DB
func (db *Database) Info() {
	fmt.Println(db)
}

// Append an item to the DB according to intervall and row rules
func (db *Database) Append(ticker tick) {
	if ticker.Timestamp-db.ticks[len(db.ticks)].Timestamp < db.Intervall {
		log.Printf("%v is within the limit. Not added\n", ticker.Timestamp)
	} else {
		if len(db.ticks) >= db.Rows {
			db.ticks = append(db.ticks[:0], db.ticks[1:]...)
			db.ticks = append(db.ticks, ticker)
		} else {
			db.ticks = append(db.ticks, ticker)
		}
	}
}
