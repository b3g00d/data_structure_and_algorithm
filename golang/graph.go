package main

import "fmt"

type Graph struct {
	name  string
	nodes []*Node
	edges map[Node][]*Node
}

type Node struct {
	value string
}

func (g *Graph) insertNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) insertEdge(pointer *Node, pointee *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*pointer] = append(g.edges[*pointer], pointee)
}


func (g Graph) preview() {
	for i := 0; i < len(g.nodes); i++ {
		node := g.nodes[i]
		near := g.edges[*node]
		for j := 0; j < len(near); j++ {
			child := near[j]
			fmt.Printf("%s -> %s\n", node.value, child.value)
		}
	}
}


func main() {
	my_g := Graph{name: "My_G"}

	nQuang := Node{"Quang"}
	nHoang := Node{"Hoang"}
	nNguyen := Node{"NGuyen"}
	nVu := Node{"Vu"}
	nGia := Node{"Gia"}
	nTrang := Node{"Trang"}
	my_g.insertNode(&nQuang)
	my_g.insertNode(&nHoang)
	my_g.insertNode(&nNguyen)
	my_g.insertNode(&nVu)
	my_g.insertNode(&nGia)
	my_g.insertNode(&nTrang)


	my_g.insertEdge(&nNguyen, &nQuang)
	my_g.insertEdge(&nQuang, &nHoang)
	my_g.insertEdge(&nVu, &nQuang)
	my_g.insertEdge(&nVu, &nGia)
	my_g.insertEdge(&nGia, &nTrang)

	my_g.preview()
}
