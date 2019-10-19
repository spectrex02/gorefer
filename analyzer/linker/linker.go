package linker

import (
	"flag"
	"fmt"
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
	pkgInfo := pass.ResultOf[detectDecl.Analyzer].(*gorefer.PackageInfo)
	call := pass.ResultOf[findcall.Analyzer].(gorefer.Call)
	fmt.Println(pkgInfo.Name)
	Link(*pkgInfo, call)
	return nil, nil
}


//link called and caller
func Link(pkgInfo gorefer.PackageInfo, call gorefer.Call) {
	for _, f := range pkgInfo.Function {
		called := call[f.FuncInfo]
		if called == nil { continue }
		f.Called = called
		f.Show()
	}
}

