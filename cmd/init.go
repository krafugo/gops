package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/krafugo/gops/utils"
)

const ext = ".tmpl"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start a new project",
	Long: `This command allows to create the whole structure of a new project using some defined template.
	Usage: 
		$ gops init <projectName> [template]
	
	Example: 
		$ gops init new-api sample
		The above command create a new project named "new-api" and it generate all the files and directories from sample.tmpl file template. By default it initialize a git repository but you can omit it using the --norepo flag`,

	Run: func(cmd *cobra.Command, args []string) {
		norepo, _ := cmd.Flags().GetBool("norepo")
		initProject(args[0], args[1], norepo) // Project name, template name
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().Bool("norepo", false, "Avoid creating Git repository")
}

func initProject(name, tmpl string, norepo bool) {
	tmplPath := os.Getenv("GOPS_TEMPLATES") + tmpl + ext
	template, err := utils.New(tmplPath, name, tmpl)
	if err != nil {
		fmt.Println("Error reading Template ", err)
	}
	err = template.Build(norepo)
	if err != nil {
		fmt.Println("Error creating Template ", err)
	}
	fmt.Println("DONE!")
}
