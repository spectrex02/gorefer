package gorefer

import (
	"fmt"
	"github.com/spectrex02/goref/pkg/parse"
	"io/ioutil"
	"strings"

	//"log"
	"path/filepath"
)

type Node interface {
	Show()
	IsDir() bool
}

type Dir struct {
	Node
	Child []Node
	Name string
	Path string
	Package string

}

type File struct {
	Node
	Name string
	Path string
	Info FileInfo
	Package string
}

//don't use Root struct
type Root struct {
	Node
	Nodes []Node
	RootPath string
	Package string
}
func (f File) IsDir() bool { return false}
func (d Dir) IsDir() bool { return true}
func (root Root) IsDir() bool { return true}

func (f File) Show() {
	fmt.Printf("File name -> %v\n", f.Name)
	//f.Info.Show()
}

func (d Dir) Show() {
	fmt.Printf("Directory name and package name -> %v, %v\n", d.Name, d.Package)
	for _, child := range d.Child {
		if child.IsDir() {
			child.Show()
			continue
		}
		child.Show()
	}
}

func (root Root) Show() {
	//fmt.Printf("Root directory name -> %v\n", root.RootPath)
	for _, node := range root.Nodes {
		if node.IsDir() {
			node.Show()
			continue
		}
		node.Show()
	}
}

func New(path string) Dir {
	p := parse.New()
	nodes, packageName := getTree(path, p)
	root := Dir{
		Child: nodes,
		Name: trim(path),
		Path: path,
		Package: packageName,
	}
	root.GivePackageName()
	return root
}

//trim root directory name
func trim(path string) string {
	if strings.Contains(path, "/") {
		index := strings.LastIndex(path, "/")
		length := len(path)
		name := path[index+1 : length]
		return name
	}
	return path
}


func getTree(path string, p *parse.Parser) ([]Node, string) {
	//fmt.Println("now in getTree function. path ->", path)
	var nodes []Node
	var name string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		//if file is directory, we have to search files in this directory recursively
		if file.IsDir() {
			child, packageName := getTree(filepath.Join(path, file.Name()), p)
			dir := Dir{
				Child: child,
				Name:file.Name(),
				Path:path,
				Package: packageName,
			}
			nodes = append(nodes, dir)
			continue
		}
		//if file is file
		//we have to check whether go file or not.
		if !strings.Contains(file.Name(), ".go") { continue }
		//fmt.Println("file path ===== ", filepath.Join(path, file.Name()))
		info := p.GetFileInfo(filepath.Join(path, file.Name()))
		f := File{
			Name: file.Name(),
			Path: filepath.Join(path, file.Name()),
			Info: info,
			Package: info.Package,
		}
		//fmt.Println(f.Name)
		nodes = append(nodes, f)
		name = info.Package
	}
	//fmt.Println("length of nodes: ", len(nodes))
	return nodes, name
}

//initially, Dir struct can't have its own package name.
//so this method give package name field to it.
func (d *Dir) GivePackageName() {
	for _, node := range d.Child {
		if node.IsDir() {
			node.(Dir).givePackageName()
		}
	}
}

func (d Dir) givePackageName() {
	for _, node := range d.Child {
		if node.IsDir() {
			node.(Dir).givePackageName()
			continue
		}

		d.Package = node.(File).Package
	}
}

//traverse package and return Dir.
//Because some files belong to one package, so to identify some package, Dir is better.
func (d Dir) TraversePackage(packageName string) (Dir, bool) {
	for _, node := range d.Child {
		if node.IsDir() {
			if node.(Dir).Package == packageName {
				return node.(Dir), true
			}
			_, _ = node.(Dir).TraversePackage(packageName)
		}
	}
	if packageName == d.Package {
		return d, true
	}
	return d, false
}