package main

import (
	"bufio"
	"fmt"
	"os"
	quiz "quiz-game"
)

const csvFileName = "problems.csv"

func parseFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var problemStrings []string
	for scanner.Scan() {
		problemStrings = append(problemStrings, scanner.Text())
	}
	return problemStrings[1:]
}

func main() {
	file, err := os.Open(csvFileName)
	defer file.Close()
	if err != nil {
		fmt.Errorf("got error %s", err)
	}
	problems := quiz.ConvertStringsToProblems(parseFile(file))
	game := quiz.NewQuiz(problems)
	cli := quiz.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayQuiz()
}
