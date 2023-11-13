package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	g := NewGraph()
	g.InitializeFromInput(input)
	return int(g.GetPath(Shortest))
}

func PartTwo(input string) int {
	g := NewGraph()
	g.InitializeFromInput(input)
	return int(g.GetPath(Longest))
}

type Path int

const (
	Shortest Path = iota
	Longest
)

type Node map[string]uint64
type Nodes map[string]Node
type Graph struct {
	nodes Nodes
}

func NewGraph() *Graph {
	var g Graph
	g.nodes = make(Nodes)
	return &g
}

func (g *Graph) AddConnection(first string, second string, distance uint64) {
	if _, ok := g.nodes[first]; !ok {
		g.nodes[first] = make(Node)
	}
	if _, ok := g.nodes[second]; !ok {
		g.nodes[second] = make(Node)
	}

	g.nodes[first][second] = distance
	g.nodes[second][first] = distance
}

func (g *Graph) InitializeFromInput(input string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		distance, _ := strconv.ParseUint(words[4], 10, 0)
		g.AddConnection(words[0], words[2], distance)
	}
}

func (g *Graph) GetPath(p Path) uint64 {
	var paths []uint64
	for k := range g.nodes {
		paths = append(paths, GetPathFromStart(g.nodes, k, p))
	}

	if p == Shortest {
		return slices.Min(paths)
	} else {
		return slices.Max(paths)
	}
}

func GetPathFromStart(nodes Nodes, start string, p Path) uint64 {
	// if this node only has one destination left, return the distance
	if len(nodes[start]) == 1 {
		for _, v := range nodes[start] {
			return v
		}
	}

	// make a copy of the nodes so that we can remove start for recrusive loops,
	// but, still have access to distances from the start node
	nodesCopy := make(Nodes)
	for outerK, outerV := range nodes {
		nodesCopy[outerK] = make(Node)
		for innerK, innerV := range outerV {
			nodesCopy[outerK][innerK] = innerV
		}
	}

	// delete the start node
	delete(nodesCopy, start)

	// delete all references to the start node
	for _, v := range nodesCopy {
		delete(v, start)
	}

	// find all paths from the start node
	var paths []uint64
	for k := range nodesCopy {
		paths = append(paths, nodes[start][k]+GetPathFromStart(nodesCopy, k, p))
	}

	if p == Shortest {
		return slices.Min(paths)
	} else {
		return slices.Max(paths)
	}
}
