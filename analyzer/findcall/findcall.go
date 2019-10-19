package findcall

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"
)

//called function or method information



//type for mapping between caller function and called function



type Linker struct {
	Pkg gorefer.PackageInfo
	CallList []gorefer.Call
}

var Analyzer = &analysis.Analyzer{
	Name:             "findcall",
	Doc:              "findcall for package",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	ResultType:       reflect.TypeOf(*new(gorefer.Call)),
	FactTypes:        nil,
}



func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	call := make(gorefer.Call)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	//parse AST and link function or method
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			{
				called := parseBody(n.Body, pass)
				callerReturnType := pass.TypesInfo.Defs[n.Name].(*types.Func).Type().String()
				caller := gorefer.Func{
					Name:         n.Name.Name,
					ReturnType:   callerReturnType,
					Receiver:     resolveReceiverName(gorefer.GetReceiver(n)),
					ReceiverType: resolveReceiverName(gorefer.GetReceiverType(n)),
					Package:      pass.TypesInfo.Defs[n.Name].Pkg().Name(),
				}
				for _, c := range called {
					c.Show()
				}
				call[caller] = called
			}
		}
	})
	return call, nil
}

func resolveReceiverName(r interface{}) string {
	if r == nil { return "" }
	return r.(string)
}