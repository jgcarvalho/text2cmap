package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/awalterschulze/gographviz"
)

type node struct {
	name  string
	depth int
	attr  map[string]string
}

type extraconn struct {
	id    string
	nodes [2]string
}

func cm(lines []string) {

	nodes := make([]node, len(lines))
	var nodeName string
	var nodeDepth int
	var content string

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			nodeAttr := make(map[string]string)
			nodeName = fmt.Sprintf("node_%d", i)
			nodeDepth = strings.Count(lines[i], "\t")
			content = strings.TrimSpace(lines[i])
			if strings.Count(content, "--") >= 2 {
				nodeAttr["label"] = "\"" + strings.Trim(content, "--") + "\""
				nodeAttr["shape"] = "plaintext"
			} else {
				nodeAttr["label"] = "\"" + content + "\""
				nodeAttr["shape"] = "box"
				nodeAttr["style"] = "\"rounded,filled\""
				nodeAttr["fillcolor"] = "aliceblue"
			}
			nodes[i] = node{nodeName, nodeDepth, nodeAttr}
		}
	}

	conn := make(map[string][]string)

	var ndepth int
	for i := 0; i < len(nodes); i++ {
		ndepth = nodes[i].depth
		for j := i + 1; j < len(nodes); j++ {
			if nodes[j].depth <= (ndepth) {
				break
			}
			if nodes[j].depth == (ndepth + 1) {
				if _, ok := conn[nodes[i].name]; !ok {
					conn[nodes[i].name] = make([]string, 1)
					conn[nodes[i].name][0] = nodes[j].name
				} else {
					conn[nodes[i].name] = append(conn[nodes[i].name], nodes[j].name)
				}
			}
		}
	}

	// fmt.Println(nodes)
	// fmt.Println(conn)

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	for _, v := range nodes {
		g.AddNode("G", v.name, v.attr)
	}

	for key, val := range conn {
		for _, v := range val {
			g.AddEdge(key, v, true, nil)
		}
	}
	s := g.String()
	fmt.Println(s)

}

// func createCM(lines []string) {
// 	var nodeName string
// g := gographviz.NewGraph()
// g.SetName("G")
// g.SetDir(true)
// for i, v := range lines {
// 	nodeName = fmt.Sprintf("node_%d", i)
// 	nodeAttrs := map[string]string{"label": "\"" + v + "\""}
// 	g.AddNode("G", nodeName, nodeAttrs)
// }
// // g.AddNode("G", "Hello", nil)
// // g.AddNode("G", "World", nil)
// // g.AddEdge("Hello", "World", true, nil)
// s := g.String()
// fmt.Println(s)
// }

func parserT2CM(fn string) {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		//fail
	}
	lines := strings.Split(string(content), "\n")
	// fmt.Println(lines)
	cm(lines)
}

func main() {
	// g := gographviz.NewGraph()
	// g.SetName("G")
	// g.SetDir(true)
	// g.AddNode("G", "Hello", nil)
	// g.AddNode("G", "World", nil)
	// g.AddEdge("Hello", "World", true, nil)
	// s := g.String()
	// fmt.Println(s)
	parserT2CM("./teste.t2c")
}
