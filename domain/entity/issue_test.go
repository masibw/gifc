package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssue_generateTodoCommentWithGithubInfo(t *testing.T) {
	type fields struct {
		Number  int
		Title   string
		Content string
		URL     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Generate new TODO comment with GithubInfo(URL).",
			fields: fields{
				Number:  1,
				Title:   "we must implement this.",
				Content: "we must implement this.",
				URL:     "https://github.com/masibw/gifc/issues/1",
			},
			want: "// TODO-#1{https://github.com/masibw/gifc/issues/1}: we must implement this.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Issue{
				Number:  tt.fields.Number,
				Title:   tt.fields.Title,
				Content: tt.fields.Content,
				URL:     tt.fields.URL,
			}
			if got := i.GenerateTodoCommentWithGithubInfo(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
