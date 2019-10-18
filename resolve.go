package gorefer

import "fmt"

//this file contain resolvers from *ast.FuncDecl or function name(string) to FunctionInfo

//to do
//resolve method list
//


func (pkg *PackageInfo) ResolveMethodList() {
	for _, typ := range pkg.Struct {
		var lis []FunctionInfo
		for _, f := range pkg.Function {
			if typ.Name == f.ReceiverType {
				fmt.Printf("function:%v is methods of %v type.\n", f.Name, typ.Name)
				lis = append(lis, f)
			}
		}
		typ.MethodList = lis
		fmt.Printf("methods list(type %v)-> %v\n", typ.Name, typ.MethodList)
	}
}

