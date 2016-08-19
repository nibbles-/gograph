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
	Ticks     []tick
}

// Info prints information about the DB
func (db *Database) Info() {
	fmt.Println(db)
}

// Append an item to the DB according to intervall and row rules
func (db *Database) Append(ticker tick) {
	if len(db.Ticks) > 0 {
		if ticker.Timestamp-db.Ticks[len(db.Ticks)-1].Timestamp < db.Intervall {
			log.Printf("%v is within the limit. Not added\n", ticker.Timestamp)
		} else {
			if len(db.Ticks) >= db.Rows {
				fmt.Println(db.Ticks)
				db.Ticks = append(db.Ticks[:0], db.Ticks[1:]...)
				db.Ticks = append(db.Ticks, ticker)
			} else {
				fmt.Println(db.Ticks)
				db.Ticks = append(db.Ticks, ticker)
			}
		}
	} else {
		db.Ticks = append(db.Ticks, ticker)
	}
}
