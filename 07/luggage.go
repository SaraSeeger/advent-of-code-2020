package main

import (
  "log"
  "fmt"
  "strings"
  "regexp"
  "strconv"

  "advent/utils"
)

func partOne(rules []string) (int, error) {
  var headRE = regexp.MustCompile(`([a-z]+ [a-z]+) bags? contain \d ([a-z]+ [a-z]+) bags?`)
  var tailRE = regexp.MustCompile(`([a-z]+ [a-z]+) bags?`)
  // Store the luggage rules in a graph, ingnore numbers for now.
  graph := NewGraph()
  for _, rule := range rules {
    var node ID
    parents := []ID{}
    // Split the rule by commas, apply headRE to the first result, and tailRE to
    // any other results. If headRE fails, we are dealing with a leaf node. This
    // will find the string values for the node and its children.
    processing := strings.Split(rule, ",")
    if matched := headRE.FindStringSubmatch(processing[0]); matched == nil {
      node = ID(tailRE.FindStringSubmatch(processing[0])[1])
    } else {
      node = ID(matched[1])
      parents = append(parents, ID(matched[2]))
    }
    for i := 1; i < len(processing); i++ {
      parents = append(parents, ID(tailRE.FindStringSubmatch(processing[i])[1]))
    }
    // Now add the nodes and relationships in the graph.
    graph.Nodes[node] = true
    if graph.Edges[node] == nil {
      graph.Edges[node] = map[ID]Edge{}
    }
    for _, parent := range parents {
      graph.Nodes[parent] = true
      if graph.Edges[parent] == nil {
        graph.Edges[parent] = map[ID]Edge{}
      }
      graph.Edges[parent][node] = Edge{}
    }
  }
  return graph.CountUniqueDecendents(ID("shiny gold"))
}

func partTwo(rules []string) int {
  var headRE = regexp.MustCompile(`([a-z]+ [a-z]+) bags? contain (\d) ([a-z]+ [a-z]+) bags?`)
  var headNoChildrenRE = regexp.MustCompile(`([a-z]+ [a-z]+) bags?`)
  var tailRE = regexp.MustCompile(`(\d) ([a-z]+ [a-z]+) bags?`)
  // Store the luggage rules in a graph with weighted edges.
  graph := NewGraph()
  for _, rule := range rules {
    var node ID
    children := map[ID]Edge{}
    // Split the rule by commas, apply headRE to the first result, and tailRE to
    // any other results. If headRE fails, we are dealing with a leaf node. This
    // will find the string values for the node and its children.
    processing := strings.Split(rule, ",")
    if matched := headRE.FindStringSubmatch(processing[0]); matched == nil {
      node = ID(headNoChildrenRE.FindStringSubmatch(processing[0])[1])
    } else {
      node = ID(matched[1])
      weight, _ := strconv.Atoi(matched[2]) // I know this is bad style :P
      children[ID(matched[3])] = Edge{Weight: weight}
    }
    for i := 1; i < len(processing); i++ {
      matched := tailRE.FindStringSubmatch(processing[i])
      weight, _ := strconv.Atoi(matched[1])
      children[ID(matched[2])] = Edge{Weight: weight}
    }
    // Now add the nodes and edges to the graph.
    graph.Nodes[node] = true
    graph.Edges[node] = children
    for child, _ := range children {
      graph.Nodes[child] = true
    }
  }
  return graph.TotalDecendentWeight(ID("shiny gold"))
}

func main() {
  log.SetFlags(0)

  // Read the luggage rules from a file.
  rules, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Part One:
  // if val, err := partOne(rules); err != nil {
  //   fmt.Println(err)
  // } else {
  //   fmt.Println(val)
  // }
  // Part Two:
  fmt.Println(partTwo(rules))
}
