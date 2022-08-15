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
	cycles [][]int
}

func NewGraph() *Graph {
	return &Graph{
		edges:  []*edge{},
		indeg:  map[int]map[int]struct{}{},
		nodes:  map[string]int{},
		names:  map[int]string{},
		queue:  []int{},
		visted: map[int]struct{}{},
		cycles: [][]int{},
	}
}

func (g *Graph) AddEdge(start, end string) ([][]string, bool) {
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

	g.buildCycle(g.nodes[start])

	if len(g.cycles) > 0 {
		cycles := make([][]string, 0, len(g.cycles))
		for _, cycle := range g.cycles {
			for i := 0; i < len(cycle)>>1; i++ {
				cycle[i], cycle[len(cycle)-i-1] = cycle[len(cycle)-i-1], cycle[i]
			}
			cycles = append(cycles, g.node2name(cycle))
		}
		return cycles, true
	}

	return nil, false
}

func (g *Graph) buildCycle(start int) {
	for p := range g.indeg[start] {
		if p == g.queue[0] {
			g.queue = append(g.queue, p)
			tmp := make([]int, len(g.queue)-1)
			copy(tmp, g.queue[1:])
			g.cycles = append(g.cycles, tmp)
			continue
		}

		g.queue = append(g.queue, p)

		i := len(g.queue)
		g.buildCycle(p)
		g.queue = g.queue[:i-1]
	}
}

func (g *Graph) node2name(nodes []int) []string {
	names := make([]string, 0, len(nodes))
	for _, node := range nodes {
		names = append(names, g.names[node])
	}
	return names
}

func (g *Graph) TopoSequence() (topos []string, cycl bool) {
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
	g.printCycle(g.queue)
}

func (g *Graph) printCycle(cycle []int) {
	println("-------------")
	for _, n := range cycle {
		print(g.names[n], ", ")
	}
	println("\n-------------")
}
