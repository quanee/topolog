package topo_test

import (
	"fmt"
	"testing"

	topo "github.com/quanee/topolog"
)

func TestTopo(t *testing.T) {
	g := topo.NewGraph()
	edges := [][]string{
		{"1", "3"},
		{"2", "3"},
		{"2", "4"},
		{"3", "5"},
		{"3", "6"},
		{"4", "6"},
		{"5", "7"},
		{"6", "7"},
		{"7", "8"},
		{"7", "9"},
		{"10", "4"},
		{"10", "11"},
		{"4", "11"},
		{"11", "12"},
		{"12", "13"},
		{"8", "13"},
		{"9", "13"},
		{"8", "11"},
		//{"11", "6"},
		{"6", "11"},
		{"11", "3"},
	}
	for _, edge := range edges {
		if cycle, ok := g.AddEdge(edge[0], edge[1]); ok {
			fmt.Println(cycle)
		}
	}

	/*g.AddEdge("1", "2")
	g.AddEdge("1", "3")
	g.AddEdge("2", "3")
	g.AddEdge("2", "4")
	g.AddEdge("4", "5")
	g.AddEdge("5", "2")*/

	/* g.AddEdge("0", "1")
	g.AddEdge("0", "2")
	g.AddEdge("0", "3")
	g.AddEdge("1", "4")
	g.AddEdge("3", "4")
	g.AddEdge("2", "5")
	g.AddEdge("2", "6")
	g.AddEdge("3", "7")
	g.AddEdge("6", "8")
	g.AddEdge("8", "9")
	g.AddEdge("7", "9")
	g.AddEdge("8", "10")
	g.AddEdge("8", "11")
	g.AddEdge("9", "12")
	g.AddEdge("5", "10") */
	//g.AddEdge("8", "0")
	//g.AddEdge("9", "6")

	//fmt.Println(g.TopoSequence())
}
