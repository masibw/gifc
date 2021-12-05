package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v41/github"
	"github.com/masibw/gifc/cmd"
	"github.com/masibw/gifc/domain/entity"
	"github.com/masibw/gifc/infrastructure"
	"github.com/masibw/gifc/usecase"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalf("you must provide github access token.")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get working directory: %v", err)
	}

	client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
	issueRepository := infrastructure.NewIssue(client)
	gitRepository, err := infrastructure.NewGit(".git")
	if err != nil {
		log.Fatalf("can't initialize git repository: %v", err)
	}
	var git *entity.Git
	git, err = gitRepository.Get()
	if err != nil {
		log.Fatalf("failed to get git information: %v", err)
	}
	commentUC := usecase.NewCommentUseCase(issueRepository, git)
	cmd.Run(commentUC, wd)

	fmt.Println("Finish!")
}
