package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file containing data for quiz; "+
		"format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "amount of time to give the quiz taker in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the given csv file, %s\n", *csvFileName))
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the csv file."))
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// correct count
	correct := 0

	// iterate through and ask the problems and keep count
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			_, _ = fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	probs := make([]problem, len(lines))

	for i, line := range lines {
		probs[i] = problem{
			question: line[0],
			answer:   strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return probs
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
