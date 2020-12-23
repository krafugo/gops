package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Command represent a simple command in Bash
type Command struct {
	CMD    string
	Args   []string
	Dir    string
	Stdout bool
}

// NewC return a new Command struct
func NewC(cmd, dir string, stdout bool) Command {
	args := splitCMD(cmd)
	return Command{args[0], args[1:], dir, stdout}
}

// SplitCMD splits the string cmd around each instance of one or more
// consecutive white space but detect any valid substring between quotes
// "" and admit it as a whole single string
// Ex: cmd = `git commit -m "Initial commit"`` return
// ["git", "commit", "-m", "Initial commit"]
func splitCMD(cmd string) []string {
	fields := strings.Fields(cmd)
	if strings.Count(cmd, `"`)%2 != 0 {
		return fields
	}
	var result []string
	n := len(fields)
	for i := 0; i < n; i++ {
		f := fields[i]
		if strings.HasPrefix(f, `"`) {
			var aux string
			for j := i; j < n; j++ {
				if strings.HasSuffix(fields[j], `"`) {
					aux += fields[j]
					i = j
					break
				}
				aux += fields[j] + " "
			}
			result = append(result, aux)
			continue
		}
		result = append(result, f)
	}
	return result
}

// Execute a Command c on Terminal and return error if
// something was wrong
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

// Commands is a list of instances of Command struct
type Commands []Command

// ExecuteAll runs a sorted list of Commands
func (c Commands) ExecuteAll() error {
	for _, cmd := range c {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Execute the command n in a list of Commands
func (c Commands) Execute(n int) error {
	if n < 0 || n >= len(c) {
		return fmt.Errorf("n: %d is out of the range", n)
	}
	return c[n].Execute()
}
