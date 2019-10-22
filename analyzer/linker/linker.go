package linker

import (
	"flag"
	"github.com/spectrex02/gorefer"
	"github.com/spectrex02/gorefer/analyzer/detectDecl"
	"github.com/spectrex02/gorefer/analyzer/findcall"
	"github.com/spectrex02/gorefer/util"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:             "linker",
	Doc:              "linker for the package",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: false,
	Requires:         []*analysis.Analyzer{detectDecl.Analyzer, findcall.Analyzer},
	//ResultType:       reflect.TypeOf(new(gorefer.PackageInfo)),
	FactTypes:        nil,
}

func run(pass *analysis.Pass) (interface{}, error) {
	pkgInfo := pass.ResultOf[detectDecl.Analyzer].(*gorefer.PackageInfo)
	call := pass.ResultOf[findcall.Analyzer].(gorefer.Call)
	//fmt.Println(pkgInfo.Name)
	result := Link(pkgInfo, call)
	for _, f := range result.Function {
		f.Show()
	}
	resultJson := util.New(*result)
	resultJson.OutputResult()
	rel := gorefer.ResolveFuncRelationship(result.Function)
	graph := util.NewGraph(result.Function, rel)
	util.OutputDot(graph, result.Name)
	//api.Serve()
	return nil, nil
}


//link called and caller
func Link(pkgInfo *gorefer.PackageInfo, call gorefer.Call) *gorefer.PackageInfo {
	var newFunctionList []gorefer.FunctionInfo
	for _, f := range pkgInfo.Function {
		called := call[f.FuncInfo]
		if called == nil { continue }
		f.Call = called
		//f.Show()
		newFucntionInfo := gorefer.FunctionInfo{
			Id: f.Id,
			FuncInfo: f.FuncInfo,
			Call: called,
		}
		newFunctionList = append(newFunctionList, newFucntionInfo)
	}

	return &gorefer.PackageInfo{
		Name: pkgInfo.Name,
		Struct: pkgInfo.Struct,
		Interface: pkgInfo.Interface,
		Function: newFunctionList,
		Var: pkgInfo.Var,
	}
}

