package infrastructure

import (
	"github.com/go-git/go-git/v5"
	"github.com/masibw/gifc/domain/entity"
	"github.com/pkg/errors"
	"strings"
)

type Git struct {
	Repository *git.Repository
}

func NewGit(path string) (*Git, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		err = errors.Wrap(err, "failed to git PlainOpen")
		return nil, err
	}
	return &Git{
		Repository: repository,
	}, nil
}

func (g *Git) Get() (*entity.Git, error) {
	remotes, err := g.Repository.Remotes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Remotes")
	}
	if len(remotes) == 0 {
		return nil, errors.New("there is no remote repository")
	}
	if len(remotes[0].Config().URLs) == 0 {
		return nil, errors.New("there is no remote repository's urls")
	}

	url := remotes[0].Config().URLs[0]

	return &entity.Git{
		Owner:      extractOwner(url),
		Repository: extractRepoName(url),
	}, nil
}

func extractOwner(url string) string {
	return url[strings.Index(url, ":")+1 : strings.Index(url, "/")]
}

func extractRepoName(url string) string {
	return url[strings.Index(url, "/")+1 : strings.Index(url, ".git")]
}
