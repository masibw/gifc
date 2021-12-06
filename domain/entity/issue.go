package entity

import "fmt"

type Issue struct {
	Number  int
	Title   string
	Content string
	URL     string
}

func NewIssue(title, content string) *Issue {
	return &Issue{
		Title:   title,
		Content: content,
	}
}

func (i *Issue) SetGithubInfo(nullableNumber *int, nullableURL *string) *Issue {
	var number int
	if nullableNumber == nil {
		number = 0
	} else {
		number = *nullableNumber
	}

	var url string
	if nullableURL == nil {
		url = ""
	} else {
		url = *nullableURL
	}

	return &Issue{
		Number:  number,
		Title:   i.Title,
		Content: i.Content,
		URL:     url,
	}
}

func (i *Issue) GenerateTodoCommentWithGithubInfo() string {
	return fmt.Sprintf("// TODO-#%d{%s}: %s", i.Number, i.URL, i.Content)
}
