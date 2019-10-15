package gendecl

import (
	"github.com/spectrex02/gorefer"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:		"gendecl",
	Doc:		"detect struct or interface, custom type declaration",
	Run:		run,
	Requires:	[]*analysis.Analyzer{inspect.Analyzer},

}

func run(pass *analysis.Pass, parser *gorefer.Parser) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			{
				switch n.Tok {
				case token.TYPE:
					getGenDeclInfo(n, parser)
				}
			}
		}

	})

	return nil, nil
}

func getGenDeclInfo(n *ast.GenDecl, p *gorefer.Parser) {
	for _, spec := range n.Specs {
		switch spec.(*ast.TypeSpec).Type.(type) {
		case *ast.InterfaceType:
			p.GetInterfaceInfo(spec)
		case *ast.StructType:
			p.GetStructDeclInfo(spec)
		}
	}
}