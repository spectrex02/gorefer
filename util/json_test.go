package util

import (
	"github.com/spectrex02/gorefer"
	"github.com/spectrex02/gorefer/analyzer/linker"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestJson(t *testing.T) {
	testdata := analysistest.TestData()
	r := analysistest.Run(t, testdata, linker.Analyzer, "testfiles")
	info := r[0].Result.(*gorefer.PackageInfo)
	j := New(*info)
	data := j.ToJson()
	WriteJsonFile(j.Name, data)
}