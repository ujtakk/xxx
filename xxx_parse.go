package main

import (
  "io"
  "fmt"
  "bufio"
  "strings"
)

func parseLine(line *strings.Reader) *XXXData {
  stack := new(XXXData)

  loop: for {
    ch, _, err := line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch ch {
    case ' ':
      continue
    case '#':
      break loop
    case '\x00':
      break loop
    }

    *stack = append(*stack, ch)
  }
  fmt.Printf("%v\n\n", stack)

  return stack
}

func Parse(file io.Reader) *XXXData {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    line := strings.NewReader(text)

    fmt.Println(text)
    parseLine(line)
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return new(XXXData)
}
