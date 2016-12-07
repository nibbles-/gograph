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

type table struct {
	Name     string
	Interval int
	Rows     int
	Overflow *table
	Ticks    []Tick
}

// Database is a type for storing stats in
type Database struct {
	Name   string
	File   string
	Tables []*table
}

// Info prints information about the DB
func (db *Database) Info() {
	fmt.Println(db)
}

// Append an item to the DB according to intervall and row rules
func (tbl *table) Append(ticker Tick) {
	if len(tbl.Ticks) > 0 {
		if ticker.Timestamp-tbl.Ticks[len(tbl.Ticks)-1].Timestamp < int64(tbl.Interval) {
			log.Printf("%v is less than %v from last entry. Not added\n", ticker.Timestamp, tbl.Interval)
		} else {
			if len(tbl.Ticks) >= tbl.Rows {
				tbl.Ticks = append(tbl.Ticks[:0], tbl.Ticks[1:]...)
				tbl.Ticks = append(tbl.Ticks, ticker)
				if tbl.Overflow != nil {
					// Calculate average of the table and overflow it
					tblAverage := tbl.GetAverage()
					ticker.Value = tblAverage
					tbl.Overflow.Append(ticker)
				}
			} else {
				tbl.Ticks = append(tbl.Ticks, ticker)
			}
		}
	} else {
		tbl.Ticks = append(tbl.Ticks, ticker)
	}
}

// NewTable creates a new table in the Database
func (db *Database) NewTable(name string, interval int, rows int, overflow *table) *table {
	tbl := table{Name: name, Interval: interval, Rows: rows, Overflow: overflow}
	db.Tables = append(db.Tables, &tbl)
	return &tbl
}

// GetAverage gets the average value of the ticks.Value
func (tbl *table) GetAverage() int {
	var average int
	for _, ticker := range tbl.Ticks {
		average += ticker.Value
	}
	average = average / len(tbl.Ticks)
	return average
}

// Save writes the db as json to the file specified db.file
func (db *Database) Save() {
	dBytes, err := json.Marshal(db)
	if err != nil {
		log.Printf("%v is not valid json. Something is really broken", db)
	}
	ioutil.WriteFile(db.File, dBytes, 0600)
	if err != nil {
		log.Printf("Unable to save database to %v", db.File)
	}
}
