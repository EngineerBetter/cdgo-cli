package main

import (
	"fmt"
	"github.com/EngineerBetter/goto/dir"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("directory to look for was not specified")
	}
	needle := os.Args[1]

	haystack, err := getAbsoluteWorkspace()
	printAndBombIfOccurred(err)
	result, err := dir.Find(needle, haystack)
	printAndBombIfOccurred(err)

	fmt.Println(result)
}

func getAbsoluteWorkspace() (dir string, err error) {
	usr, err := user.Current()
	printAndBombIfOccurred(err)
	dir = filepath.Join(usr.HomeDir, "workspace")
	_, err = os.Lstat(dir)
	return
}

func printAndBombIfOccurred(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
