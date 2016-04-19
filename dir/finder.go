package dir

import "errors"
import "io"
import "os"
import "path/filepath"
import "sort"

type DirectoryFinder interface {
	Find(directory string, in string) (string, error)
}

type RecursiveFinder struct{}

func (*RecursiveFinder) Find(needle string, haystack string) (result string, errOut error) {
	fi, err := os.Lstat(haystack)

	if err != nil {
		return result, errors.New(haystack + " did not exist")
	} else if !fi.IsDir() {
		return result, errors.New(haystack + " is not a directory")
	}

	filepath.Walk(haystack, func(path string, fi os.FileInfo, errIn error) (errOut error) {
		if fi.Name() == needle {
			result = path
			return io.EOF
		}

		return
	})

	if result == "" {
		errOut = errors.New(needle + " not found in " + haystack)
	}

	return
}

func (finder *RecursiveFinder) startWalk(root string, walkFn filepath.WalkFunc) error {
	info, _ := os.Lstat(root)
	return finder.walk(root, info, walkFn)
}

func (finder *RecursiveFinder) walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	err := walkFn(path, info, nil)
	if err != nil {
		if info.IsDir() && err == filepath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}

	names, err := readDirNames(path)
	if err != nil {
		return walkFn(path, info, err)
	}

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil && err != filepath.SkipDir {
				return err
			}
		} else {
			err = finder.walk(filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != filepath.SkipDir {
					return err
				}
			}
		}
	}
	return nil
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}
