package linker

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestLinker(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "testfiles")
}