package linker

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"github.com/spectrex02/gorefer/analyzer/detectDecl"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

//called function or method information
type Called struct {
	Name string
	ReturnType string
	Receiver string
	ReceiverType string
	Package string
}
//type for mapping between caller function and called function
type Call map[*ast.FuncDecl][]Called


type Linker struct {
	Pkg gorefer.PackageInfo
	CallList []Call
}

var Analyzer = &analysis.Analyzer{
	Name:             "linker",
	Doc:              "linker for package",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer, detectDecl.Analyzer},
	ResultType:       nil,
	FactTypes:        nil,
}



func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	pkgInfo := pass.ResultOf[detectDecl.Analyzer].(*gorefer.PackageInfo)
	call := make(Call)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	//parse AST and link function or method
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			{
				parseBody(n.Body, pass)
			}
		case *ast.FuncLit:
		}
	})
	return nil, nil
}
