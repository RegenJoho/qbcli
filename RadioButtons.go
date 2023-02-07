package qbcli

import (
	"fmt"
	"strconv"
)

type Options struct {
	options  []string
	callback QuestionCallback
	name     string
}

func NewOptions(name string, back QuestionCallback, options ...string) *Options {
	return &Options{
		name:     name,
		callback: back,
		options:  options,
	}
}

func (r *Options) GetName() string {
	return r.name
}

func (r *Options) GetOptions() []string {
	return r.options
}
func (r *Options) Ask() error {
	for {
		for k, v := range r.options {
			fmt.Println(fmt.Sprint("(", (k+1), ")"), v)
		}
		var answer string = ""
		_, err := fmt.Scanln(&answer)
		if err != nil {
			return err
		}
		num, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		if num-1 >= len(r.options) || num < 1 {
			fmt.Println("Your answer is out of bounds. Please try again.")
		} else {
			r.callback(r.options[num-1])
			return nil
		}
	}
}
