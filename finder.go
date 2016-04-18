package main

type NestedDirectoryFinder interface {
	find(directory string, in string) string
}

type Thing struct{}

func (*Thing) Find(directory string, in string) string {
	return "foo"
}
