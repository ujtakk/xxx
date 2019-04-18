package main

import (
  "io"
  "fmt"
  "bufio"
  "strings"
)

func parseLine(line string) {
  fmt.Println(line)
  r := strings.NewReader(line)
}

func Parse(file io.Reader) *XXXData {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    parseLine(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return new(XXXData)
}
