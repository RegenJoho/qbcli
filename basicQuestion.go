package qbcli

import "fmt"

// A basic question. Asks the user something and then calls a
type Question struct {
	callback QuestionCallback
	question string
	name     string
}

func (q *Question) Ask() error {
	fmt.Println(q.question)
	var answer string = ""
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return err
	}
	q.callback(answer)
	return nil
}

func NewQuestion(name, question string, callback QuestionCallback) *Question {
	return &Question{question: question, callback: callback, name: name}
}

func (q *Question) GetOptions() []string {
	return []string{q.question}
}
func (q *Question) GetName() string {
	return q.name
}
