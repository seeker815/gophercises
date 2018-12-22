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
	"time"
)

var (
	fn             *string
	qt             *time.Duration
	counter, tally int
)

func main() {
	fn = flag.String("csv", "problems.csv", "Name of the csv file with Problems")
	qt = flag.Duration("timer", 10, "Quiz Time duration")
	flag.Parse()

	afterFuncTimer := time.AfterFunc(time.Second**qt, func() {
		fmt.Printf("\n Please put your pencils down, Quiz time is over!!! \n")
		fmt.Printf("\n Total number of Questions = %d", counter)
		fmt.Printf("\n Total number of Correctly answered questions = %d \n", tally)
		os.Exit(0)
	})

	defer afterFuncTimer.Stop()

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
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
