package gorefer

import (
	"fmt"
)
type Info interface {
	Show()
}

type FunctionInfo struct {
	Info
	Id       FuncId
	FuncInfo Func
	//Name     string
	//Receiver interface{}	//string or nil
	//ReceiverType interface{}	//receiver type (string or nil)
	//ReturnType interface{}	//return type (string or nil)
	//Package  string
	Call   []Func
}

//struct for storing information of struct
type StructInfo struct {
	Info
	Id StructId
	Name string
	//Member map[string]string
	Member []string
	MethodList []FunctionInfo
	Package string
}

//struct for storing information of interface
type InterfaceInfo struct {
	Info
	Id InterfaceId
	Name string
	//Methods map[string][]string
	Methods []string
	Package string
}

type VarInfo struct {
	Info
	Id VarId
	Name string
	Type string
}

type PackageInfo struct {
	Name string
	Struct []StructInfo
	Interface []InterfaceInfo
	Var []VarInfo
	Function []FunctionInfo
}

type ProjectInfo struct {
	Pkgs []PackageInfo
}
//struct for get testfiles method or function list called by some function

type Func struct {
	Name string
	ReturnType string
	Receiver string
	ReceiverType string
	Package string
}

type Call map[Func][]Func

func (f *Func) Show() {
	fmt.Println("==========Called in function==========")
	fmt.Printf("name : %v\n", f.Name)
	fmt.Printf("return type : %v\n", f.ReturnType)
	fmt.Printf("receiver: %v\n", f.Receiver)
	fmt.Printf("receiver type : %v\n", f.ReceiverType)
	fmt.Printf("package : %v\n", f.Package)
}

//interface some type
type Id interface {
	AllocateId() int
}
//id type for function
type FuncId struct {
	Id int
}

//id type for function
type StructId struct {
	Id int
}
type InterfaceId struct {
	Id int
}
type VarId struct {
	Id int
}

func (fId *FuncId) AllocateId() FuncId {
	fId.Id = fId.Id + 1
	return *fId
}

func (sId *StructId) AllocateId() StructId {
	sId.Id = sId.Id + 1
	return *sId
}

func (iId *InterfaceId) AllocateId() InterfaceId {
	iId.Id = iId.Id + 1
	return *iId
}

func (vId *VarId) AllocateId() VarId {
	vId.Id = vId.Id + 1
	return *vId
}

func (f FunctionInfo) Show() {
	fmt.Println("-----show function info-----")
	fmt.Println("Id:", f.Id.Id)
	fmt.Println("Name:", f.FuncInfo.Name)
	fmt.Println("Package:", f.FuncInfo.Package)
	fmt.Println("Receiver:", f.FuncInfo.Receiver)
	fmt.Println("Receiver Type:", f.FuncInfo.ReceiverType)
	fmt.Println("Return Type:", f.FuncInfo.ReturnType)
	fmt.Println("Call:", f.Call)
	fmt.Printf("\n")
}

func (s StructInfo) Show() {
	fmt.Println("-----show struct info-----")
	fmt.Println("Id:", s.Id.Id)
	fmt.Println("Name:", s.Name)
	fmt.Println("Package:", s.Package)
	fmt.Println("Member:", s.Member)
	fmt.Println("Methods:", s.MethodList)
	fmt.Printf("\n")
}

func (i InterfaceInfo) Show() {
	fmt.Println("-----show interface info-----")
	fmt.Println("Id:", i.Id.Id)
	fmt.Println("Name:", i.Name)
	fmt.Println("Methods:", i.Methods)
}

func (v VarInfo) Show() {
	fmt.Println("-----show var info-----")
	fmt.Println("Id:", v.Id.Id)
	fmt.Println("Name:", v.Name)
	fmt.Println("Type:", v.Type)
}
