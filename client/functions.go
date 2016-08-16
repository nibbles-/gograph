package main

import (
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
		ioutil.WriteFile("database.json", []byte("{}"), 0600)
		data, e = ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}
	} else if e != nil {
		panic(e)
	}
	return data
}
