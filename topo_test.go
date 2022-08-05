package topo_test

import (
	"fmt"
	"testing"

	topo "github.com/quanee/topolog"
)

func TestTopo(t *testing.T) {
	g := topo.NewGraph()
	g.AddEdge("1", "3")
	g.AddEdge("2", "3")
	g.AddEdge("2", "4")
	g.AddEdge("3", "5")
	g.AddEdge("3", "6")
	g.AddEdge("4", "6")
	g.AddEdge("5", "7")
	g.AddEdge("6", "7")
	g.AddEdge("7", "8")
	g.AddEdge("7", "9")
	g.AddEdge("10", "4")
	g.AddEdge("10", "11")
	g.AddEdge("4", "11")
	g.AddEdge("11", "12")
	g.AddEdge("12", "13")
	g.AddEdge("8", "13")
	g.AddEdge("9", "13")
	g.AddEdge("8", "11")
	//g.AddEdge("11", "6")

	fmt.Println(g.TopoSequence())
}
