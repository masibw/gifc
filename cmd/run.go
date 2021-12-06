package cmd

import (
	"github.com/masibw/gifc/pkg/dir"
	"github.com/masibw/gifc/usecase"
	"log"
	"strings"
)

func Run(commentUC *usecase.CommentUseCase, startDir string) {
	filePaths, err := dir.ListFilePaths(startDir)
	if err != nil {
		log.Fatalf("can't list files :%+v", err)
		return
	}

	for _, filePath := range filePaths {
		if !strings.HasPrefix(filePath, startDir+"/.git") {
			err = commentUC.InspectFile(filePath)
			if err != nil {
				log.Printf("failed to InspectFile: filePath: %s, err: %v\n", filePath, err)
				continue
			}
		}
	}
}
