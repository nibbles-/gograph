package main

// tick is an entry in the database
type tick struct {
	Timestamp int64
	Value     int
}

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
