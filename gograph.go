package main

// make needed imports
import (
	"fmt"
	"flag"
)

var (
	flag1 = flag.String("flag1", "default1", "Description1")
	flag2 = flag.Int("flag2", 123, "Description2")
	flag3 = flag.String("flag3", "Default3", "Description3")
)

func init(){
	// read settings from settings file and flags
	flag.Parse()
}
func main(){
	// get data from cucm
	// save data to a storage format
	// read data from storage format
	// read html template
	// put data in html files
	fmt.Println(*flag1)
	fmt.Println(*flag2)
	fmt.Println(*flag3)		
}

