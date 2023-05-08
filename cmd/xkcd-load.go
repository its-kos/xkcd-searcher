package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const url = "https://xkcd.com/"

func main() {

	var (
		output io.WriteCloser = os.Stdout
		err error
		fails = 0
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			log.Fatalln("Error creating file - ", err)
		}
		defer output.Close()
	}

	if _, err = io.Copy(output, strings.NewReader("[")); err != nil {
		log.Fatalln("Error getting URL: ", err)
	}

	for i := 1; fails < 2; i++ {

		nextURL := url + strconv.Itoa(i) + "/info.0.json"

		resp, err := http.Get(nextURL)
		if err != nil {
			log.Fatalln("Error getting URL: ", err)
		}
		
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Skipping #%d, got status code %d\n", i, resp.StatusCode)
			fails++
			resp.Body.Close()
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Can't read body: ", err)
		}

		fmt.Fprint(output, string(body) + ",")
		_, err = io.Copy(output, bytes.NewBuffer(body))
		if err != nil {
			log.Fatalln("Can't write body: ", err)
		}

		fails = 0
		resp.Body.Close()
	}
	if _, err = io.Copy(output, strings.NewReader("]")); err != nil {
		log.Fatalln("Error getting URL: ", err)
	}
}