package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show gops help",
	Run: func(cmd *cobra.Command, args []string) {
		help()
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}

func help() {
	help := `
GOPS is a tool for generating structures of projects

Usage:
    gops <command> [arguments]

Available commands:

    help        show the help of the project
    init       	create a new project structure
    list       	list all of available templates

Use "gops help <command>" for more information about a command.`
	fmt.Println(help)
}
