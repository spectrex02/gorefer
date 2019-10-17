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

}

//parse go stmt
/*
type GoStmt struct {
	Go   token.Pos // position of "go" keyword
	Call *CallExpr
}
*/
func parseGoStmt(stmt *ast.GoStmt, pass *analysis.Pass) {

}

//parse return stmt
/*
type ReturnStmt struct {
    Return  token.Pos // position of "return" keyword
    Results []Expr    // result expressions; or nil
}
 */
func parseReturnStmt(stmt *ast.ReturnStmt, pass *analysis.Pass) {

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
