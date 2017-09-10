package main

import "encoding/xml"

// --- Struct for the Soap-XML response from CUCM ---
// Written downward / inward.
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
