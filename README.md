# GOPS

Generator Of Projects Structures is a CLI for generating new project based on predefined templates. You can define a template defining the structure of the project in a file _your-template-name.tmpl_

## Install

```bash
go get github.com/krafugo/gops
```

## Commands

```bash
gops <command> [arguments]

gops help       # show the help of the project
gops init       # create a new project structure
gops list       # list all of available templates
```

Some examples of use:

```bash
gops init store-app sample  # Generate a new project named "store-app" using the "sample" template and create a new repository
gops init store-app sample --norepo # It's the same as above but without a repository

gops list sample # Show the structure of the "sample" template
```
