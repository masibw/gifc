package infrastructure

import (
	"context"
	"github.com/google/go-github/v41/github"
	"github.com/masibw/gifc/domain/entity"
	"github.com/masibw/gifc/domain/repository"
	"github.com/pkg/errors"
	"log"
)

type issue struct {
	client *github.Client
}

func NewIssue(client *github.Client) repository.Issue {
	return &issue{client: client}
}

func (i issue) Create(issue *entity.Issue, git *entity.Git) (createdIssue *entity.Issue, err error) {
	ctx := context.Background()
	issueRequest := &github.IssueRequest{
		Title:     &issue.Title,
		Body:      &issue.Content,
		Labels:    nil,
		Assignee:  nil,
		State:     nil,
		Milestone: nil,
		Assignees: nil,
	}

	var gIssue *github.Issue
	gIssue, _, err = i.client.Issues.Create(ctx, git.Owner, git.Repository, issueRequest)
	if err != nil {
		err = errors.Wrap(err, "failed to create issues")
		return
	}

	createdIssue = issue.SetGithubInfo(gIssue.Number, gIssue.HTMLURL)
	log.Printf("issue created. title: %s, url: %s \n", createdIssue.Title, createdIssue.URL)
	return
}
