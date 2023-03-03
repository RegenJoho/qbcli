package qbcli

import (
	"io"
	"os"
	"os/exec"
)

type Command struct {
	cmd        *exec.Cmd
	name       string
	shallClear bool
}

func NewCommand(name, cmmd string, args ...string) *Command {
	command := exec.Command(cmmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return &Command{
		name: name,
		cmd:  command,
	}
}

func (c *Command) Clear(shallClear bool) *Command {
	c.shallClear = shallClear
	return c
}

func NewCommandCustom(name string, stin io.Reader, stout, sterr io.Writer, cmmd string, args ...string) *Command {
	command := exec.Command(cmmd, args...)
	command.Stdin = stin
	command.Stdout = stout
	command.Stderr = sterr
	return &Command{
		name: name,
		cmd:  command,
	}
}

func (c *Command) GetOptions() []string {
	return []string{}
}
func (c *Command) GetName() string {
	return c.name
}

func (c *Command) Ask() error {
	return c.cmd.Run()
}
