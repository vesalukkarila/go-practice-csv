package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type TLD struct {
	Domain      string
	Description string
	Type        string
}

func main() {

	f, err := os.Open("data/country.csv")
	defer f.Close()

	if err != nil {
		log.Panic(err)
	}

	tlds, err := parseTlds(f)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(tlds)
}

func parseTlds(source io.Reader) ([]TLD, error) { //io.Reader = interface

	var tlds []TLD
	var err error

	scanner := bufio.NewScanner(source) //study

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")

		tld := TLD{
			Domain:      parts[0],
			Description: parts[1],
			Type:        parts[2],
		}

		tlds = append(tlds, tld)
	}

	return tlds, err
}
