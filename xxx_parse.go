package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
  "unicode"
)

func parseDot(line *strings.Reader, env *XXXEnv) (XXXData, *XXXEnv, error) {
  ch, _, err := line.ReadRune()
  if !unicode.IsSpace(ch) {
    fmt.Fprintln(os.Stderr, "not valid use of dot command")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return nil, nil, nil
}

func parseVar(line *strings.Reader, env *XXXEnv) (XXXData, error) {
  ch, _, err := line.ReadRune()
  if !unicode.IsLetter(ch) {
    fmt.Fprintln(os.Stderr, "not valid use of dot command")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return nil, nil
}

func parseAssign(line *strings.Reader, env *XXXEnv) (*XXXEnv, error) {
  return env, nil
}

func parseNumber(line *strings.Reader) {
  ch, _, err := line.ReadRune()
  if !unicode.IsNumber(ch) {
    fmt.Fprintln(os.Stderr, "not valid format of number")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }
}

func parseLine(line *strings.Reader, env *XXXEnv) (XXXData, *XXXEnv) {
  data := make(XXXData, 0)

  loop: for {
    ch, _, err := line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsSpace(ch):
      continue
    case unicode.IsLetter(ch):
      line.UnreadRune()
      env, _ = parseAssign(line, env)
    case unicode.IsNumber(ch):
      line.UnreadRune()
      parseNumber(line)
    case ch == '.':
      parseDot(line, env)
    case ch == '$':
      val, err := parseVar(line, env)
      if err != nil {
        panic(err)
      }
      data = append(data, val...)
    case ch == '#':
      break loop
    case ch == '\x00':
      break loop
    default:
      fmt.Fprintln(os.Stderr, "not valid input")
      os.Exit(1)
    }
  }

  return data, env
}

func Parse(file io.Reader) []XXXData {
  env := new(XXXEnv)
  scanner := bufio.NewScanner(file)
  pool := make([]XXXData, 0)

  for scanner.Scan() {
    text := scanner.Text()
    line := strings.NewReader(text)

    data, env := parseLine(line, env)
    if data != nil {
      pool = append(pool, data)
    }

    fmt.Printf("%v\n%v\n\n", text, data)
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return pool
}
