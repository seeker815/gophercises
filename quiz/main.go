package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var counter, tally int

	fn := flag.String("filename", "problems.csv", "Name of the csv file with Problems")
	flag.Parse()

	f, err := os.Open(*fn)
	check(err)
    defer f.Close()
	r := csv.NewReader(f)
	ri := bufio.NewReader(os.Stdin)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		// Start Quiz
		fmt.Println(record[0])
		counter++
		sol := strings.TrimSpace(record[1])
		answer, _ := ri.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.Compare(sol, answer) == 0 {
			tally++
		}
	}

	f.Close()

	fmt.Println("Total questions in the test and correctly answered", counter, tally)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
