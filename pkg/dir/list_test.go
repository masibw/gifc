package dir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListFiles(t *testing.T) {
	testDir := "../../testdata/src/testdir"
	filePaths, err := ListFilePaths(testDir)
	if err != nil {
		t.Fatalf("failed to list files dir: %s, err: %v", testDir, err)
	}

	expected := []string{"../../testdata/src/testdir/a.go", "../../testdata/src/testdir/b.html", "../../testdata/src/testdir/recursive/c.xml"}
	assert.Equal(t, expected, filePaths)
}
