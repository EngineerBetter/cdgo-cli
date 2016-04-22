package main

import (
	"flag"
	"fmt"
	"github.com/EngineerBetter/cdgo/dir"
	"github.com/EngineerBetter/cdgo/goto/installer"
	"log"
	"os"
	"path/filepath"
)

func main() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		log.Fatal("GOPATH is not set")
	}

	haystack := filepath.Join(gopath, "src")

	installFilePtr := flag.String("install", "", "path to install Bash functions to")
	flag.Parse()
	installTo := *installFilePtr

	fmt.Println(installTo)

	if installTo != "" {
		installer.Install(installTo)
		fmt.Println("Added Bash functions to " + installTo)
	} else {
		if len(os.Args) < 2 {
			log.Fatal("directory to look for was not specified")
		}
		needle := os.Args[1]

		result, err := dir.Find(needle, haystack)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	}
}
