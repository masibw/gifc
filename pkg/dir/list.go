package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ListFilePaths(start string) (filePaths []string, err error) {
	var dirEntries []os.DirEntry
	dirEntries, err = os.ReadDir(start)
	if err != nil {
		err = errors.Wrap(err, "failed to ListFilePaths")
		return
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			var files []string
			files, err = ListFilePaths(filepath.Join(start, dirEntry.Name()))
			if err != nil {
				err = errors.Wrap(err, "failed to ListFilePaths")
				return
			}
			filePaths = append(filePaths, files...)
			continue
		}
		filePaths = append(filePaths, filepath.Join(start, dirEntry.Name()))
	}

	return
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
