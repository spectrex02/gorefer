package main

import (
	"github.com/spectrex02/gorefer/analyzer/detectDecl"
	"github.com/spectrex02/gorefer/analyzer/findcall"
	"github.com/spectrex02/gorefer/analyzer/linker"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {

	/*
		progname := filepath.Base(os.Args[0]) <- get directory path
	*/
	multichecker.Main(detectDecl.Analyzer, findcall.Analyzer)
	singlechecker.Main(linker.Analyzer)
}
