package main

import . "github.com/EngineerBetter/goto/dir"
import "os"
import "os/user"
import "path/filepath"
import "fmt"
import "log"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("directory to look for was not specified")
	}
	needle := os.Args[1]

	finder := new(RecursiveFinder)
	haystack, err := getAbsoluteWorkspace()
	printAndBombIfOccurred(err)
	result, err := finder.Find(needle, haystack)
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
	log.Fatal(err)
}
