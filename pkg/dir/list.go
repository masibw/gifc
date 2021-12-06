package dir

import (
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
