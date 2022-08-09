package topo

import (
	"sort"
)

type edge struct {
	start int
	end   int
}

type Graph struct {
	count  int
	edges  []*edge
	indeg  map[int]map[int]struct{}
	nodes  map[string]int
	names  map[int]string
	queue  []int
	visted map[int]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		edges:  []*edge{},
		indeg:  make(map[int]map[int]struct{}),
		nodes:  make(map[string]int),
		names:  make(map[int]string),
		queue:  make([]int, 0),
		visted: make(map[int]struct{}),
	}
}

func (g *Graph) AddEdge(start, end string) ([]string, bool) {
	if start == end {
		return nil, false
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
	g.visted = make(map[int]struct{})
	if g.buildCycle(g.nodes[start]) {
		return g.node2name(g.queue[1:]), true
	}

	return nil, false
}

func (g *Graph) buildCycle(start int) bool {
	for p := range g.indeg[start] {
		if _, ok := g.visted[p]; ok {
			continue
		}
		g.visted[p] = struct{}{}

		if p == g.queue[0] {
			g.queue = append(g.queue, p)
			return true
		}
		g.queue = append(g.queue, p)

		i := len(g.queue)
		if !g.buildCycle(p) {
			g.queue = g.queue[:i-1]
		} else {
			return true
		}
	}

	return false
}

func (g *Graph) node2name(nodes []int) (names []string) {
	for i := len(nodes) - 1; i >= 0; i-- {
		names = append(names, g.names[nodes[i]])
	}
	return
}

func (g *Graph) TopoSequence() ([]string, bool) {
	sorted, ok := g.topoSort()
	if !ok {
		return nil, true
	}

	return g.topoSequence(sorted), false
}

func (g *Graph) topoSort() ([]*edge, bool) {
	var topoedge []*edge
	delnode := make(map[int]struct{})
	oldLen := len(delnode)

	for len(delnode) != g.count {
		var zerodegnode []int
		for node := range g.nodes {
			if len(g.indeg[g.nodes[node]]) == 0 {
				zerodegnode = append(zerodegnode, g.nodes[node])
			}
		}
		sort.Ints(zerodegnode)
		if len(zerodegnode) > 0 {
			for _, node := range zerodegnode {
				if _, ok := delnode[node]; !ok {
					topoedge = append(topoedge, g.deleteEdge(node)...)
					delnode[node] = struct{}{}
				}
			}
			if len(delnode) == oldLen {
				return nil, false
			}
			oldLen = len(delnode)
		} else {
			break
		}
	}

	if len(topoedge) != len(g.edges) {
		return nil, false
	}

	return topoedge, true
}

func (g *Graph) deleteEdge(currnode int) (nodes []*edge) {
	for _, edge := range g.edges {
		if edge.start == currnode {
			nodes = append(nodes, edge)
			delete(g.indeg[edge.end], edge.start)
		}
		if len(g.indeg[currnode]) == 0 {
			delete(g.indeg, currnode)
		}
	}

	return
}

func (g *Graph) topoSequence(sorted []*edge) []string {
	set := map[int]struct{}{}
	sequences := make([]int, 0, len(g.nodes))

	for _, node := range sorted {
		if _, ok := set[node.start]; !ok {
			set[node.start] = struct{}{}
			sequences = append(sequences, node.start)
		}
	}

	for node := range g.names {
		if _, ok := set[node]; !ok {
			sequences = append(sequences, node)
		}
	}

	return g.node2name(sequences)
}

func (g *Graph) printParent() {
	for node, parent := range g.indeg {
		println(g.names[node], " -> ")
		for p := range parent {
			print(g.names[p], " ")
		}
		println()
	}
}

func (g *Graph) printTopoEdges(sorted []*edge) {
	for _, e := range sorted {
		println(g.names[e.start], " -> ", g.names[e.end])
	}
}

func (g *Graph) printQ() {
	println("-------------")
	for _, n := range g.queue {
		print(g.names[n], ", ")
	}
	println("\n-------------")
}
