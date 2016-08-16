package main

// make needed imports
import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	// read settings
	settings := configuration{}
	settings.Username = "administrator"
	settings.Password = "password"
	settings.Servers = append(settings.Servers, "127.0.0.1:8080", "127.0.0.2:8080")
	settings.Counters = append(settings.Counters, "Cisco SIP", "Cisco MGCP Gateways", "Cisco MGCP PRI Device")

	// load database into to a map
	var err error // Error container
	var dBytes []byte

	dBytes = dbFileReadCreate("database.json")
	var mapStore = map[string][]tick{}
	err = json.Unmarshal(dBytes, &mapStore)
	check(err)

	// get data from cucm
	client := &http.Client{}
	var result = map[string]int{} // Init empty resultmap to contain the totals of all counters
	for _, counter := range settings.Counters {
		soaprequest := []byte(fmt.Sprintf("%v", counter))
		for _, server := range settings.Servers {
			perfmonresult := soapResponse{}
			url := fmt.Sprintf("http://%v/perfmonservice/services/PerfmonPort", server)
			request, _ := http.NewRequest("POST", url, bytes.NewBuffer(soaprequest))
			request.Header.Set("SOAPAction", "perfmonCollectCounterData")
			response, err := client.Do(request)
			check(err)
			defer response.Body.Close()
			responseBody, _ := ioutil.ReadAll(response.Body)
			err = xml.Unmarshal(responseBody, &perfmonresult)
			check(err)
			// Add current request results to the resultmap
			for _, item := range perfmonresult.Soap.PerfmonCollectCounterData.Item {
				device := strings.Split(item.Name, "\\")
				fmt.Println(device)
				result[device[3]] = result[device[3]] + item.Value // device[3] because thats where "Counter(Device)" ends up
			}
		}
	}
	// save result to the database map
	for key, value := range result {
		ticker := tick{fmt.Sprint(time.Now().Unix()), value}
		mapStore[key] = append(mapStore[key], ticker)
	}

	// encode into json and write to database.json
	dBytes, err = json.Marshal(mapStore)
	check(err)
	ioutil.WriteFile("database.json", dBytes, 0600)
	check(err)
	// read html template
	// put data in html files
	//fmt.Println(settings.Username, settings.Password, settings.Servers, settings.Counters)
	//fmt.Println(mapStore)
}
