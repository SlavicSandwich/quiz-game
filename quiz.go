package quiz

import (
	"fmt"
	"strings"
)

type Quiz struct {
	problems          []Problem
	NumberOfQuestions int
	currentIndex      int
	rightAnswersCount int
}

func NewQuiz(problems []Problem) Quiz {
	return Quiz{
		problems:          problems,
		NumberOfQuestions: len(problems),
		currentIndex:      0,
	}
}

func (quiz *Quiz) GetQuestion() string {
	return quiz.getCurrentProblem().question
}

func (quiz *Quiz) SubmitAnswer(answer string) {
	currentProblem := quiz.getCurrentProblem()
	expectedAnswer := currentProblem.answer
	if answer == expectedAnswer {
		quiz.rightAnswersCount++
	}
	quiz.currentIndex++
}

func (quiz *Quiz) getCurrentProblem() Problem {
	return quiz.problems[quiz.currentIndex]
}

func (quiz *Quiz) GetResult() string {
	return fmt.Sprintf("Got %v out of %v", quiz.rightAnswersCount, quiz.NumberOfQuestions)
}

type Problem struct {
	question string
	answer   string
}

func ConvertStringsToProblems(problemStrings []string) []Problem {
	var problems []Problem
	for _, str := range problemStrings {
		splitStr := strings.Split(str, ";")
		problems = append(problems, Problem{splitStr[0], splitStr[1]})
	}
	return problems
}
