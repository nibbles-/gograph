package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	body, _ := ioutil.ReadAll(r.Body)
	// Create a regexp to be able to generate unique devicenames
	re := regexp.MustCompile("[A-Za-z]*$")
	fmt.Printf("Connection: %v %v %v %v\n", time.Now(), string(body), r.Host, re.FindString(string(body)))
	fmt.Fprintf(w,
		`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
    <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <soapenv:Body>
      <ns1:perfmonCollectCounterDataResponse xmlns:ns1="http://schemas.cisco.com/ast/soap/" soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
      <ArrayOfCounterInfo xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/" soapenc:arrayType="ns1:CounterInfoType[40]" xsi:type="soapenc:Array">
        <item xsi:type="ns1:CounterInfoType">
          <Name xsi:type="ns1:CounterNameType">\\%[1]v\%[4]v(%[5]v_Trunk_1)\CallsInProgress</Name>
          <Value xsi:type="xsd:long">%[2]v</Value>
          <CStatus xsi:type="xsd:unsignedInt">1</CStatus>
        </item>
        <item xsi:type="ns1:CounterInfoType">
          <Name xsi:type="ns1:CounterNameType">\\%[1]v\%[4]v(%[5]v_Trunk_2)\CallsInProgress</Name>
          <Value xsi:type="xsd:long">%[3]v</Value>
          <CStatus xsi:type="xsd:unsignedInt">1</CStatus>
        </item>
      </ArrayOfCounterInfo>
      </ns1:perfmonCollectCounterDataResponse>
    </soapenv:Body>
    </soapenv:Envelope>`, r.Host, rand.Intn(100), rand.Intn(100), string(body), re.FindString(string(body)))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
