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

	timer := cli.quiz.StartTimer()
	inputChan := make(chan string)

	for i := 0; i < cli.quiz.GetNumberOfQuestions(); i++ {
		fmt.Fprint(cli.out, cli.quiz.GetQuestion()+" ")
		go cli.readLine(inputChan)

		select {
		case <-timer.C:
			fmt.Fprint(cli.out, "\nTime ran out!\n")
			fmt.Fprint(cli.out, cli.quiz.GetResult())
			return
		case answer := <-inputChan:
			cli.quiz.SubmitAnswer(answer)
		}
	}

	fmt.Fprint(cli.out, cli.quiz.GetResult())
}

func (cli *CLI) readLine(inputChan chan string) {
	cli.in.Scan()
	inputChan <- strings.TrimSpace(cli.in.Text())
}
