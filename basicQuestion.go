package qbcli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/inancgumus/screen"
)

// A basic question. Asks the user something and then calls a
type Question struct {
	callback   QuestionCallback
	question   string
	name       string
	shallClear bool
}

func (q *Question) Ask() error {
	fmt.Println(q.question)
	buffer := bufio.NewReader(os.Stdin)
	answer, err := buffer.ReadString('\n')
	answer = strings.ReplaceAll(answer, "\r", "")
	answer = strings.ReplaceAll(answer, "\n", "")
	if err != nil {
		return err
	}
	q.callback(answer)
	if q.shallClear {
		screen.Clear()
		screen.MoveTopLeft()
	}
	return nil
}

func NewQuestion(name, question string, callback QuestionCallback) *Question {
	return &Question{question: question, callback: callback, name: name}
}

func (q *Question) Clear(shallClear bool) *Question {
	q.shallClear = shallClear
	return q
}

func (q *Question) GetOptions() []string {
	return []string{q.question}
}
func (q *Question) GetName() string {
	return q.name
}
