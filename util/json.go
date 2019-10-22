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

type Relation struct {
	Call FunctionJson `json:"call"`
	Called []FunctionJson `json:"called"`
}
func New(info gorefer.PackageInfo) PackageJson {
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
func writeJsonFile(path string, jsonResult []byte) {
	filepath := path + ".json"
	err := ioutil.WriteFile(filepath, jsonResult, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

//convert all package information to json file
func (pkg *PackageJson) OutputResult() {
	//make directory
	path := "result/" + pkg.Name + "/"
	makeResultDir(pkg.Name)
	//struct
	structData := structToJson(pkg.Struct)
	writeJsonFile(path + "struct-list", structData)

	//interface
	interfaceData := interfaceToJson(pkg.Interface)
	writeJsonFile(path + "interface-list", interfaceData)

	//function
	functionData := functionToJson(pkg.Function)
	writeJsonFile(path + "function-list", functionData)

	//var
	varData := varToJson(pkg.Var)
	writeJsonFile(path + "var-list", varData)
}

func structToJson(list []StructJson) []byte {
	type structList struct {
		List []StructJson `json:"list"`
	}
	s := structList{List:list}
	result, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}

func interfaceToJson(list []InterfaceJson) []byte {
	type interfaceList struct {
		List []InterfaceJson `json:"list"`
	}
	i := interfaceList{List:list}
	result, err := json.MarshalIndent(i, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}

func functionToJson(list []FunctionJson) []byte {
	type functionList struct {
		List []FunctionJson `json:"list"`
	}
	f := functionList{List:list}
	result, err := json.MarshalIndent(f, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}

func varToJson(list []VarJson) []byte {
	type varList struct {
		List []VarJson `json:"list"`
	}
	v := varList{List:list}
	result, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}

func makeResultDir(name string) {
	dir := "result/" + name
	isExist := os.Mkdir("result", os.ModePerm)
	if os.IsNotExist(isExist) {
		log.Printf("Directory to store the result is already exist.")
	}
	isExist = os.Mkdir(dir, os.ModePerm)
	if os.IsNotExist(isExist) {
		log.Printf("Directory to store the result is already exist.")
	}
}

