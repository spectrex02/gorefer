package linker

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func parseBody(block *ast.BlockStmt, pass *analysis.Pass) []Called {
	if len(block.List) == 0 { return nil }
	var called []Called
	for _, stmt := range block.List {
		switch s := stmt.(type) {
		case *ast.BlockStmt:
			called = append(called, parseBody(s, pass)...)
		case *ast.AssignStmt:
				called = append(called, parseAssignStmt(s, pass)...)
		case *ast.ExprStmt:
			called = append(called, parseExprStmt(s, pass))
		case *ast.RangeStmt:
			called = append(called, parseRangeStmt(s, pass)...)
		case *ast.ForStmt:
			called = append(called, parseForStmt(s, pass)...)
		case *ast.IfStmt:
			called = append(called, parseIfStmt(s, pass)...)
		case *ast.ReturnStmt:
			called = append(called, parseReturnStmt(s, pass)...)
		case *ast.GoStmt:
			called = append(called, parseGoStmt(s, pass)...)
		case *ast.SwitchStmt:
			called = append(called, parseSwitchStmt(s, pass)...)
		case *ast.TypeSwitchStmt:
			called = append(called, parseTypeSwitchStmt(s, pass)...)
		}
	}
	return called
}

//parse call expr
func parseCallExpr(expr *ast.CallExpr, pass *analysis.Pass) Called {
	obj := pass.TypesInfo.Types[expr]
	returnType := obj.Type.String()

	switch c := expr.Fun.(type) {
	case *ast.Ident:
		{
			return Called{
				Name:         c.Name,
				ReturnType:   returnType,
				Receiver:     "",
				ReceiverType: "",
				Package:      "",
			}
		}
	case *ast.SelectorExpr:
		{
			f := parseSelectorExpr(c, pass)
			return f
		}
	}

	return Called{
		Name: expr.Fun.(*ast.Ident).Name,
		ReturnType: returnType,
	}
}

//parse selector expr
func parseSelectorExpr(expr *ast.SelectorExpr, pass *analysis.Pass) Called {
	name := expr.Sel.Name
	returnType := pass.TypesInfo.Types[expr].Type.String()
	switch r := expr.X.(type) {
	case *ast.Ident:
		{
			receiver := r.Name
			var receiverType string
			var packageName string
			if pass.TypesInfo.Types[r].Type != nil {
				receiverType = pass.TypesInfo.Types[r].Type.String()
				packageName = ""
			} else {
				packageName = receiver
				receiverType = ""
			}
			return Called{
				Name:         name,
				ReturnType:   returnType,
				Receiver:     receiver,
				ReceiverType: receiverType,
				Package:      packageName,
			}
		}
	case *ast.SelectorExpr:
	//case *ast.StarExpr:
		{
			c :=  parseSelectorExpr(r, pass)
			receiver := c.Receiver
			receiverType := c.ReceiverType
			fmt.Printf("-----------> receiver :%v\n-----------> receiver type :%v\n", receiver, receiver)
			return Called{
				Name:name,
				ReturnType: returnType,
				Receiver: receiver,
				ReceiverType: receiverType,
				Package: c.Package,
			}
		}
	case *ast.CallExpr:
		{
			c := parseCallExpr(r, pass)
			return Called{
				Name:         name,
				ReturnType:   returnType,
				Receiver:     c.Name,
				ReceiverType: c.ReturnType,
				Package:      c.Package,
			}
		}
	}
	//NewBBB().hoge() <- reach here
	return Called{
		Name:         name,
		ReturnType:   returnType,
		Receiver:     "",
		ReceiverType: "",
		Package:      "hogehoge",
	}

}


//parse expr stmt -> *ast.CallExpr, *ast.SelectorExpr
func parseExprStmt(stmt *ast.ExprStmt, pass *analysis.Pass) Called {
	switch f := stmt.X.(type) {
	case *ast.CallExpr:
		{
			called := parseCallExpr(f, pass)
			//called.show()
			return called
		}
	case *ast.SelectorExpr:
		{
			called := parseSelectorExpr(f, pass)
			//called.show()
			return called
		}
	}
	return Called{}
}


//parse call expr

