package gorefer


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