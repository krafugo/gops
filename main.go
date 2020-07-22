package main

import (
	"fmt"
	"gofio/utils"
)

func main() {
	template, err := utils.New("templates/sample.tmpl", "sample")
	if err != nil {
		fmt.Println("Error reading Template ", err)
	}
	err = template.Build()
	if err != nil {
		fmt.Println("Error creating Template ", err)
	}
	fmt.Println("DONE!")
}