//parse assign stmt
func parseAssignStmt(a *ast.AssignStmt, pass *analysis.Pass) []Called {
	var called []Called
	for _, r := range a.Rhs {
		switch s := r.(type) {
		case *ast.SelectorExpr:
			{
				c := parseSelectorExpr(s, pass)
				called = append(called, c)
			}

		case *ast.CallExpr:
			{
				c := parseCallExpr(s, pass)
				called = append(called, c)
			}
		}
	}
	return called
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
func parseRangeStmt(stmt *ast.RangeStmt, pass *analysis.Pass) []Called {
	var called []Called
	x := stmt.X
	switch x := x.(type) {
	case *ast.CallExpr:
		called = append(called, parseCallExpr(x, pass))
	case *ast.SelectorExpr:
		called = append(called, parseSelectorExpr(x, pass))
	}
	called = append(called, parseBody(stmt.Body, pass)...)
	return called
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
func parseForStmt(stmt *ast.ForStmt, pass *analysis.Pass) []Called {
	return parseBody(stmt.Body, pass)

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
func parseIfStmt(stmt *ast.IfStmt, pass *analysis.Pass) []Called {
	var called []Called
	init := stmt.Init
	switch init.(type) {
	case *ast.AssignStmt:
		{
			c := parseAssignStmt(init.(*ast.AssignStmt), pass)
			called = append(called, c...)
		}
	}
	called = append(called, parseBody(stmt.Body, pass)...)

	if stmt.Else != nil {
		called = append(called, parseIfStmt(stmt.Else.(*ast.IfStmt), pass)...)
	}
	return called
}

//parse go stmt
/*
type GoStmt struct {
	Go   token.Pos // position of "go" keyword
	Call *CallExpr
}
*/
func parseGoStmt(stmt *ast.GoStmt, pass *analysis.Pass) []Called {
	var called []Called
	c := stmt.Call
	called = append(called, parseCallExpr(c, pass))
	return called
}


//parse return stmt
/*
type ReturnStmt struct {
    Return  token.Pos // position of "return" keyword
    Results []Expr    // result expressions; or nil
}
 */
func parseReturnStmt(stmt *ast.ReturnStmt, pass *analysis.Pass) []Called {
	var called []Called
	if len(stmt.Results) == 0 {
		return called
	}
	for _, r := range stmt.Results {
		switch r := r.(type) {
		case *ast.CallExpr:
			called = append(called, parseCallExpr(r, pass))
		case *ast.SelectorExpr:
			called = append(called, parseSelectorExpr(r, pass))
		}
	}
	return called
}

//parse switch stmt
/*
type SwitchStmt struct {
    Switch token.Pos  // position of "switch" keyword
    Init   Stmt       // initialization statement; or nil
    Tag    Expr       // tag expression; or nil
    Body   *BlockStmt // CaseClauses only
}
 */
func parseSwitchStmt(stmt *ast.SwitchStmt, pass *analysis.Pass) []Called {
	var called []Called
	init := stmt.Init
	switch init.(type) {
	case *ast.AssignStmt:
		{
			c := parseAssignStmt(init.(*ast.AssignStmt), pass)
			called = append(called, c...)
		}
	}
	body := stmt.Body
	if body == nil {
		return called
	}
	if len(body.List) == 0 {
		return called
	}
	for _, b := range body.List {

		for _, stmt := range b.(*ast.CaseClause).Body {
			switch s := stmt.(type) {
			case *ast.BlockStmt:
				called = append(called, parseBody(s, pass)...)
			case *ast.AssignStmt:
				called = append(called, parseAssignStmt(s, pass)...)
			case *ast.ExprStmt:
				called = append(called, parseExprStmt(s, pass))
			case *ast.RangeStmt:
				called = append(called, parseRangeStmt(s, pass)...)
			case *ast.ForStmt:
				called = append(called, parseForStmt(s, pass)...)
			case *ast.IfStmt:
				called = append(called, parseIfStmt(s, pass)...)
			case *ast.ReturnStmt:
				called = append(called, parseReturnStmt(s, pass)...)
			case *ast.GoStmt:
				called = append(called, parseGoStmt(s, pass)...)
			case *ast.SwitchStmt:
				called = append(called, parseSwitchStmt(s, pass)...)
			case *ast.TypeSwitchStmt:
				called = append(called, parseTypeSwitchStmt(s, pass)...)
			}
		}
	}
	return called
}

//parse type switch stmt
/*
type TypeSwitchStmt struct {
        Switch token.Pos  // position of "switch" keyword
        Init   Stmt       // initialization statement; or nil
        Assign Stmt       // x := y.(type) or y.(type)
        Body   *BlockStmt // CaseClauses only
}
 */
func parseTypeSwitchStmt(stmt *ast.TypeSwitchStmt, pass *analysis.Pass) []Called {
	var called []Called
	body := stmt.Body
	if body == nil {
		return called
	}
	if len(body.List) == 0 {
		return called
	}
	for _, b := range body.List {

		for _, stmt := range b.(*ast.CaseClause).Body {
			switch s := stmt.(type) {
			case *ast.BlockStmt:
				called = append(called, parseBody(s, pass)...)
			case *ast.AssignStmt:
				called = append(called, parseAssignStmt(s, pass)...)
			case *ast.ExprStmt:
				called = append(called, parseExprStmt(s, pass))
			case *ast.RangeStmt:
				called = append(called, parseRangeStmt(s, pass)...)
			case *ast.ForStmt:
				called = append(called, parseForStmt(s, pass)...)
			case *ast.IfStmt:
				called = append(called, parseIfStmt(s, pass)...)
			case *ast.ReturnStmt:
				called = append(called, parseReturnStmt(s, pass)...)
			case *ast.GoStmt:
				called = append(called, parseGoStmt(s, pass)...)
			case *ast.SwitchStmt:
				called = append(called, parseSwitchStmt(s, pass)...)
			case *ast.TypeSwitchStmt:
				called = append(called, parseTypeSwitchStmt(s, pass)...)
			}
		}
	}
	return called
}

//parse func lit
/*
type FuncLit struct {
        Type *FuncType  // function type
        Body *BlockStmt // function body
}
 */
func parseFuncLit(lit *ast.FuncLit, pass *analysis.Pass) []Called {
	var called []Called
	body := lit.Body
	//typ := lit.Type.Results.List[0].Type
	if len(body.List) == 0 { return called }

	called = append(called, parseBody(body, pass)...)
	return called
}
