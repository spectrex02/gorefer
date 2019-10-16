package linker

import (
	"flag"
	"fmt"
	//"github.com/spectrex02/gorefer"
	"github.com/spectrex02/gorefer/analyzer/detectDecl"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

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
	//pkgInfo := pass.ResultOf[detectDecl.Analyzer].(*gorefer.PackageInfo)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	//parse AST and link function or method
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			parseBody(n.Body, pass)
		case *ast.FuncLit:
		}
	})
	return nil, nil
}

func parseBody(block *ast.BlockStmt, pass *analysis.Pass) {
	if len(block.List) == 0 { return }
	for _, stmt := range block.List {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			{
				parseAssign(s, pass)
			}
		case *ast.ExprStmt:
			parseExprStmt(s, pass)
		case *ast.RangeStmt:
		case *ast.ForStmt:
		case *ast.IfStmt:
		}
	}
}

//parse expr stmt -> *ast.CallExpr, *ast.SelectorExpr
func parseExprStmt(expr *ast.ExprStmt, pass *analysis.Pass) {
	switch f := expr.X.(type) {
	case *ast.CallExpr:
		{
			info := f.Fun
			obj := pass.TypesInfo.Types[f]
			objType := obj.Type.String()
			//objValue := obj.Value.String()
			fmt.Println("---------------------------------------")
			fmt.Printf("func info -> %v\nfunc obj -> %v\nfunc type -> %v\n", info, obj, objType)
			fmt.Println("---------------------------------------")
		}
	case *ast.SelectorExpr:
		{
			obj := pass.TypesInfo.Selections[f]
			fmt.Println("---------------------------------------")
			fmt.Printf("method name -> %v\nmethod type -> %v\nmethod receiver -> %v\n", obj.String(), obj.Type().String(), obj.Recv().String())
			fmt.Println("---------------------------------------")
		}
	}
}

//parse assign stmt
func parseAssign(a *ast.AssignStmt, pass *analysis.Pass) {
	for _, r := range a.Rhs {
		typ := pass.TypesInfo.Types[r].Type.String()
		//value := pass.TypesInfo.Types[r].Value.String()
		fmt.Printf("Assign stnt -> (type: %v)\n", typ)
	}
}

//parse range stmt
func parseRangeStmt(stmt *ast.RangeStmt, pass *analysis.Pass) {

}

//parse for stmt
func parseForStmt(stmt *ast.ForStmt, pass *analysis.Pass) {

}

//parse if stmt
func parseIfStmt(stmt *ast.IfStmt, pass *analysis.Pass) {

}

//parse go stmt
func parseGoStmt(stmt *ast.GoStmt, pass *analysis.Pass) {

}