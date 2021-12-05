package usecase

import (
	"bufio"
	"bytes"
	"github.com/masibw/gifc/domain/entity"
	"github.com/masibw/gifc/domain/repository"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type CommentUseCase struct {
	issueRepository repository.Issue
	git             *entity.Git
}

func NewCommentUseCase(issueRepository repository.Issue, git *entity.Git) *CommentUseCase {
	return &CommentUseCase{
		issueRepository: issueRepository,
		git:             git,
	}
}

func (c *CommentUseCase) InspectFile(filePath string) (err error) {
	var file *os.File
	file, err = os.Open(filePath)
	defer file.Close()
	if err != nil {
		err = errors.Wrap(err, "failed to open file")
		return
	}

	var bs []byte
	buf := bytes.NewBuffer(bs)
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if isTodoComment(line) && notCreated(extractCommentContent(line)) {
			commentContent := extractCommentContent(line)
			todoContent := extractTodoContent(commentContent)
			var createdIssue *entity.Issue
			createdIssue, err = c.issueRepository.Create(entity.NewIssue(todoContent, todoContent), c.git)
			if err != nil {
				err = errors.Wrap(err, "failed to create issue")
				return
			}

			newComment := createdIssue.GenerateTodoCommentWithGithubInfo()
			err = writeLine(buf, newComment)
			if err != nil {
				err = errors.Wrap(err, "failed to writeLine")
				return
			}
		} else {
			err = writeLine(buf, line)
			if err != nil {
				err = errors.Wrap(err, "failed to writeLine")
				return
			}
		}
	}

	if err = fileScanner.Err(); err != nil {
		err = errors.Wrap(err, "error while reading file")
		return
	}

	err = os.WriteFile(filePath, buf.Bytes(), 0666)
	if err != nil {
		err = errors.Wrap(err, "failed to write file")
		return
	}

	return
}

func isTodoComment(line string) bool {

	if !isComment(line) {
		return false
	}

	return isTodo(extractCommentContent(line))
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}

func extractCommentContent(line string) string {
	commentContent := strings.TrimPrefix(strings.TrimSpace(line), "//")
	return strings.TrimSpace(commentContent)
}

// commentContent means that there are no comment prefix('//') in front of it.
func isTodo(commentContent string) bool {
	upperLine := strings.ToUpper(strings.TrimSpace(commentContent))
	return strings.HasPrefix(upperLine, "TODO")
}

func notCreated(commentContent string) bool {
	return !strings.HasPrefix(strings.TrimSpace(commentContent), "TODO-#")
}

func extractTodoContent(commentContent string) string {
	commentContent = strings.TrimSpace(commentContent)

	// remove TODO(todo)
	noTodo := strings.TrimSpace(commentContent[4:])

	noTodo = strings.Trim(noTodo, ":")
	return strings.TrimSpace(noTodo)
}

func writeLine(buf *bytes.Buffer, content string) (err error) {
	_, err = buf.WriteString(content)
	if err != nil {
		err = errors.Wrap(err, "failed to WriteString")
		return
	}
	_, err = buf.WriteString("\n")
	if err != nil {
		err = errors.Wrap(err, "failed to WriteString")
		return
	}
	return
}
