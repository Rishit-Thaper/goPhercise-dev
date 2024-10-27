package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "There should be a csv with format of question and answer")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Problem reading the provided csv %s.csv\n", *csvFileName)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Problem reading the provided csv %s.csv\n", *csvFileName)
		os.Exit(1)
	}
	problems := createProblems(lines)
	fmt.Println(problems)
}

func createProblems(rawProblemSet [][]string) []problem {
	problemSet := make([]problem, len(rawProblemSet))
	for i, rawProblem := range rawProblemSet {
		problemSet[i] = problem{
			q: rawProblem[0],
			a: rawProblem[1],
		}
	}
	return problemSet
}