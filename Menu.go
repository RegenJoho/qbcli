package qbcli

import (
	"fmt"
	"strconv"
)

type Menu struct {
	choices             []QuestionHandler
	name, question      string
	shallExitImmediatly bool
}

func NewMenu(name, question string, exitImmediatly bool) *Menu {
	return &Menu{
		choices:             []QuestionHandler{},
		name:                name,
		question:            question,
		shallExitImmediatly: exitImmediatly,
	}
}
func NewMenuWithoutQuestion(name string, exitImmediatly bool) *Menu {
	return NewMenu(name, "", exitImmediatly)
}
func NewMenuWithOptions(name, question string, exitImmediatly bool, handlers ...QuestionHandler) *Menu {
	return &Menu{
		name:                name,
		question:            question,
		shallExitImmediatly: exitImmediatly,
		choices:             handlers,
	}
}
func (c *Menu) AddMenu(choice QuestionHandler) *Menu {
	c.choices = append(c.choices, choice)
	return c
}

func (c *Menu) GetName() string {
	return c.name
}

func (c *Menu) GetOptions() []string {
	var options []string = make([]string, 0)
	for _, v := range c.choices {
		options = append(options, v.GetName())
	}
	return options
}

func (c *Menu) Ask() error {
	options := c.GetOptions()
	shallAllBreak := false
	for {
		var num int
		for {
			var lastNum int = 0
			if c.question != "" {
				fmt.Println(c.question)
			}
			for k, v := range options {
				fmt.Println(fmt.Sprint("(", (k+1), ")"), v)
				lastNum = k + 2
			}
			if !c.shallExitImmediatly {
				fmt.Println(fmt.Sprint("(", lastNum, ")"), "exit")
			}
			var answer string = ""
			_, err := fmt.Scanln(&answer)
			if err != nil {
				return err
			}
			num, err = strconv.Atoi(answer)
			if err != nil {
				return err
			}
			if !c.shallExitImmediatly {
				if num == lastNum {
					shallAllBreak = true
					break
				}
			}
			if num-1 >= len(options) || num < 1 {
				fmt.Println("Your answer is out of bounds. Please try again.")
			} else {
				break
			}
		}
		if shallAllBreak == true {
			break
		}
		c.choices[num-1].Ask()
		if c.shallExitImmediatly == true {
			break
		}
	}
	return nil
}
