package libdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Tick is a Tick is a tick is a tick
type Tick struct {
	Timestamp int64
	Value     int
}

// Database is a type for storing stats in
type Database struct {
	name     string
	interval int
	rows     int
	file     string
	child    *Database
	ticks    []Tick
}

// Info prints information about the DB
func (db *Database) Info() {
	fmt.Println(db)
}

// Append an item to the DB according to intervall and row rules
func (db *Database) Append(ticker Tick) {
	if len(db.ticks) > 0 {
		if ticker.Timestamp-db.ticks[len(db.ticks)-1].Timestamp < int64(db.interval) {
			log.Printf("%v is within the less than %v from last entry. Not added\n", ticker.Timestamp, db.interval)
		} else {
			if len(db.ticks) >= db.rows {
				db.ticks = append(db.ticks[:0], db.ticks[1:]...)
				db.ticks = append(db.ticks, ticker)
			} else {
				db.ticks = append(db.ticks, ticker)
			}
		}
	} else {
		db.ticks = append(db.ticks, ticker)
	}
}

// SetName sets the Name of the Database
func (db *Database) SetName(name string) {
	db.name = name
}

// Name returns the name of the Database
func (db *Database) Name() string {
	return db.name
}

// SetRows sets the number of rows the database can handle before overflowing
func (db *Database) SetRows(rows int) {
	db.rows = rows
}

// Rows returns the number of rows the database can handle before overflowing
func (db *Database) Rows() int {
	return db.rows
}

// SetInterval sets the interval in SECONDS between entries in the database.
// Entries with a lower interval will be ignored
func (db *Database) SetInterval(interval int) {
	db.interval = interval
}

// Interval returns the interval of the database
func (db *Database) Interval() int {
	return db.interval
}

// SetFile sets the path to the file where to store the db
func (db *Database) SetFile(file string) {
	db.file = file
}

// File returns the file-string of the database
func (db *Database) File() string {
	return db.file
}

// GetAverage gets the average value of the ticks.Value
func (db *Database) GetAverage() int {
	var average int
	for _, ticker := range db.ticks {
		average += ticker.Value
	}
	average = average / len(db.ticks)
	return average
}

// Save writes the db as json to the file specified db.file
func (db *Database) Save() {
	dBytes, err := json.Marshal(db)
	if err != nil {
		log.Printf("%v is not valid json. Something is really broken", db)
	}
	ioutil.WriteFile("db.json", dBytes, 0600)
	if err != nil {
		log.Printf("Unable to save database to %v", db.file)
	}
}
