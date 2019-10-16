package detectDecl

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"go/ast"
	"go/token"
	"reflect"

	//"go/types"
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
	ResultType:       reflect.TypeOf(new(gorefer.PackageInfo)),
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
	var structList []gorefer.StructInfo
	var interfaceList []gorefer.InterfaceInfo
	var varList []gorefer.VarInfo
	var functionList []gorefer.FunctionInfo
	//helper function
	var getGenDeclInfo func(*ast.GenDecl, *gorefer.Parser)
	getGenDeclInfo = func(n *ast.GenDecl, p *gorefer.Parser) {
		for _, spec := range n.Specs {
			switch spec.(*ast.TypeSpec).Type.(type) {
			case *ast.InterfaceType:
				{
					obj := pass.TypesInfo.Defs[spec.(*ast.TypeSpec).Name]
					info := p.GetInterfaceInfo(spec, obj)
					info.Show()
					interfaceList = append(interfaceList, *info)
				}
			case *ast.StructType:
				{
					obj := pass.TypesInfo.Defs[spec.(*ast.TypeSpec).Name]
					info := p.GetStructDeclInfo(spec, obj)
					info.Show()
					structList = append(structList, *info)
				}
			}
		}
	}

	var getFuncDeclInfo func(*ast.FuncDecl, *gorefer.Parser)
	getFuncDeclInfo = func(n *ast.FuncDecl, p *gorefer.Parser) {
		obj := pass.TypesInfo.Defs[n.Name]
		info := p.GetFunctionInfo(n, obj)
		info.Show()
		functionList = append(functionList, info)
	}

	var getVarDeclInfo func(*ast.GenDecl, *gorefer.Parser)
	getVarDeclInfo = func(n *ast.GenDecl, p *gorefer.Parser) {
		for _, spec := range n.Specs {
			obj := pass.TypesInfo.Defs[spec.(*ast.ValueSpec).Names[0]]
			switch vl := spec.(type) {
			case *ast.ValueSpec:
				{
					info := p.GetVarDecl(vl, obj)
					varList = append(varList, info...)
					for _, i := range info {
						i.Show()
					}

				}
			}
		}
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			{
				getFuncDeclInfo(n, parser)
			}
		case *ast.GenDecl:
			{
				switch n.Tok {
				case token.TYPE:
					getGenDeclInfo(n, parser)
				case token.VAR:
					getVarDeclInfo(n, parser)
				case token.CONST:
					getVarDeclInfo(n, parser)

				}
			}
		}
	})
	info := &gorefer.PackageInfo{
		Struct:    structList,
		Interface: interfaceList,
		Var:       varList,
		Function:  functionList,
	}
	return info, nil
}

