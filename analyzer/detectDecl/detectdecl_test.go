package detectDecl

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestDetectDecl(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "testfiles")
}