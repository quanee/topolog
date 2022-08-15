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
		/*{"0", "1"},
		{"1", "2"},
		{"2", "3"},
		{"3", "4"},
		{"1", "3"},
		{"4", "1"},*/
	}
	for _, edge := range edges {
		cycle, ok := g.AddEdge(edge[0], edge[1])
		if ok {
			fmt.Printf("## start: %v end: %v, cycle: %v, cycled: %v\n", edge[0], edge[1], cycle, ok)
		}
	}

	/* edges = [][]string{
		{"0", "1"},
		{"1", "2"},
		{"2", "3"},
		{"3", "4"},
		{"1", "3"},
		{"4", "1"},
	} */

	/*edges = [][]string{
		{"0", "1"},
		{"0", "2"},
		{"0", "3"},
		{"1", "4"},
		{"3", "4"},
		{"2", "5"},
		{"2", "6"},
		{"3", "7"},
		{"6", "8"},
		{"8", "9"},
		{"7", "9"},
		{"8", "10"},
		{"8", "11"},
		{"9", "12"},
		{"5", "10"},
		//{"8", "0"},
		{"9", "6"},
	}*/
	/* for _, edge := range edges {
		if cycle, ok := g.AddEdge(edge[0], edge[1]); ok {
			fmt.Println(cycle)
		}
	} */

	//fmt.Println(g.TopoSequence())
}
