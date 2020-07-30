package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of available templates",
	Run: func(cmd *cobra.Command, args []string) {
		listTemplates()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// listTemplates shows all of templates available
func listTemplates() {
	ext := ".tmpl"
	tmplPath := os.Getenv("GOPS_TEMPLATES")
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
