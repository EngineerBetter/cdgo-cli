package finder

type NestedDirectoryFinder interface {
	find(directory string, in string) string
}

type Thing struct{}

func (*Thing) find(directory string, in string) string {
	return "foo"
}
