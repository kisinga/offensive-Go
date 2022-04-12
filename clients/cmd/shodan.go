package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kisinga/offensive-Go/clients/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits: %d\nScan Credits:  %d\n\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, val := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", val.IPString, val.Port)
	}
}
