package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Template keeps the structure of a project
type Template struct {
	Dirs  []string // directories to create
	Files []string // files to create
	Root  string   // project name
	Name  string   // template name
}

// New return an instance of a Template
func New(filename, root, tmplName string) (Template, error) {
	var dirs, files []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return Template{}, err
	}
	defer file.Close()

	// Use bufio scanner, the default Scan method is by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := fixLine(scanner.Text())
		if len(line) == 0 {
			continue
		} else if strings.HasSuffix(line, "/") {
			line = line[:len(line)-1]
			dirs = append(dirs, line)
		} else {
			files = append(files, line)
		}
	}
	return Template{dirs, files, root, tmplName}, nil
}

//Build the project generating all the files and directories
func (t Template) Build(norepo bool) error {
	// Make dirs
	for _, dir := range t.Dirs {
		dir = filepath.Join(t.Root, dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	// Make files
	for _, file := range t.Files {
		file = filepath.Join(t.Root, file)
		err := ioutil.WriteFile(file, []byte(""), 0644)
		if err != nil {
			return err
		}
	}
	if norepo {
		return nil
	}
	return t.CreateRepo()
}

// CreateRepo creates a new Git repository into a root of a project
func (t Template) CreateRepo() error {
	gitInit := NewC("git init", t.Root, true)
	gitAdd := NewC("git add .", t.Root, true)
	gitCommit := NewC(`git commit -q -m "Initial Commit"`, t.Root, true)
	commands := Commands{gitInit, gitAdd, gitCommit}
	return commands.ExecuteAll()
}

// fixLine erases comments inside a line and left and right white spaces
func fixLine(line string) string {
	// Removing comments
	i := strings.Index(line, "#")
	if i != -1 {
		line = line[:i]
	}
	if len(line) == 0 || line == "\n" {
		return ""
	}
	// Removing white spaces
	if strings.HasPrefix(line, " ") {
		line = strings.TrimLeft(line, " ")
	}
	line = strings.TrimRight(line, " ")
	return line
}
