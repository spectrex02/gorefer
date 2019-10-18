package linker

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"github.com/spectrex02/gorefer/analyzer/detectDecl"
	"github.com/spectrex02/gorefer/analyzer/findcall"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:             "linker",
	Doc:              "linker for the package",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: false,
	Requires:         []*analysis.Analyzer{detectDecl.Analyzer, findcall.Analyzer},
	ResultType:       nil,
	FactTypes:        nil,
}

func run(pass *analysis.Pass) (interface{}, error) {
	pkginfo := pass.ResultOf[detectDecl.Analyzer].(gorefer.PackageInfo)
	call := pass.ResultOf[findcall.Analyzer].(findcall.Call)

	return nil, nil
}