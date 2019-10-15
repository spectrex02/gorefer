package funcdecl

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

//func init() {
//	Analyzer.Flags.Set()
//}

func TestFuncDecl(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "func_decl_test_file")
}