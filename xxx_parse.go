package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
  "unicode"
)

func parseImport(line *strings.Reader, env *XXXEnv) (XXXData, *XXXEnv, error) {
  ch, _, err := line.ReadRune()
  if !unicode.IsSpace(ch) {
    fmt.Fprintln(os.Stderr, "not valid use of import command (.)")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return nil, nil, nil
}

func parseAssign(line *strings.Reader, env *XXXEnv) (*XXXEnv, error) {
  ch, _, err := line.ReadRune()
  if !unicode.IsSpace(ch) {
    fmt.Fprintln(os.Stderr, "not valid use of assign command (=)")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return env, nil
}

func parseNum(line *strings.Reader) (XXXData, error) {
  ch, _, _ := line.ReadRune()
  if ch == '0' {
  } else {
  }

  return nil, nil
}

func parseVar(line *strings.Reader, env *XXXEnv) (XXXData, error) {
  return nil, nil
}

func parseLine(line *strings.Reader, env *XXXEnv) (XXXData, *XXXEnv) {
  var ch rune
  var err error

  data := make(XXXData, 0)

  ch, _, err = line.ReadRune()
  switch ch {
  case '.':
    parseImport(line, env)
  case '=':
    parseAssign(line, env)
  }
  line.UnreadRune()

loop:
  for {
    ch, _, err = line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsSpace(ch):
      continue
    case unicode.IsNumber(ch):
      line.UnreadRune()
      data, _ = parseNum(line)
    case unicode.IsLetter(ch):
      line.UnreadRune()
      data, _ = parseVar(line, env)
    case ch == '#':
      break loop
    case ch == '\x00':
      break loop
    case ch == '.':
      fmt.Fprintln(os.Stderr,
        "import command (.) have to be issued at start of the line")
      os.Exit(1)
    case ch == '=':
      fmt.Fprintln(os.Stderr,
        "assign command (=) have to be issued at start of the line")
      os.Exit(1)
    default:
      fmt.Fprintf(os.Stderr, "%v is not valid input\n", ch)
      os.Exit(1)
    }
  }

  return data, env
}

func Parse(file io.Reader) []XXXData {
  var data XXXData
  var env *XXXEnv = new(XXXEnv)
  var pool []XXXData = make([]XXXData, 0)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    line := strings.NewReader(text)

    data, env = parseLine(line, env)
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
