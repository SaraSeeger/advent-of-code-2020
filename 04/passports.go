package main

import (
  "log"
  "fmt"
  "strings"
  "regexp"

  "advent/utils"
)

// Part One checks for valid passports by verifying they contain all required
// fields, except Contry ID.
func partOne(passports []string) int {
  requiredFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
  valid := len(passports)
  for _, passport := range passports {
    for _, field := range requiredFields {
      if !strings.Contains(passport, field) {
        valid--
        break
      }
    }
  }
  return valid
}

// Part Two checks for valid passports by verifying they contain all required
// fields, and that those fields are correctly formmatted.
func partTwo(passports []string) int {
  fieldFormats := []*regexp.Regexp{
    regexp.MustCompile(`byr:(19[2-9][0-9]|200[0-2])\b`), // Birth Year
    regexp.MustCompile(`iyr:(201[0-9]|2020)\b`), // Issue Year
    regexp.MustCompile(`eyr:(202[0-9]|2030)\b`), // Experation Year
    regexp.MustCompile(`hgt:((59|6[0-9]|7[0-6])in|(1[5-8][0-9]|19[0-3])cm)\b`), // Hair Color
    regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)\b`), // Eye Color
    regexp.MustCompile(`hcl:#[0-9a-f]{6}\b`), // Hair Color
    regexp.MustCompile(`pid:[0-9]{9}\b`), // Passport ID
  }
  valid := len(passports)
  for _, passport := range passports {
    for _, field := range fieldFormats {
      if !field.MatchString(passport) {
        valid--
        break
      }
    }
  }
  return valid
}

func main() {
  log.SetFlags(0)

  // Read the input passports from a file.
  passports, err := utils.ReadChunks("input.txt")
  if err != nil {
    log.Fatalf("utils.ReadFile: %s", err)
  }

  // Part One:
  // fmt.Println(partOne(passports))
  // Part Two:
  fmt.Println(partTwo(passports))
}
