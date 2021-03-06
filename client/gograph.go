package main

// make needed imports
import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// read settings
	settings := configuration{}
	settings.Username = "administrator"
	settings.Password = "password"
	//settings.Servers = append(settings.Servers, "127.0.0.1:8080", "192.168.1.100", "127.0.0.2:8080", "192.168.1.110")
	settings.Servers = append(settings.Servers, "127.0.0.1:8080", "127.0.0.2:8080")
	settings.Counters = append(settings.Counters, "Cisco SIP", "Cisco MGCP Gateways", "Cisco MGCP PRI Device")

	// LoadCreate a database
	db := Database{}
	db.Name = "LeDatabase"
	db.File = "LeFile.json"

	// Create a client with a 10 second timeout
	client := &http.Client{Timeout: time.Second * 10}
	// Init empty resultmap to contain the totals of all counters
	var result = map[string]int{}
	// Get data from cucm
	for _, counter := range settings.Counters {
		soaprequest := []byte(fmt.Sprintf("%v", counter))
		for _, server := range settings.Servers {
			perfmonresult := soapResponse{}
			url := fmt.Sprintf("http://%v/perfmonservice/services/PerfmonPort", server)
			request, _ := http.NewRequest("POST", url, bytes.NewBuffer(soaprequest))
			request.Header.Set("SOAPAction", "perfmonCollectCounterData")
			response, err := client.Do(request)
			if err != nil {
				// If client.Do generates an error log it and move on.
				log.Println(err)
				continue
			}
			defer response.Body.Close()
			responseBody, _ := ioutil.ReadAll(response.Body)
			err = xml.Unmarshal(responseBody, &perfmonresult)
			check(err)
			// Add current request results to the resultmap
			for _, item := range perfmonresult.Soap.PerfmonCollectCounterData.Item {
				devicestring := []string{}
				// Create a regexp to be able to generate unique devicenames
				switch counter {
				// If we are looking at a SIP device we want to use CallsInProgress
				case "Cisco SIP":
					devicestring = regSip.FindStringSubmatch(item.Name)
					// If we are looking at a MGCP GW we want to use PRIChannelsActive
				case "Cisco MGCP Gateways":
					devicestring = regMgcpGw.FindStringSubmatch(item.Name)
					// If we are looking at a MGCP PRI we want to use CallsActive
				case "Cisco MGCP PRI Device":
					devicestring = regMgcpPri.FindStringSubmatch(item.Name)
				default:
					log.Panic("Unsupported Counter: ", counter)
				}
				// We only save matched values (i.e devicestring is not empty)
				if len(devicestring) > 0 {
					device := devicestring[1]
					result[device] = result[device] + item.Value // add current device and value to result
				}
			}
		}
	}
	//fmt.Println(result)

	// save result to the database
	for key, value := range result {
		ticker := Tick{Timestamp: time.Now().Unix(), Value: value}
		weekTbl := db.NewTable(key+"_week", 10, 5, nil)
		tbl := db.NewTable(key, 1, 10, weekTbl)
		tbl.Append(ticker)
		time.Sleep(2 * time.Second)
	}

	fmt.Println(db)
	db.Save()

	// read html template
	// put data in html files
	//fmt.Println(settings.Username, settings.Password, settings.Servers, settings.Counters)
	//fmt.Println(mapStore)
}
