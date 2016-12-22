package main

import (
	"blueprints/ch4/thesaurus"

	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Using environment variables for configuration
	//apiKey := os.Getenv("BHT_APIKEY")
	apiKey := "159974fb005b2b9a8d319c418b704cf2"
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for \""+word+"\"", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for \"" + word + "\"")
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
