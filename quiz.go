package quiz

import (
	"fmt"
	"strings"
	"time"
)

type Quiz struct {
	problems          []Problem
	currentIndex      int
	rightAnswersCount int
	timeLimit         time.Duration
}

func NewQuiz(problems []Problem, timeLimit time.Duration) Quiz {
	return Quiz{
		problems:     problems,
		currentIndex: 0,
		timeLimit:    timeLimit,
	}
}

func (quiz *Quiz) GetNumberOfQuestions() int {
	return len(quiz.problems)
}

func (quiz *Quiz) StartTimer() *time.Timer {
	return time.NewTimer(quiz.timeLimit)
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
	return fmt.Sprintf("Got %v out of %v", quiz.rightAnswersCount, quiz.GetNumberOfQuestions())
}

type Problem struct {
	question string
	answer   string
}

func ConvertStringsToProblems(problemStrings []string) []Problem {
	var problems []Problem
	for _, str := range problemStrings {
		splitStr := strings.Split(str, ";")
		problems = append(problems, Problem{splitStr[0], strings.TrimSpace(splitStr[1])})
	}
	return problems
}
