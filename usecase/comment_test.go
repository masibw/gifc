package usecase

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractCommentContent(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can remove '//' from comment.",
			args: args{
				line: "// comment.",
			},
			want: "comment.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractCommentContent(tt.args.line); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_extractTodoContent(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can extract TODO content from 'TODO'.",
			args: args{
				line: "TODO content.",
			},
			want: "content.",
		},
		{
			name: "Can extract TODO content from 'TODO:'.",
			args: args{
				line: "TODO: content.",
			},
			want: "content.",
		},
		{
			name: "Can extract todo content from 'todo'.",
			args: args{
				line: "todo content.",
			},
			want: "content.",
		},
		{
			name: "Can extract todo content from 'todo:'.",
			args: args{
				line: "todo: content.",
			},
			want: "content.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractTodoContent(tt.args.line); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_isComment(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Lines starting with '//' are comments.",
			args: args{
				line: "// This is comment.",
			},
			want: true,
		},
		{
			name: "Lines starting with ' //' are comments.",
			args: args{
				line: " // This is comment.",
			},
			want: true,
		},
		{
			name: "Lines not starting with '//' are not comments.",
			args: args{
				line: "This is not comment.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isComment(tt.args.line); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_isTodo(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Starting from 'TODO' is a TODO.",
			args: args{
				line: "TODO this is todo.",
			},
			want: true,
		},
		{
			name: "Starting from ' TODO' is a TODO.",
			args: args{
				line: " TODO this is todo.",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTodo(tt.args.line); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_isTodoComment(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "'//TODO:' is a todo comment.",
			args: args{
				line: "//TODO:",
			},
			want: true,
		},
		{
			name: "'// TODO:' is a todo comment.",
			args: args{
				line: "// TODO:",
			},
			want: true,
		},
		{
			name: "'// TODO :' is a todo comment.",
			args: args{
				line: "//TODO :",
			},
			want: true,
		},
		{
			name: "' // TODO :' is a todo comment.",
			args: args{
				line: " //TODO :",
			},
			want: true,
		},
		{
			name: "'// build:' is a todo comment.",
			args: args{
				line: "// build:",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTodoComment(tt.args.line); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_writeLine(t *testing.T) {
	type args struct {
		buf     *bytes.Buffer
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write with line feed code.",
			args: args{
				buf:     bytes.NewBuffer(make([]byte, 0)),
				content: "content",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeLine(tt.args.buf, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("writeLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.args.content+"\n", tt.args.buf.String())
		})
	}
}

func Test_notCreated(t *testing.T) {
	type args struct {
		commentContent string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TODO comment with issue number has already created.",
			args: args{
				commentContent: "// TODO-#1{https://github.com/masibw/gifc/issues/1}: we must implement this.",
			},
			want: false,
		},
		{
			name: "TODO comment with issue number has not created yet.",
			args: args{
				commentContent: "// TODO: we must implement this.",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notCreated(extractCommentContent(tt.args.commentContent)); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
