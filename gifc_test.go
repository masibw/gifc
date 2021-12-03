package gifc_test

import (
	"testing"

	"github.com/masibw/gifc"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gifc.Analyzer, "a")
}

