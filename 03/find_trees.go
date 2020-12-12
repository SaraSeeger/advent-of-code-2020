package main

import (
  "fmt"
  "log"
  "strings"

  "advent/utils"
)

const tree = "#"
const open = "."


// In part one, we calculate how many trees we will hit if we go at a slope of
// right 3, down 1.
func partOne(trees [][]string, hSlope, vSlope int) int {
  treesHit := 0
  // Horizontal coordinate.
  i := 0
  for j, row := range trees {
    if (j % vSlope) == 0 {
      if row[i] == tree {
        treesHit++
      }
      i = (i + hSlope) % len(row)
    }
  }
  return treesHit
}

// In part two, we need to check several slopes and multiply the number of trees
// encountered.
func partTwo(trees [][]string) int {
  // Slopes to check
  slopes := [][]int{
    []int{1, 1},
    []int{3, 1},
    []int{5, 1},
    []int{7, 1},
    []int{1, 2},
  }
  treesHitMult := 1
  for _, slope := range slopes {
    treesHitMult = treesHitMult * partOne(trees, slope[0], slope[1])
  }
  return treesHitMult
}

func main() {
  log.SetFlags(0)

  // Read the input geology from a file.
  lines, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Create a slice of slices to represent the repeating grid of trees.
  trees := [][]string{}
  for _, line := range lines {
    trees = append(trees, strings.Split(line, ""))
  }

  // Part One:
  // fmt.Println(partOne(trees, 3, 1))
  // Part Two:
  fmt.Println(partTwo(trees))
}
