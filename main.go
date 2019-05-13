package quizboi

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file containing data for quiz; "+
		"format of 'question,answer'")
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

	// correct count
	correct := 0

	// iterate through and ask the problems and keep count
	for i, p := range problems {
		correct = askQuestion(i, p, correct)
	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

func askQuestion(i int, p problem, correct int) int {
	fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
	var answer string
	_, _ = fmt.Scanf("%s\n", &answer)
	if answer == p.answer {
		correct++
	}
	return correct
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	probs := make([]problem, len(lines))

	for i, line := range lines {
		probs[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return probs
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
