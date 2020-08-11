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

const ext = ".tmpl"

// New return an instance of a Template
func New(root, tmplName string) (Template, error) {
	var dirs, files []string
	filename := os.Getenv("GOPS_SCHEMA") + tmplName + ext
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
		}
		dir, file := splitFilename(line)
		if len(dir) != 0 {
			dirs = append(dirs, dir)
		}
		if len(file) != 0 {
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
		fmt.Printf("> Created dir: %s\n", dir)
	}
	// Make files
	for _, file := range t.Files {
		_, filename := splitFilename(file)
		content := loadContent(t.Name, filename)
		content = replaceTokens(content, t.Root)
		file = filepath.Join(t.Root, file)
		err := ioutil.WriteFile(file, content, 0644)
		if err != nil {
			return err
		}
		fmt.Printf("> Created file: %s\n", file)
	}
	if norepo {
		return nil
	}
	return t.CreateRepo()
}

func (t Template) String() string {
	r := "Template: " + t.Name + "\n"
	if len(t.Name) != 0 {
		r += "Project: " + t.Root + "\n"
	}
	if len(t.Dirs) != 0 {
		r += "Dirs: \n"
		for _, dir := range t.Dirs {
			r += "\t" + dir + "\n"
		}
	}
	if len(t.Files) != 0 {
		r += "Files: \n"
		for _, file := range t.Files {
			r += "\t" + file + "\n"
		}
	}
	return r
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

// loadContent load all the content of an specific file
func loadContent(tmpl, filename string) []byte {
	contentPath := os.Getenv("GOPS_CONTENT")
	content, err := ioutil.ReadFile(contentPath + tmpl + "/" + filename)
	if err == nil {
		return content
	}
	communPath := os.Getenv("GOPS_COMMUN")
	content, err = ioutil.ReadFile(communPath + filename)
	if err == nil {
		return content
	}
	return []byte("")
}

func splitFilename(file string) (string, string) {
	if i := strings.LastIndex(file, "/"); i != -1 {
		return file[:i+1], file[i+1:] // dir, filename
	}
	return "", file
}

//! This should be a map of token-value
func replaceTokens(content []byte, projectName string) []byte {
	return []byte(strings.ReplaceAll(string(content), "{# projectName #}", projectName))
}
