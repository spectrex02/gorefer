package main

import (
	"github.com/spectrex02/gorefer/analyzer/linker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {

	/*
		progname := filepath.Base(os.Args[0]) <- get directory path
	*/
	singlechecker.Main(linker.Analyzer)
}
