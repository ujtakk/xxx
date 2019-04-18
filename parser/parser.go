package parser

import (
  "io"
  "fmt"
  "bufio"
)

func Parse(file io.Reader) {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
