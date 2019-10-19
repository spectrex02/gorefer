package findcall

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestFindCall(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "testfiles")
	analysistest.Run(t, testdata, Analyzer, "testfiles/check_statement")
}
