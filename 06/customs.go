package main

import (
  "log"
  "fmt"
  "strings"

  "advent/utils"
)

func partOne(forms []string) int {
  sum := 0
  for _, group := range forms {
    singleAns := strings.Split(group, "")
    groupAns := make(map[string]bool)
    for _, answer := range singleAns {
      groupAns[answer] = true
    }
    // Sanity check, cant be more than 26
    if len(groupAns) > 26 {
      panic(fmt.Sprintf("More than 26 (%v) answers in %q", len(groupAns), group))
    }
    sum += len(groupAns)
  }
  return sum
}

func partTwo(forms [][]string) int {
  sum := 0
  for _, group := range forms {
    if len(group) < 1 {
      panic(fmt.Sprintf("Found a group with no one in it: %v", group))
    }
    // Create a slice of slices with each persons answers.
    singleAns := [][]string{}
    for _, person := range group {
      singleAns = append(singleAns, strings.Split(person, ""))
    }
    // Look for any ansers that appear in all peoples slices.
    groupAns := []string{}
    for _, answer := range singleAns[0] {
      unanimous := true
      for _, person := range singleAns {
        found := false
        for _, personAnswer := range person {
          if personAnswer == answer {
            found = true
          }
        }
        if !found {
          unanimous = false
          break
        }
      }
      if unanimous {
        groupAns = append(groupAns, answer)
      }
    }
    sum += len(groupAns)
  }
  return sum
}

func main() {
  log.SetFlags(0)

  // Part One:
  // Read the input customs declaration forms from a file.
  //forms, err := utils.ReadChunksNoSpace("input.txt")
  //if err != nil {
  //  log.Fatalf("utils.ReadFile: %s", err)
  //}
  //fmt.Println(partOne(forms))

  // Part Two:
  forms, err := utils.ReadLinesInChunks("input.txt")
  if err != nil {
   log.Fatalf("utils.ReadFile: %s", err)
  }

  fmt.Println(partTwo(forms))
}
