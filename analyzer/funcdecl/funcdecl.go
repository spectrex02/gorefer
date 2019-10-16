package funcdecl

import (
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:             "funcdecl",
	Doc:              "analyze function declaration",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: false,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	ResultType:       nil,
	FactTypes:        nil,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	var funcDecl []*types.Func
	var funcLit []*ast.FuncLit

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		//
		switch n := n.(type) {
		case *ast.FuncDecl:
			{
				o := pass.TypesInfo.Defs[n.Name].(*types.Selection)
				f := pass.TypesInfo.Defs[n.Name].(*types.Func)
				fmt.Printf("object type -> %v\n", o.Id())
				t := f.Id()
				a := f.Pkg().Name()
				//b := pass.TypesInfo.Implicits[n].String()
				c := f.FullName()
				d, e, g := types.LookupFieldOrMethod(o.Type(), true, o.Pkg(), o.Name())
				fmt.Printf("Look up -> %v, %v, %v\n", d, e, g)
				fmt.Printf("detected -> %v, Id -> %s, underlying -> %s\n", f, t, a)
				fmt.Printf("types --> %v\n", c)
				funcDecl = append(funcDecl, f)
			}
		case *ast.FuncLit:
			{
				funcLit = append(funcLit, n)
			}
		}
	})

	return nil, nil
}