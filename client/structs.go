package main

import "encoding/xml"

// For settings in the program
type configuration struct {
	Username string
	Password string
	Servers  []string
	Counters []string
}

// --- Struct for the XML response from CUCM ---
type soapResponse struct {
	// Envelope part of the soap response
	XMLName xml.Name `xml:"Envelope"`
	Soap    *soapBody
}
type soapBody struct {
	XMLName                   xml.Name `xml:"Body"`
	PerfmonCollectCounterData *perfmonCollectCounterDataResponse
}
type perfmonCollectCounterDataResponse struct {
	XMLName xml.Name `xml:"perfmonCollectCounterDataResponse"`
	Item    []item   `xml:"ArrayOfCounterInfo>item"`
}

type item struct {
	XMLName xml.Name `xml:"item"`
	Name    string
	Value   int
	CStatus string
}

// ---------------------------------------------

// tick is an entry in the database
type tick struct {
	Timestamp int64
	Value     int
}
