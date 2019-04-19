package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
  "unicode"
)

func parseNum(line *strings.Reader) XXXToken {
  token := NewToken()

  ch, _, err := line.ReadRune()
  token = token.AddRune(ch)
  fmt.Printf("%1v %2x %x\n", string(ch), ch, token)

  if ch == '0' {
    ch, _, _ := line.ReadRune()
    switch {
    case ch == 'b':
    case ch == 'x':
    case unicode.IsNumber(ch):
    default:
      fmt.Fprintln(os.Stderr, "ERROR: not valid number format (2, 8, 10, 16)")
      os.Exit(1)
    }
    token = token.AddRune(ch)
    fmt.Printf("%1v %2x %x\n", string(ch), ch, token)
  }

loop:
  for {
    ch, _, err = line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsNumber(ch):
    case unicode.IsSpace(ch):
      break loop
    default:
      fmt.Fprintln(os.Stderr, "ERROR: not valid number format (2, 8, 10, 16)")
      os.Exit(1)
    }
    token = token.AddRune(ch)
    fmt.Printf("%1v %2x %x\n", string(ch), ch, token)
  }

  fmt.Println()
  return token
}

func parseVar(line *strings.Reader) XXXToken {
  token := NewToken()

loop:
  for {
    ch, _, err := line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsLetter(ch):
    case unicode.IsSpace(ch):
      break loop
    default:
      fmt.Fprintln(os.Stderr, "ERROR: not valid number format (2, 8, 10, 16)")
      os.Exit(1)
    }
    token.AddRune(ch)
  }

  return token
}

func parseAssign(line *strings.Reader, env *XXXEnv) *XXXEnv {
  var name string
  var token XXXToken

  ch, _, err := line.ReadRune()
  if !unicode.IsSpace(ch) {
    fmt.Fprintln(os.Stderr, "ERROR: commands have to be separated by spaces")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

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
      token = parseNum(line)
      env.Add(name, token)
      break loop
    case unicode.IsLetter(ch):
      line.UnreadRune()
      token = parseVar(line)
      env.Add(name, token)
      break loop
    }
  }

  return env
}

func parseImport(line *strings.Reader, env *XXXEnv) ([]XXXToken, *XXXEnv) {
  ch, _, err := line.ReadRune()
  if !unicode.IsSpace(ch) {
    fmt.Fprintln(os.Stderr, "ERROR: commands have to be separated by spaces")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return nil, env
}

func parseLine(line *strings.Reader, env *XXXEnv) ([]XXXToken, *XXXEnv) {
  var ch rune
  var err error

  tokens := make([]XXXToken, 0)

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
      token := parseNum(line)
      tokens = append(tokens, token)
    case unicode.IsLetter(ch):
      line.UnreadRune()
      token := parseVar(line)
      tokens = append(tokens, token)
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

  return tokens, env
}

func Parse(src string) ([]XXXToken, *XXXEnv) {
  var token []XXXToken
  var env *XXXEnv = NewEnv()
  var pool []XXXToken = make([]XXXToken, 0)

  file, err := os.Open(src)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    line := strings.NewReader(text)

    token, env = parseLine(line, env)
    if token != nil {
      pool = append(pool, token...)
    }

    fmt.Printf("%v\n%v\n\n", text, token)
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return pool, env
}
