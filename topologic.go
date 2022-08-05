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

	g.queue = []int{g.nodes[end]}
	if g.buildCycle(g.nodes[start]) {
		fmt.Println("###", g.queue)
	}

	return nil
}

func (g *Graph) printQ() {
	for _, n := range g.queue {
		fmt.Printf("%v, ", g.names[n])
	}
	println()
}

func (g *Graph) buildCycle(start int) bool {
	for p := range g.indeg[start] {
		if p == g.queue[0] {
			return true
		}
		g.queue = append(g.queue, p)

		if g.buildCycle(p) {
			g.queue = g.queue[:1]
		}
	}

	return false
}

func (g *Graph) genSequence(sorted []int) []string {
	retSet := map[int]struct{}{}

	for _, node := range g.nodes {
		if _, ok := retSet[node]; !ok {
			retSet[node] = struct{}{}
			sorted = append(sorted, node)
		}
	}

	rets := []string{}
	for _, n := range sorted {
		rets = append(rets, g.names[n])
	}
	return rets
}

func (g *Graph) deleteEdge(curnode int, nodes *[]int, nodeset map[int]struct{}) {
	for _, delEdge := range g.edges {
		if delEdge.start == curnode {
			println(delEdge.start)
			if _, ok := nodeset[delEdge.start]; !ok {
				*nodes = append(*nodes, delEdge.start)
				nodeset[delEdge.start] = struct{}{}
				fmt.Println(*nodes)
			}
			delete(g.indeg[delEdge.end], delEdge.start)
		}
		if len(g.indeg[curnode]) == 0 {
			delete(g.indeg, curnode)
		}
	}
}

func (g *Graph) topoSort() ([]int, bool) {
	nodes := []int{}
	nodeset := make(map[int]struct{})
	sorted := make(map[int]struct{})
	oldLen := len(sorted)
	/* defer func() {
		retSet := map[int]struct{}{}
		for _, e := range nodes {
			if _, ok := retSet[e]; !ok {
				retSet[e] = struct{}{}
			}
		}
	}() */

	for len(sorted) != g.count {
		zeroDegreeNodes := []int{}
		for _, node := range g.nodes {
			if len(g.indeg[node]) == 0 {
				zeroDegreeNodes = append(zeroDegreeNodes, node)
			}
		}
		sort.Ints(zeroDegreeNodes)
		if len(zeroDegreeNodes) > 0 {
			for _, node := range zeroDegreeNodes {
				if _, ok := sorted[node]; !ok {
					g.deleteEdge(node, &nodes, nodeset)
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

	if len(nodes) != len(g.edges) {
		return nil, false
	}

	return nodes, true
}

func (g *Graph) TopoSequence() ([]string, bool) {
	sorted, ok := g.topoSort()
	if !ok {
		return nil, false
	}

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
