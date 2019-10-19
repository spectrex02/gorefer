package gorefer

import (
	"errors"
	"fmt"
	"log"
)

//this file contain resolvers from *ast.FuncDecl or function name(string) to FunctionInfo

//to do
//resolve method list
//


func (pkg *PackageInfo) ResolveMethodList() {
	for _, typ := range pkg.Struct {
		var lis []FunctionInfo
		for _, f := range pkg.Function {
			if typ.Name == f.FuncInfo.ReceiverType {
				fmt.Printf("function:%v is methods of %v type.\n", f.FuncInfo.Name, typ.Name)
				lis = append(lis, f)
			}
		}
		typ.MethodList = lis
		fmt.Printf("methods list(type %v)-> %v\n", typ.Name, typ.MethodList)
	}
}

//util function
func FindFunctionFromId(pkg PackageInfo, id int) (interface{}, error) {
	for _, f := range pkg.Function {
		if f.Id.Id == id {
			return f, nil
		}
	}
	log.Printf("Not found such function.")
	err := errors.New("Not found such function.")
	return nil, err
}

func FindStructFromId(pkg PackageInfo, id int) (interface{}, error) {
	for _, s := range pkg.Struct {
		if s.Id.Id == id {
			return s, nil
		}
	}
	log.Printf("Not found such struct.")
	err := errors.New("Not found such struct.")
	return nil, err
}

func FindInterfaceFromId(pkg PackageInfo, id int) (interface{}, error) {
	for _, i := range pkg.Interface {
		if i.Id.Id == id {
			return i, nil
		}
	}
	log.Printf("Not found such interface.")
	err := errors.New("Not found such interface.")
	return nil, err
}

func FindVarFromId(pkg PackageInfo, id int) (interface{}, error) {
	for _, v := range pkg.Var {
		if v.Id.Id == id {
			return v, nil
		}
	}
	log.Printf("Not found such valiable.")
	err := errors.New("Not found such variable.")
	return nil, err
}

