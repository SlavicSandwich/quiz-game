package quiz

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	quiz Quiz
}

func NewCLI(in io.Reader, out io.Writer, quiz Quiz) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		quiz: quiz,
	}
}

func (cli *CLI) PlayQuiz() {
	fmt.Fprint(cli.out, "Answer all the questions right, if you can :)\n")

	for i := 0; i < cli.quiz.NumberOfQuestions; i++ {
		fmt.Fprint(cli.out, cli.quiz.GetQuestion()+" ")
		playerAnswer := cli.readLine()
		cli.quiz.SubmitAnswer(playerAnswer)
	}

	fmt.Fprint(cli.out, cli.quiz.GetResult())
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return strings.TrimSpace(cli.in.Text())
}
