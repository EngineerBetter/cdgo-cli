package dir

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
)

func Find(needle string, haystack string, maxDepth int) (result string, errOut error) {
	fi, err := os.Lstat(haystack)

	if err != nil {
		return result, errors.New(haystack + " did not exist")
	} else if !fi.IsDir() {
		return result, errors.New(haystack + " is not a directory")
	}

	result = walk(haystack, needle, maxDepth, 0)

	if result == "" {
		errOut = errors.New(needle + " not found in " + haystack)
	}

	return
}

func walk(path string, needle string, maxDepth int, currentDepth int) (result string) {
	names, _ := readDirNames(path)
	subdirs := make([]string, len(names))

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, _ := os.Lstat(filename)

		if fileInfo.IsDir() {
			if fileInfo.Name() == needle {
				return filename
			} else {
				subdirs = append(subdirs, filename)
			}
		}
	}

	if shouldDescend(maxDepth, currentDepth) {
		for _, subdir := range subdirs {
			result = walk(subdir, needle, maxDepth, currentDepth+1)
			if result != "" {
				return result
			}
		}
	}

	return
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

func shouldDescend(maxDepth, currentDepth int) bool {
	return maxDepth < 1 || currentDepth < maxDepth
}
