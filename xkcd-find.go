package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
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

	fn := "cmd/" + os.Args[1]

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
		log.Fatalln("Error - Can't decode file")
	}
}