package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Hit struct {
	Created  int64
	Id       string
	Msisdn   string
	Site     string
	Themes   string
	Resource string
	Provider string
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	inputFile, err := os.Open("syslog")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// scanner.Scan() advances to the next token returning false if an error was encountered
	//	var records [][]string
	var nextline bool = false
	var msisdn string
	for scanner.Scan() {
		line := scanner.Text()

		linesplit := strings.Split(line, " ")

		//		fmt.Println(linesplit[5])

		if linesplit[5] == "msisdn" {
			nextline = true
			msisdn = linesplit[6]
		} else {
			nextline = false
		}

		if !nextline {
			if len(linesplit) > 11 {
				if linesplit[12] == "adult" {
					outline := "curl -G -H \"X-UP-CALLING-LINE-ID: " + msisdn + "\" -d \"site=" + linesplit[8] + "\" -d \"id=" + linesplit[6] + "\" -d \"resource=" + linesplit[10] + "\" -d \"themes=adult\" -d \"provider=sonera\" 127.0.0.1/sonera?callback=lslslslsl"
					//			fmt.Println("curl -G -H \"X-UP-CALLING-LINE-ID: \"",msisdn,linesplit[6],linesplit[8],linesplit[10],linesplit[12])
					fmt.Println(outline)
				}

			}
		}

	}

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}

}
