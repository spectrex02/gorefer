package gorefer

import (
	"fmt"
)
type Info interface {
	Show()
}

type FunctionInfo struct {
	Info
	Id       int
	Name     string
	Receiver interface{}	//string or nil
	ReceiverType interface{}	//receiver type (string or nil)
	ReturnType interface{}	//return type (string or nil)
	Called   []int
	Package  string
	Body Block	//function body
}

//struct for storing information of struct
type StructInfo struct {
	Info
	Id int
	Name string
	//Member map[string]string
	Member []string
	MethodList []FunctionInfo
	Package string
}

//struct for storing information of interface
type InterfaceInfo struct {
	Info
	Id int
	Name string
	//Methods map[string][]string
	Methods []string
	Package string
}

type VarInfo struct {
	Info
	Id int
	Name string
	Type string
}

type PackageInfo struct {
	Struct []StructInfo
	Interface []InterfaceInfo
	Var []VarInfo
	Function []FunctionInfo
}

//struct for get testfiles method or function list called by some function
type CalledFunction struct {
	Info
	Receiver string
	Name string
}


//struct for block stmt {...}
type Block struct {
	Info
	Called []CalledFunction
	AssignVars []VarInfo
	Sub []SubBlock
}

type SubBlock struct {
	Block
}


func (c *CalledFunction) Show() {
	fmt.Printf("recv -> %s : name -> %s\n", c.Receiver, c.Name)
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

func (fId *FuncId) AllocateId() int {
	fId.Id = fId.Id + 1
	return fId.Id
}

func (sId *StructId) AllocateId() int {
	sId.Id = sId.Id + 1
	return sId.Id
}

func (iId *InterfaceId) AllocateId() int {
	iId.Id = iId.Id + 1
	return iId.Id
}

func (vId *VarId) AllocateId() int {
	vId.Id = vId.Id + 1
	return vId.Id
}

func (f FunctionInfo) Show() {
	fmt.Println("-----show function info-----")
	fmt.Println("Id:", f.Id)
	fmt.Println("Name:", f.Name)
	fmt.Println("Package:", f.Package)
	fmt.Println("Receiver:", f.Receiver)
	fmt.Println("Receiver Type:", f.ReceiverType)
	fmt.Println("Return Type:", f.ReturnType)
	fmt.Println("Called:", f.Called)
	fmt.Println("Body:", f.Body)
	fmt.Printf("\n")
}

func (s StructInfo) Show() {
	fmt.Println("-----show struct info-----")
	fmt.Println("Id:", s.Id)
	fmt.Println("Name:", s.Name)
	fmt.Println("Package:", s.Package)
	fmt.Println("Member:", s.Member)
	fmt.Println("Methods:", s.MethodList)
	fmt.Printf("\n")
}

func (i InterfaceInfo) Show() {
	fmt.Println("-----show interface info-----")
	fmt.Println("Id:", i.Id)
	fmt.Println("Name:", i.Name)
	fmt.Println("Methods:", i.Methods)
}

func (v VarInfo) Show() {
	fmt.Println("-----show var info-----")
	fmt.Println("Id:", v.Id)
	fmt.Println("Name:", v.Name)
	fmt.Println("Type:", v.Type)
}
