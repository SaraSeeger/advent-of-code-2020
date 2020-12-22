package main

import (
  "log"
  "fmt"
  "strconv"
  "regexp"
  "errors"

  "advent/utils"
)

var instructionRE = regexp.MustCompile(`(jmp|acc|nop) (\+\d+|-\d+)`)

// Returns the final accumulator value and "true" if the program exited normally
// or "false" if we found an infinite loop.
func partOne(instructions []string) (int, bool) {
  accumulator := 0
  visited := make([]bool, len(instructions))
  i := 0
  for {
    // Check for notmal exit.
    if i >= len(instructions) {
      return accumulator, true
    }
    // Check for infinite loop.
    if visited[i] {
      return accumulator, false
    }
    visited[i] = true
    matched := instructionRE.FindStringSubmatch(instructions[i])
    val, _ := strconv.Atoi(matched[2])
    switch matched[1] {
    case "nop":
      i++
    case "acc":
      accumulator += val
      i++
    case "jmp":
      i += val
    }
  }
}

// Brute force solution, just try changing each instruction and see if the
// program terminates.
func partTwo(instructions []string) (int, error) {
  for i := 0; i < len(instructions); i++ {
    matched := instructionRE.FindStringSubmatch(instructions[i])
    if matched[1] == "acc" {
      continue
    }
    original := instructions[i]
    if matched[1] == "jmp" {
      instructions[i] = "nop " + matched[2]
    } else {
      instructions[i] = "jmp " + matched[2]
    }
    if acc, terminated := partOne(instructions); terminated {
      return acc, nil
    }
    // Make sure we set the instruction back to the original
    instructions[i] = original
  }
  return 0, errors.New(fmt.Sprint("No change found that would let the program terminate."))
}

func main() {
  log.SetFlags(0)

  // Read the boot instructions from a file.
  instructions, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Part One:
  fmt.Println(partOne(instructions))
  // Part Two:
  fmt.Println(partTwo(instructions))
}
