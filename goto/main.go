package main

import (
	"flag"
	"fmt"
	"github.com/EngineerBetter/cdgo/dir"
	"github.com/EngineerBetter/cdgo/goto/installer"
	"os"
	"path/filepath"
)

func main() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		printAndExit("GOPATH is not set")
	}

	haystack := filepath.Join(gopath, "src")

	installFilePtr := flag.String("install", "", "path to install Bash functions to")
	flag.Parse()
	installTo := *installFilePtr

	if installTo != "" {
		err := installer.Install(installTo)
		if err != nil {
			printAndExit(err)
		}
		fmt.Println("Added Bash functions to " + installTo)
	} else {
		if len(os.Args) < 2 {
			printAndExit("directory to look for was not specified")
		}
		needle := os.Args[1]

		result, err := dir.Find(needle, haystack)

		if err != nil {
			printAndExit(err)
		}

		fmt.Println(result)
	}
}

func printAndExit(message interface{}) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
