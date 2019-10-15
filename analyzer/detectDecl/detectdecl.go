package detectDecl

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:             "detectdecl",
	Doc:              "detect declarations of some object",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	ResultType:       nil,
	FactTypes:        nil,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
		(*ast.GenDecl)(nil),
	}

	parser := gorefer.NewParser()
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
		case *ast.GenDecl:
		}
	})
	return nil, nil
}