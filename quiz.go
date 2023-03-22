package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Quiz Game!")
	var filename = flag.String("filename", "problems.csv", "filename to use")
	var timeLimit = flag.Int("timeLimit", 30, "time limit for the quiz")
	var shuffle = flag.Bool("shuffle", false, "shuffle questions?")
	flag.Parse()

	fmt.Println("filename:", *filename, "| time limit:", *timeLimit, "| shuffle:", *shuffle)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	allQuestions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// shuffle questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allQuestions), func(i, j int) { allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i] })

	correct := 0
	fmt.Println("Ready to go? Press Enter to start timer.")
	fmt.Scanln()

	// set timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	go func() {
		<-timer.C
		fmt.Println("\nTime's up!")
		fmt.Println("Total questions:", len(allQuestions))
		fmt.Println("Correct:", correct)
		os.Exit(0)
	}()

	for _, rec := range allQuestions {
		// ask question
		fmt.Print(rec[0], " ")
		var answer string
		fmt.Scan(&answer)
		if cleanString(answer) == cleanString(rec[1]) {
			correct++
		}
	}

    // stop timer if questions are finished before time's up
	stop := timer.Stop()
	if stop {
		fmt.Println("Timer stopped!")
	}

	fmt.Println("Total questions:", len(allQuestions))
	fmt.Println("Correct:", correct)
}

func cleanString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
