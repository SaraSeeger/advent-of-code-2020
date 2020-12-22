// This file specifies a custom graph type with nodes and edges.
package main

import (
  "fmt"
  "errors"
)

type ID string

type Graph struct {
  Nodes map[ID]bool
  Edges map[ID]map[ID]Edge
}

type Edge struct {
  Weight int
}

// FindUniqueDecendents returns a count of the number of unique decendednts of
// the input node.
func (g *Graph) CountUniqueDecendents(node ID) (int, error) {
  if !g.Nodes[node] {
    return 0, errors.New(fmt.Sprintf("Node %q does not exist in the graph", node))
  }
  return g.countUniqueDecendentsRecursively(node, map[ID]bool{}) - 1, nil
}

func (g *Graph) countUniqueDecendentsRecursively(node ID, visited map[ID]bool) int {
  visited[node] = true
  // Base Case:
  if len(g.Edges[node]) == 0 {
    return 1
  }
  // Recursive Case:
  numDecendents := 1
  for childID, _ := range g.Edges[node] {
    if !visited[childID] {
      numDecendents += g.countUniqueDecendentsRecursively(childID, visited)
    }
  }
  return numDecendents
}

// TotalDecendentWeight returns the total number of weighted nodes decended
// from the input node. This can loop forever if there are cycles in the graph.
func (g *Graph) TotalDecendentWeight(node ID) int {
  // Base Case:
  if g.Edges[node] == nil || len(g.Edges[node]) == 0 {
    return 0
  }
  // Recursive Case:
  totalDecendents := 0
  for id, edge := range g.Edges[node] {
    totalDecendents += edge.Weight + (edge.Weight * g.TotalDecendentWeight(id))
  }
  return totalDecendents
}

func NewGraph() *Graph {
  return &Graph{
    Nodes: map[ID]bool{},
    Edges: map[ID]map[ID]Edge{},
  }
}
