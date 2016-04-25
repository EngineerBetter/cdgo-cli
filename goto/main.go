package main

import (
	"flag"
	"fmt"
	"github.com/EngineerBetter/cdgo/bash"
	"github.com/EngineerBetter/cdgo/dir"
	"os"
	"path/filepath"
)

func main() {
	var installTo, needle, haystackType string
	flag.StringVar(&installTo, "install", "", "path to install Bash functions to")
	flag.StringVar(&needle, "needle", "", "directory to find")
	flag.StringVar(&haystackType, "haystackType", "go", "go/work")
	flag.Parse()

	if installTo != "" {
		doInstall(installTo)
	} else {
		if needle == "" {
			printAndExit("-needle must be provided")
		}
		doFind(needle, haystackType)
	}
}

func doInstall(installTo string) {
	err := bash.Install(installTo)
	if err != nil {
		printAndExit(err)
	}
	fmt.Println("Added Bash functions to " + installTo)
	fmt.Println("To load new functions: source " + installTo)
}

func doFind(needle string, haystackType string) {
	haystack, maxDepth := getHaystack(haystackType)
	result, err := dir.Find(needle, haystack, maxDepth)

	if err != nil {
		printAndExit(err)
	}

	fmt.Println(result)
}

func getHaystack(haystackType string) (haystack string, maxDepth int) {
	switch haystackType {
	case "go":
		haystack = getGoSrc()
		maxDepth = -1
	case "work":
		haystack = getWorkspace()
		maxDepth = 1
	default:
		printAndExit("-haystackType must be either go or work")
	}

	return
}

func getGoSrc() string {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		printAndExit("GOPATH is not set")
	}

	return filepath.Join(os.Getenv("GOPATH"), "src")
}

func getWorkspace() string {
	home := os.Getenv("HOME")

	if home == "" {
		printAndExit("HOME is not set. This tool depends on Bash, and only works on Windows when using MinGW or equivalent.")
	}

	return filepath.Join(os.Getenv("HOME"), "workspace")
}

func printAndExit(message interface{}) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
