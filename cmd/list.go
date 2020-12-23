package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/krafugo/gops/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of available templates",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			template, err := utils.New("", args[0])
			if err != nil {
				fmt.Println("Error reading template: ", err)
			} else {
				fmt.Println(template)
			}
		} else {
			listTemplates()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// listTemplates shows all of templates available
func listTemplates() {
	ext := ".tmpl"
	tmplPath := os.Getenv("GOPS_SCHEMA")
	// os.FileInfo
	files, err := ioutil.ReadDir(tmplPath)
	if err != nil {
		fmt.Println("Error listing templates: ", err)
	}
	fmt.Printf("————[ List of available Templates ]————\n\n")
	for _, file := range files {
		f := file.Name()
		if strings.HasSuffix(f, ext) {
			fmt.Println("\t      + " + strings.Replace(f, ext, "", 1))
		}
	}
	fmt.Printf("\n> You can choose anyone of above templates!\n")
}
