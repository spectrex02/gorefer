package util

import (
	"encoding/json"
	"github.com/spectrex02/gorefer"
	"io/ioutil"
	"log"
	"os"
)

//for json output
type PackageJson struct {
	Name string `json:"name"`
	Struct []StructJson `json:"struct"`
	Interface []InterfaceJson `json:"interface"`
	Function []FunctionJson `json:"function"`
	Var []VarJson `json:"var"`
}

type StructJson struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Package string `json:"package"`
	Member []string `json:"member"`
	MethodList []FunctionJson `json:"methods"`
}

type FunctionJson struct {
	Id int `json:"id"`
	Info FJson `json:"info"`
	Call []FJson `json:"call"`
}

type InterfaceJson struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Package string `json:"package"`
	Methods []string `json:"methods"`
}

type VarJson struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type FJson struct {
	Name string `json:"name"`
	ReturnType string `json:"return_type"`
	Receiver string `json:"receiver"`
	ReceiverType string `json:"receiver_type"`
	Package string `json:"package"`
}

func PackageInfoToJson(info gorefer.PackageInfo) PackageJson {
	var sj []StructJson
	var ij []InterfaceJson
	var fj []FunctionJson
	var vj []VarJson
	for _, s := range info.Struct {
		sj = append(sj, StructInfoToJson(s))
	}
	for _, i := range info.Interface {
		ij = append(ij, InterfaceInfoToJson(i))
	}
	for _, f := range info.Function {
		fj = append(fj, FunctionInfoToJson(f))
	}
	for _, v := range info.Var {
		vj = append(vj, VarInfoToJson(v))
	}
	return PackageJson{
		Name:      info.Name,
		Struct:    sj,
		Interface: ij,
		Function:  fj,
		Var:       vj,
	}
}

func FunctionInfoToJson(info gorefer.FunctionInfo) FunctionJson {
	var call []FJson
	for _, f := range info.Call {
		call = append(call, FuncToJson(f))
	}
	return FunctionJson{
		Id:   info.Id.Id,
		Info: FuncToJson(info.FuncInfo),
		Call: call,
	}
}

func StructInfoToJson(info gorefer.StructInfo) StructJson {
	var methods []FunctionJson
	for _, f := range info.MethodList {
		methods = append(methods, FunctionInfoToJson(f))
	}
	return StructJson{
		Id:         info.Id.Id,
		Name:       info.Name,
		Package:    info.Package,
		Member:     info.Member,
		MethodList: methods,
	}
}

func InterfaceInfoToJson(info gorefer.InterfaceInfo) InterfaceJson {
	return InterfaceJson{
		Id:      info.Id.Id,
		Name:    info.Name,
		Package: info.Package,
		Methods: info.Methods,
	}
}

func VarInfoToJson(info gorefer.VarInfo) VarJson {
	return VarJson{
		Id:   info.Id.Id,
		Name: info.Name,
		Type: info.Type,
	}
}

func FuncToJson(f gorefer.Func) FJson {
	return FJson{
		Name:         f.Name,
		ReturnType:   f.ReturnType,
		Receiver:     f.Receiver,
		ReceiverType: f.ReceiverType,
		Package:      f.Package,
	}
}
//write to the result file as json format
func WriteJsonFile(filename string, jsonResult []byte) {
	isExist := os.Mkdir("result", os.ModePerm)
	if os.IsNotExist(isExist) {
		log.Printf("Directory to store the result is not exist. Make it.")
	}
	filepath := "result/" + filename + ".json"
	err := ioutil.WriteFile(filepath, jsonResult, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func New(info gorefer.PackageInfo) PackageJson {
	return PackageInfoToJson(info)
}

func (pkg *PackageJson) ToJson() []byte {
	result, err := json.MarshalIndent(pkg, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}