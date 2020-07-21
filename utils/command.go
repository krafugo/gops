package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// Command ...
type Command struct {
	CMD    string
	Args   []string
	Dir    string
	Stdout bool
}

// Execute ...
func (c Command) Execute() error {
	cmd := exec.Command(c.CMD, c.Args...)
	cmd.Dir = c.Dir + "/"
	if c.Stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	return err
}

// Commands ...
type Commands []Command

// ExecuteAll ...
func (c Commands) ExecuteAll() error {
	for _, cmd := range c {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Execute ...
func (c Commands) Execute(n int) error {
	if n < 0 || n >= len(c) {
		return fmt.Errorf("n: %d is out of the range", n)
	}
	return c[n].Execute()
}
