package templates

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Template ...
type Template struct {
	Dirs  []string
	Files []string
	Root  string
}

// New return a Template
func New(filename string, root string) (Template, error) {
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
		line := scanner.Text()
		if strings.HasSuffix(line, "/") {
			line = line[:len(line)-1]
			dirs = append(dirs, line)
		} else {
			files = append(files, line)
		}
	}
	return Template{dirs, files, root}, nil
}

//Build ...
func (t Template) Build() error {
	// Make root

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
	return nil
}
