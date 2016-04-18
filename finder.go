package main

type DirectoryFinder interface {
	Find(directory string, in string) (string, error)
}

type RecursiveFinder struct{}

func (*RecursiveFinder) Find(needle string, haystack string) (string, error) {
	return "foo", nil
}
