package gorefer

import (
	"fmt"
	"go/ast"
	"go/types"
)

type Parser struct {
	FuncId   FuncId
	StructId StructId
	InterfaceId InterfaceId
	VarId	VarId
	//Tree     path.Root
}
//this struct is result of parsing some file
type FileInfo struct {
	Path string
	Package string
	FunctionList []FunctionInfo
	StructList []StructInfo
	InterfaceList []InterfaceInfo
	VarList []VarInfo
	Imports []string
}

func NewParser() *Parser {
	return &Parser{
		FuncId: FuncId{Id: 0},
		StructId: StructId{Id: 0},
		InterfaceId: InterfaceId{Id: 0},
		VarId: VarId{Id: 0},
	}
}
//get interface information from interface type declaration
func (p *Parser) GetInterfaceInfo(spec ast.Spec, obj types.Object) *InterfaceInfo {
	//var list []InterfaceInfo
	var methods []string
	for _, m := range spec.(*ast.TypeSpec).Type.(*ast.InterfaceType).Methods.List {
		if len(m.Names) == 0 {
			continue
		}
		methods = append(methods, m.Names[0].Name)
	}
	info := &InterfaceInfo{
		Id: p.InterfaceId.AllocateId(),
		Name: spec.(*ast.TypeSpec).Name.Name,
		Methods: methods,
		Package: obj.Pkg().Name(),
	}
	return info
}

//get struct information from structure type declaration
func (p *Parser) GetStructDeclInfo(spec ast.Spec, obj types.Object) *StructInfo {
	//var list []StructInfo
	var member []string
	for _, m := range spec.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
		if len(m.Names) == 0 {
			continue
		}
		member = append(member, m.Names[0].Name)
	}
	info := &StructInfo{
		Id: p.StructId.AllocateId(),
		Name: spec.(*ast.TypeSpec).Name.Name,
		Member: member,
		Package: obj.Pkg().Name(),
	}
	return info
}

func (p *Parser) GetFunctionInfo(decl *ast.FuncDecl, obj types.Object) FunctionInfo {
	fmt.Println(decl.Name)
	typ := obj.(*types.Func)
	returnTyp := typ.Type().String()
	info := Func{
		Name:         decl.Name.Name,
		Receiver:     stringer(GetReceiver(decl)),
		ReceiverType: stringer(GetReceiverType(decl)),
		ReturnType:   returnTyp,
		Package:      typ.Pkg().Name(),
	}
	return FunctionInfo{
		Id:           p.FuncId.AllocateId(),
		FuncInfo: info,
		Called:       nil,
	}
}

func GetReceiver(f *ast.FuncDecl) interface{} {
	if f.Recv == nil { return nil }
	if len(f.Recv.List[0].Names) == 0 { return nil }
	return f.Recv.List[0].Names[0].Name
}

func GetReceiverType(f *ast.FuncDecl) interface{} {
	if f.Recv == nil { return nil }
	if len(f.Recv.List[0].Names) == 0 { return nil }
	switch f.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return f.Recv.List[0].Type.(*ast.Ident).Name
	case *ast.StarExpr:
		return f.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
	}
	return nil
}

func (p *Parser) GetVarDecl(spec *ast.ValueSpec, obj types.Object) []VarInfo {
	var lis []VarInfo
	var o string
	for _, v := range spec.Names {
		switch obj.(type) {
		case *types.Var:
			o = obj.(*types.Var).Type().String()
		case *types.Const:
			o = obj.(*types.Const).Type().String()
		}
		info := VarInfo{
			Id: p.VarId.AllocateId(),
			Name: v.Name,
			Type: o,
		}
		lis = append(lis, info)
	}
	return lis
}

func stringer(str interface{}) string {
	if str == nil { return "" }
	return str.(string)
}