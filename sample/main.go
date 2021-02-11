package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rwxrob/conf-go"
)

var name, path string

func init() {
	path, _ = os.Executable()
	name = filepath.Base(path)
}

func usage() {
	fmt.Printf(`usage:
  %[1]v (save|dump|print|file)
  %[1]v get <name>
  %[1]v set <name> <val>
  `, name,
	)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}
	config := conf.New()
	err := config.Load()
	if err != nil {
		fmt.Println("failed to load: %v\n", err)
		return
	}
	switch os.Args[1] {
	case "set":
		if len(os.Args) < 4 {
			usage()
		}
		config.Set(os.Args[2], os.Args[3])
		err := config.Save()
		if err != nil {
			fmt.Println(err)
		}
	case "get":
		if len(os.Args) < 3 {
			usage()
		}
		val := config.Get(os.Args[2])
		if val != "" {
			fmt.Println(val)
		}
	case "file":
		fmt.Println(config.Path())
	case "dump":
		fmt.Println(config)
	case "print":
		config.PrettyPrint()
	case "save":
		err = config.Save()
		if err != nil {
			fmt.Println(err)
		}
	}
}
