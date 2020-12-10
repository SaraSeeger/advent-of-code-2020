package main

import (
  "fmt"
  "log"
  "os"
  "sort"
  "strconv"

  "advent/utils"
)

// For Part One, I used an algorithm that moves inward from both sides of a
// sorted slice to find the 2 entries that will add to 2020. This does not
// generalize to 3 entries.
func twoEntries(vals sort.IntSlice) {
  vals.Sort()
  for low, high := 0, len(vals) - 1; low < high; {
    switch {
    case vals[low] + vals[high] < 2020:
      low++
      continue
    case vals[low] + vals[high] == 2020:
      fmt.Println(vals[low] * vals[high])
      os.Exit(0) // Exit program, we have the answer.
    default:
      high--
      continue
    }
  }
  log.Fatal("no 2 values add to 2020 in expense report")
}

// For Part Two, I went with a brute force solution, it was getting late.
func threeEntries(vals []int) {
  for i := range vals {
    for j := range vals {
      for k := range vals {
        if i != j && i != k && j != k && vals[i] + vals[j] + vals[k] == 2020 {
          fmt.Println(vals[i] * vals[j] * vals[k])
          os.Exit(0)
        }
      }
    }
  }
  log.Fatal("no 3 values add to 2020 in the expense report")
}

func main() {
  log.SetFlags(0)

  // Read the input expense report from a file.
  lines, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }
  var vals sort.IntSlice
  for _, line := range lines {
    if val, err := strconv.Atoi(line); err != nil {
      log.Fatalf("could not convert %q to int: %s", line, err)
    } else {
      vals = append(vals, val)
    }
  }

  // Part One.
  // twoEntries(vals)
  // Part Two.
  threeEntries(vals)
}
