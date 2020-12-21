package utils

import (
  "bufio"
  "os"
)

// ReadLines returns the lines of an input file in a string slice.
func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, scanner.Err()
}

// ReadChunks returns chunks of text seperated by blank lines in an input file
// as a string slice.
func ReadChunks(path string) ([]string, error) {
  lines, err := ReadLines(path)
  if err != nil {
    return nil, err
  }
  chunks := []string{}
  i := 0
  chunk := ""
  for {
    if i >= len(lines) {
      chunks = append(chunks, chunk)
      break
    }
    if lines[i] == "" {
      chunks = append(chunks, chunk)
      chunk = ""
    } else {
      if chunk == "" {
        chunk = chunk + lines[i]
      } else {
        chunk = chunk + " " + lines[i]
      }
    }
    i++
  }
  return chunks, nil
}

// ReadChunksNoSpace is just like ReadChunks, but doesn't add a space between lines.
func ReadChunksNoSpace(path string) ([]string, error) {
  lines, err := ReadLines(path)
  if err != nil {
    return nil, err
  }
  chunks := []string{}
  i := 0
  chunk := ""
  for {
    if i >= len(lines) {
      chunks = append(chunks, chunk)
      break
    }
    if lines[i] == "" {
      chunks = append(chunks, chunk)
      chunk = ""
    } else {
      chunk = chunk + lines[i]
    }
    i++
  }
  return chunks, nil
}

// ReadChunksNoSpace is just like ReadChunks, but doesn't add a space between lines.
func ReadLinesInChunks(path string) ([][]string, error) {
  lines, err := ReadLines(path)
  if err != nil {
    return nil, err
  }
  chunks := [][]string{}
  i := 0
  chunk := []string{}
  for {
    if i >= len(lines) {
      chunks = append(chunks, chunk)
      break
    }
    if lines[i] == "" {
      chunks = append(chunks, chunk)
      chunk = []string{}
    } else {
      chunk = append(chunk, lines[i])
    }
    i++
  }
  return chunks, nil
}
