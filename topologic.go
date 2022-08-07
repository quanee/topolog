package topo

import (
	"fmt"
	"sort"
)

type edge struct {
	start int
	end   int
}

type set struct {
	s map[int]struct{}
}

type Graph struct {
	count  int
	edges  []*edge
	indeg  map[int]map[int]struct{}
	nodes  map[string]int
	names  map[int]string
	ufs    *unionfindset
	queue  []int
	visted map[int]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		edges: []*edge{},
		indeg: make(map[int]map[int]struct{}),
		nodes: make(map[string]int),
		names: make(map[int]string),
		ufs:   newunionfindset(1000),
		queue: make([]int, 0),
	}
}

func (g *Graph) AddEdge(start, end string) error {
	if start == end {
		return nil
	}

	if _, ok := g.nodes[start]; !ok {
		g.nodes[start] = g.count
		g.names[g.count] = start
		g.count++
	}
	if _, ok := g.nodes[end]; !ok {
		g.nodes[end] = g.count
		g.names[g.count] = end
		g.count++
	}

	if g.indeg[g.nodes[end]] == nil {
		g.indeg[g.nodes[end]] = make(map[int]struct{})
	}
	g.indeg[g.nodes[end]][g.nodes[start]] = struct{}{}
	g.edges = append(g.edges, &edge{start: g.nodes[start], end: g.nodes[end]})

	g.queue = []int{g.nodes[end], g.nodes[start]}
	g.buildCycle(g.nodes[start])

	fmt.Print("### ")
	g.printQ()

	return nil
}

func (g *Graph) printQ() {
	for _, n := range g.queue {
		fmt.Printf("%v, ", g.names[n])
	}
	println()
}

func (g *Graph) buildCycle(start int) bool {
	if g.names[start] == "5" {
		println()
	}
	for p := range g.indeg[start] {
		if p == g.queue[0] {
			g.queue = append(g.queue, p)
			return true
		}
		g.queue = append(g.queue, p)

		if !g.buildCycle(p) {
			g.queue = g.queue[:1]
		}
		return true
	}

	return false
}

func (g *Graph) genSequence(sorted []*edge) []string {
	retSet := map[int]struct{}{}
	var sequences []string

	for _, node := range sorted {
		if _, ok := retSet[node.start]; !ok {
			retSet[node.start] = struct{}{}
			sequences = append(sequences, g.names[node.start])
		}
	}

	for node, name := range g.names {
		if _, ok := retSet[node]; !ok {
			sequences = append(sequences, name)
		}
	}

	return sequences
}

func (g *Graph) deleteEdge(currnode int, nodes *[]*edge) {
	for _, delEdge := range g.edges {
		if delEdge.start == currnode {
			*nodes = append(*nodes, delEdge)
			delete(g.indeg[delEdge.end], delEdge.start)
		}
		if len(g.indeg[currnode]) == 0 {
			delete(g.indeg, currnode)
		}
	}
}

func (g *Graph) topoSort() ([]*edge, bool) {
	var ret []*edge
	sorted := make(map[int]struct{})
	oldLen := len(sorted)
	/*	defer func() {
		retSet := map[string]struct{}{}
		for _, e := range ret {
			if _, ok := retSet[g.names[e.start]]; !ok {
				retSet[g.names[e.start]] = struct{}{}
			}
		}
	}()*/

	for len(sorted) != g.count {
		var zeroDegreeNodes []int
		for node := range g.nodes {
			if len(g.indeg[g.nodes[node]]) == 0 {
				zeroDegreeNodes = append(zeroDegreeNodes, g.nodes[node])
			}
		}
		sort.Ints(zeroDegreeNodes)
		if len(zeroDegreeNodes) > 0 {
			for _, node := range zeroDegreeNodes {
				if _, ok := sorted[node]; !ok {
					g.deleteEdge(node, &ret)
					sorted[node] = struct{}{}
				}
			}
			if len(sorted) == oldLen {
				return nil, false
			}
			oldLen = len(sorted)
		} else {
			break
		}
	}

	if len(ret) != len(g.edges) {
		return nil, false
	}

	return ret, true
}

func (g *Graph) TopoSequence() ([]string, bool) {
	sorted, ok := g.topoSort()
	g.printTopoEdges(sorted)
	if !ok {
		return nil, false
	}
	fmt.Println("$$$", sorted)

	return g.genSequence(sorted), true
}

func (g *Graph) PrintParent() {
	for node, parent := range g.indeg {
		fmt.Printf("%v -> ", g.names[node])
		for p := range parent {
			fmt.Printf("%v ", g.names[p])
		}
		println()
	}
}

func (g *Graph) printTopoEdges(sorted []*edge) {
	for _, e := range sorted {
		fmt.Printf("%v -> %v\n", g.names[e.start], g.names[e.end])
	}
}
