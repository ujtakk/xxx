package main

import (
  "fmt"
  "io"
  "os"
  "bufio"
)

func isTokenSound(token XXXToken) bool {
  return true
}

func DumpVar(token XXXToken) XXXData {
  return nil
}

func DumpToken(token XXXToken, env *XXXEnv) XXXData {
  if !isTokenSound(token) {
    fmt.Fprintf(os.Stderr, "ERROR: lines have to be composed by bytes")
    os.Exit(1)
  }

  return nil
}

func Dump(dst string, pool []XXXToken, env *XXXEnv) {
  fmt.Println(pool)

  var file io.Writer
  if dst == "" {
    file = os.Stdout
  } else {
    file, err := os.Create(dst)
    if err != nil {
      panic(err)
    }
    defer file.Close()
  }

  writer := bufio.NewWriter(file)
  for _, token := range pool {
    data := DumpToken(token, env)
    _, err := writer.Write(data)
    if err != nil {
      panic(err)
    }
  }
}
