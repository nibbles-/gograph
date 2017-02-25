package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Used to handle errors in a nice way
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read dbFile and Create if not existing
func dbFileReadCreate(file string) []byte {
	data, e := ioutil.ReadFile(file)
	if e != nil && os.IsNotExist(e) {
		log.Print("Database not found. Creating a new one")
		ioutil.WriteFile(file, []byte("{}"), 0600)
		data, e = ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}
	} else if e != nil {
		panic(e)
	}
	return data
}

// Info prints the database.
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
