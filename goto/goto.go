package main

import . "github.com/EngineerBetter/goto/dir"
import "os"
import "path/filepath"
import "fmt"

func main() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		fmt.Println("GOPATH is not set")
		os.Exit(1)
	}

	haystack := filepath.Join(gopath, "src")

	if len(os.Args) < 2 {
		fmt.Println("directory to look for was not specified")
		os.Exit(1)
	}
	needle := os.Args[1]

	finder := new(RecursiveFinder)
	result, err := finder.Find(needle, haystack)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
