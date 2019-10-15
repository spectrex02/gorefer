package gorefer

import (
	"go/ast"
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
func (p *Parser) GetInterfaceInfo(spec ast.Spec, pkg string) *InterfaceInfo {
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
		Package: pkg,
	}
	return info
}

//get struct information from structure type declaration
func (p *Parser) GetStructDeclInfo(spec ast.Spec, pkg string) *StructInfo {
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
		Package: pkg,
	}
	return info
}