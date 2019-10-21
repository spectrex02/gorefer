package linker

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestLinker(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "testfiles")
}

func TestChangeStructVal(t *testing.T) {
	type A struct {
		name string
	}
	a := A{name: "test"}
	fmt.Println(a.name)

	a.name = "hogehoge"

	fmt.Println(a)
}