package main

import (
	"fmt"
	"github.com/EngineerBetter/goto/dir"
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
