package main

import (
  "log"
  "fmt"
  "strings"
  "math"

  "advent/utils"
)

// seatFromCode returns the row and column number for a seat from its 10 letter
// binary space partitioning code.
func seatFromCode(seat string) (float64, float64) {
  code := strings.Split(seat, "")
  row := 0.0
  for i := 0; i <= 6; i++ {
    if code[i] == "B" {
      row += math.Pow(2.0, (6.0 - float64(i)))
    }
  }
  col := 0.0
  for j := 7; j <= 9; j++ {
    if code[j] == "R" {
      col += math.Pow(2.0, (9.0 - float64(j)))
    }
  }
  return row, col
}

// For Part One, simply find the highest seat ID in the list of boarding passes.
func partOne(passes []string) float64 {
  highest := 0.0
  for _, pass := range passes {
    row, col := seatFromCode(pass)
    if sid := row * 8 + col; sid > highest {
      highest = sid
    }
  }
  return highest
}

// For Part Two, we need to find the missing seat in the middle, that's our seat.
func partTwo(passes []string) int {
  seatsTaken := make(map[int]bool)
  for _, pass := range passes {
    row, col := seatFromCode(pass)
    seatsTaken[int(row * 8 + col)] = true
  }
  i := 0
  for ; seatsTaken[i] == false; i++ {}
  for ; seatsTaken[i] == true; i++ {}
  return i
}

func main() {
  log.SetFlags(0)

  // Read the boarding passes from a file.
  passes, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Part One:
  //fmt.Println(partOne(passes))
  // Part Two:
  fmt.Println(partTwo(passes))
}
