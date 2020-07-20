package main

import (
	"fmt"
	tmpl "gofio/templates"
)

func main() {
	template, err := tmpl.New("templates/sample.tmpl", "sample")
	if err != nil {
		fmt.Println("Error reading Template ", err)
	}
	err = template.Build()
	if err != nil {
		fmt.Println("Error creating Template ", err)
	}
	fmt.Println("DONE!")
	// os.Create()
	// os.Mkdir()
}
