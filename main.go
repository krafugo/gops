package main

import (
	"log"
	"os"

	"github.com/krafugo/gops/cmd"
)

func main() {
	cmd.Execute()
}

// GopsRoot is the path of GOPS path
const GopsRoot = "/src/github.com/krafugo/gops/"

// TemplatePath is the path of templates from $GOPATH
var TemplatePath = GopsRoot + "templates/"

// SchemaPath is the path of all of templates schemas
var SchemaPath = TemplatePath + "schema/"

// ContentPath is the path of the content of each file
var ContentPath = TemplatePath + "content/"

// CommunPath is the path of the all commun elements of templates
var CommunPath = TemplatePath + "commun/"

func init() {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("GOPATH environment variable is not set. " +
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.")
	}
	err := os.Setenv("GOPS_ROOT", gopath+GopsRoot)
	checkErr(err)

	tempPath := os.Getenv("GOPS_TEMPLATES")
	if len(tempPath) != 0 {
		TemplatePath = tempPath
	}
	err = os.Setenv("GOPS_TEMPLATES", gopath+TemplatePath)
	checkErr(err)
	err = os.Setenv("GOPS_SCHEMA", gopath+SchemaPath)
	checkErr(err)
	err = os.Setenv("GOPS_CONTENT", gopath+ContentPath)
	checkErr(err)
	err = os.Setenv("GOPS_COMMUN", gopath+CommunPath)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("GOPATH environment variable is not set. " +
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.")
	}
}
