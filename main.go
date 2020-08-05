/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
const TemplatePath = GopsRoot + "templates/"

// ContentPath is the path of the content of each file
const ContentPath = GopsRoot + "content/"

func init() {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("GOPATH environment variable is not set. " +
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.")
	}
	err := os.Setenv("GOPS_ROOT", gopath+GopsRoot)
	checkErr(err)
	err = os.Setenv("GOPS_TEMPLATES", gopath+TemplatePath)
	checkErr(err)
	err = os.Setenv("GOPS_CONTENT", gopath+ContentPath)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("GOPATH environment variable is not set. " +
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.")
	}
}
