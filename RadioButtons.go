package qbcli

import (
	"fmt"
	"strconv"
)

type Options struct {
	options    []string
	Callback   QuestionCallback
	name       string
	shallClear bool
}

func NewOptions(name string, back QuestionCallback, options ...string) *Options {
	return &Options{
		name:     name,
		Callback: back,
		options:  options,
	}
}

func (r *Options) Clear(shallClear bool) *Options {
	r.shallClear = shallClear
	return r
}

func (r *Options) GetName() string {
	return r.name
}

func (r *Options) AddOption(option string) *Options {
	r.options = append(r.options, option)
	return r
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
			r.Callback(num - 1)
			return nil
		}
	}
}
