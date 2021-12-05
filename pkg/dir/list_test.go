package dir

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestListFiles(t *testing.T) {
	testDir := "../../testdata/src/testdir"
	filePaths, err := ListFilePaths(testDir)
	if err != nil {
		t.Fatalf("failed to list files dir: %s, err: %v", testDir, err)
	}
	expected := []string{filepath.Join(testDir, "/a.go"), filepath.Join(testDir, "/b.html"), filepath.Join(testDir, "recursive", "c.xml")}
	assert.Equal(t, expected, filePaths)
}
