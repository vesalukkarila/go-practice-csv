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

	//var allTlds []TLD

	tlds, err := getTldsForType("country")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(tlds)

	// TODO...
	//allTlds = append(allTlds, tlds...)
}

func getTldsForType(tldType string) ([]TLD, error) {
	var tlds []TLD
	var err error

	filePath := fmt.Sprintf("data/%s.csv", tldType)

	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return tlds, err
	}

	tlds, err = parseTlds(f)
	if err != nil {
		return tlds, err
	}
	return tlds, nil
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
