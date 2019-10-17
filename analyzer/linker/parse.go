package linker

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func parseBody(block *ast.BlockStmt, pass *analysis.Pass) {
	if len(block.List) == 0 { return }
	for _, stmt := range block.List {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			{
				parseAssignStmt(s, pass)
			}
		case *ast.ExprStmt:
			parseExprStmt(s, pass)
		case *ast.RangeStmt:
		case *ast.ForStmt:
		case *ast.IfStmt:
		}
	}
}

//parse call expr
func parseCallExpr(expr *ast.CallExpr, pass *analysis.Pass) Called {
	obj := pass.TypesInfo.Types[expr]
	typ := obj.Type.String()
	var name string

	switch c := expr.Fun.(type) {
	case *ast.Ident:
		name = c.Name
	case *ast.SelectorExpr:
	}

	called := Called{
		Name: name,
		ReturnType: typ,
		Receiver:

	}
}

//parse selector expr
func parseSelectorExpr(expr *ast.SelectorExpr, pass *analysis.Pass) Called {

}


//parse expr stmt -> *ast.CallExpr, *ast.SelectorExpr
func parseExprStmt(stmt *ast.ExprStmt, pass *analysis.Pass) {
	switch f := stmt.X.(type) {
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


//parse call expr

//parse assign stmt
func parseAssignStmt(a *ast.AssignStmt, pass *analysis.Pass) {
	for _, r := range a.Rhs {
		typ := pass.TypesInfo.Types[r].Type.String()
		switch s := r.(type) {
		case *ast.SelectorExpr:
			pass.TypesInfo.Selections[s].Recv().String()
		}
		//value := pass.TypesInfo.Types[r].Value.String()
		fmt.Printf("Assign stnt -> (type: %v)\n", typ)
	}
}

//parse range stmt
/*
type RangeStmt struct {
        For        token.Pos   // position of "for" keyword
        Key, Value Expr        // Key, Value may be nil
        TokPos     token.Pos   // position of Tok; invalid if Key == nil
        Tok        token.Token // ILLEGAL if Key == nil, ASSIGN, DEFINE
        X          Expr        // value to range over
        Body       *BlockStmt
}
 */
func parseRangeStmt(stmt *ast.RangeStmt, pass *analysis.Pass) {
	body := stmt.Body
	value := stmt.Value
	x := stmt.X
	valueType := pass.TypesInfo.Types[value].Type.String()
	xType := pass.TypesInfo.Types[x].Type.String()

	parseBody(body, pass)

}

//parse for stmt
/*
type ForStmt struct {
        For  token.Pos // position of "for" keyword
        Init Stmt      // initialization statement; or nil
        Cond Expr      // condition; or nil
        Post Stmt      // post iteration statement; or nil
        Body *BlockStmt
}

 */
func parseForStmt(stmt *ast.ForStmt, pass *analysis.Pass) {
	body := stmt.Body
	parseBody(body, pass)

}

//parse if stmt
/*
type IfStmt struct {
        If   token.Pos // position of "if" keyword
        Init Stmt      // initialization statement; or nil
        Cond Expr      // condition
        Body *BlockStmt
        Else Stmt // else branch; or nil
}
 */
func parseIfStmt(stmt *ast.IfStmt, pass *analysis.Pass) {
	init := stmt.Init
	switch init.(type) {
	case *ast.AssignStmt:
		parseAssignStmt(init.(*ast.AssignStmt), pass)
	}
	body := stmt.Body
	parseBody(body, pass)

	if stmt.Else != nil {
		parseIfStmt(stmt.Else.(*ast.IfStmt), pass)
	}
}

//parse go stmt
/*
type GoStmt struct {
	Go   token.Pos // position of "go" keyword
	Call *CallExpr
}
*/
func parseGoStmt(stmt *ast.GoStmt, pass *analysis.Pass) {
	call := stmt.Call
	info := pass.TypesInfo.Types[call].Type
	fmt.Printf("%v\n", info)
}

//parse return stmt
/*
type ReturnStmt struct {
    Return  token.Pos // position of "return" keyword
    Results []Expr    // result expressions; or nil
}
 */
func parseReturnStmt(stmt *ast.ReturnStmt, pass *analysis.Pass) {
	if len(stmt.Results) == 0 {
		return
	}
	for _, r := range stmt.Results {
		switch r := r.(type) {
		case *ast.CallExpr:
		case *ast.SelectorExpr:
		}
	}
}

//parse func lit
/*
type FuncLit struct {
        Type *FuncType  // function type
        Body *BlockStmt // function body
}
 */
func parseFuncLit(lit *ast.FuncLit, pass *analysis.Pass) {


}
