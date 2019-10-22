package util

import (
	"github.com/awalterschulze/gographviz"
	"github.com/spectrex02/gorefer"
	"io/ioutil"
	"os"
)

//this file contains functions to output network graphs
func NewGraph(funcList []gorefer.FunctionInfo, relationship gorefer.CallRelationship) *gographviz.Graph {
	g := gographviz.NewGraph()
	graphName := "network"
	if err := g.SetName(graphName); err != nil {
		panic(err)
	}
	if err := g.SetDir(true); err != nil {
		panic(err)
	}
	if err := g.AddAttr(graphName, "bgcolor", "\"#343434\""); err != nil {
		panic(err)
	}
	if err := g.AddAttr(graphName, "layout", "dot"); err != nil {
		panic(err)
	}

	// Node設定
	nodeAttrs := make(map[string]string)
	// if _, err := gographviz.NewAttr("colorscheme"); err != nil {
	//  	panic(err)
	// }
	nodeAttrs["colorscheme"] = "white"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "black"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "white"
	nodeAttrs["fillcolor"] = "white"
	nodeAttrs["shape"] = "doublecircle"

	//add nodes
	for _, f := range funcList {
		if err := g.AddNode(graphName, f.FuncInfo.Name, nodeAttrs); err != nil {
			panic(err)
		}
	}

	//edge config
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"
	edgeAttrs["labelangle"] = "70"
	edgeAttrs["labeldistance"] = "2.5"

	//add edge
	for _, src := range funcList {
		addEdges(g, src.FuncInfo.Name, relationship, edgeAttrs)
	}
	return g
}

func addEdges(g *gographviz.Graph, src string, relations gorefer.CallRelationship, edgeAttrs map[string]string) {
	if relations[src] == nil {
		return
	}
	for _, dst := range relations[src] {
		if err := g.AddEdge(src, dst, true, edgeAttrs); err != nil {
			panic(err)
		}
	}
}

func OutputDot(g *gographviz.Graph, packageName string) {
	filepath := "result/" + packageName + "/network.dot"
	 s:= g.String()
	err := ioutil.WriteFile(filepath, []byte(s), os.ModePerm)
	if err != nil {
		panic(err)
	}
}