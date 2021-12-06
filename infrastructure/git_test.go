package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractOwner(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can extract owner name from github url",
			args: args{
				url: "git@github.com:masibw/gifc.git",
			},
			want: "masibw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractOwner(tt.args.url); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_extractRepoName(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can extract repository name from github url",
			args: args{
				url: "git@github.com:masibw/gifc.git",
			},
			want: "gifc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractRepoName(tt.args.url); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
