package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
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
	var count int = 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, problem.q)

		//channel creation to get data
		answerCh := make(chan string)

		//concurrent function Go-routine with Anonymous function

		// Goroutine: Jab go keyword se function call kiya jata hai, 
		// to ye function alag thread mein execute hota hai, 
		// jisse main program ka flow rukta nahi hai.
		go func() {
			var answer string
			//blocking statement and a seperate goroutine to get the input
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
			//get the data from channel, blocking statement
		case answer := <-answerCh:
			{
				if answer == problem.a {
					count = count + 1
				} else {
					fmt.Printf("Incorrect Answer")
					break
				}
			}
			//another concurrent function without hampering the actual flow tracking time that returns a channel
		case <-time.After(15 * time.Second):
			{
				fmt.Printf("Times Up!")
				break
			}
		}
	}
	fmt.Printf("Your final score is %d out of %d.\n", count, len(problems))

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
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
