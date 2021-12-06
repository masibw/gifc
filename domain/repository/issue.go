package repository

import (
	"github.com/masibw/gifc/domain/entity"
)

type Issue interface {
	Create(issue *entity.Issue, git *entity.Git) (createdIssue *entity.Issue, err error)
}
