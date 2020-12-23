package main

import (
  "log"
  "fmt"
  "strconv"

  "advent/utils"
)

func isValidSum(nums []int, sum int) bool {
  for i := range nums {
    for j := range nums {
      if i != j && nums[i] + nums[j] == sum {
        return true
      }
    }
  }
  return false
}

func partOne(numbers []int) int {
  for i := 25; i < len(numbers); i++ {
    if !isValidSum(numbers[(i - 25):i], numbers[i]) {
      return numbers[i]
    }
  }
  return -1
}

func partTwo(numbers []int) int {
  goal := partOne(numbers)
  for i := 0; i < len(numbers) - 1; i++ {
    sum := numbers[i]
    for j := i + 1; j < len(numbers); j++ {
      sum += numbers[j]
      if sum == goal {
        min := goal
        max := 0
        for k := i; k <= j; k++ {
          if numbers[k] < min {
            min = numbers[k]
          }
          if numbers[k] > max {
            max = numbers[k]
          }
        }
        return min + max
      }
      if sum > goal {
        break // Small optimization.
      }
    }
  }
  return -1
}

func main() {
  log.SetFlags(0)

  // Read the XMAS numbers from a file.
  lines, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }
  if len(lines) < 26 {
    log.Fatalf("Not enough input numbers (min 25)")
  }
  numbers := make([]int, len(lines))
  for i, line := range lines {
    numbers[i], _ = strconv.Atoi(line)
  }

  // Part One:
  // fmt.Println(partOne(numbers))
  // Part Two:
  fmt.Println(partTwo(numbers))
}
