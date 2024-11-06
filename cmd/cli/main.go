package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	quiz "quiz-game"
	"time"
)

const defaultCSVFileName = "problems.csv"
const defaultTimeLimit = 2

func parseFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var problemStrings []string
	for scanner.Scan() {
		problemStrings = append(problemStrings, scanner.Text())
	}
	return problemStrings[1:]
}

func main() {
	csvFileName := flag.String("csv", defaultCSVFileName, "a csv file in the format of 'question;asnwer'")
	timeLimit := flag.Int("limit", defaultTimeLimit, "time limit of a quiz game")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("got error %s", err)
	}
	problems := quiz.ConvertStringsToProblems(parseFile(file))
	game := quiz.NewQuiz(problems, time.Duration(*timeLimit)*time.Second)
	cli := quiz.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayQuiz()
}
