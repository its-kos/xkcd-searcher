package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Transcript string `json:"transcript"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}


func main() {
	
	if len(os.Args) < 3 {
		log.Fatalln("Error - No file and/or search term given")
	}

	fn := os.Args[1]

	var (
		comics []Comic
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	if input, err = os.Open(fn); err != nil {
		log.Fatalln("Error - Can't find file")
	}

	if err = json.NewDecoder(input).Decode(&comics); err != nil {
		log.Fatalln("Error - Can't decode file ", err)
	}

	fmt.Printf("Read %d comics\n", len(comics))

	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

outer:
	for _, item := range comics {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}
		fmt.Printf("A relevant one: https://xkcd.com/%d/ %s/%s/%s %q\n", item.Num, item.Day, item.Month, item.Year, item.Title)
		cnt++
	}
}