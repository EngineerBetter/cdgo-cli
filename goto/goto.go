package main

import . "github.com/EngineerBetter/goto/dir"
import "os"
import "path/filepath"
import "fmt"
import "log"

func main() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		log.Fatal("GOPATH is not set")
	}

	haystack := filepath.Join(gopath, "src")

	if len(os.Args) < 2 {
		log.Fatal("directory to look for was not specified")
	}
	needle := os.Args[1]

	finder := new(RecursiveFinder)
	result, err := finder.Find(needle, haystack)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
