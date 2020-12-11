package main

import (
  "log"
  "regexp"
  "fmt"
  "strconv"
  "strings"

  "advent/utils"
)

// In part one, we are seeing if the proper number of a spicific character appears.
func partOne(passwords [][]string) int {
  policyFormat := "^[^%s]*(%s[^%s]*){%s,%s}$"
  valid := 0
  for _, policy := range passwords {
    policyRegex := regexp.MustCompile(fmt.Sprintf(policyFormat, policy[3], policy[3], policy[3], policy[1], policy[2]))
    if policyRegex.FindStringSubmatch(policy[4]) != nil {
      valid++
    }
  }
  return valid
}

// In part two, we are seeing if a charature appears in the proper index.
func partTwo(passwords [][]string) int {
  valid := 0
  for _, policy := range passwords {
    i, err := strconv.Atoi(policy[1])
    if err != nil {
      log.Fatalf("could not convert %s (of %v) to int: %v", policy[1], policy, err)
    }
    i--
    j, err := strconv.Atoi(policy[2])
    if err != nil {
      log.Fatalf("could not convert %s (of %v) to int: %v", policy[2], policy, err)
    }
    j--
    passwordChars := strings.Split(policy[4], "")
    if (passwordChars[i] == policy[3]) != (passwordChars[j] == policy[3]) {
         valid++
    }
  }
  return valid
}

func main() {
  log.SetFlags(0)

  // Read the input expense report from a file.
  lines, err := utils.ReadLines("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Parse the passwords and policies using a regex. Each entry in the slice will be formatted:
  // []string{<full line>, <min_val>, <max_val>, <match_letter>, <password>}
  lineRegex := regexp.MustCompile(`([0-9]+)-([0-9]+)\s([a-z]):\s([a-z]+)`)
  passwords := [][]string{}
  for _, line := range lines {
    if match := lineRegex.FindStringSubmatch(line); match != nil {
      passwords = append(passwords, match)
    } else {
      log.Fatalf("invalid password line input: %q", line)
    }
  }

  // Part One:
  // fmt.Println(partOne(passwords))
  // Part Two:
  fmt.Println(partTwo(passwords))
}
