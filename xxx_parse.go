package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
  "unicode"
)

func parseNum(line *strings.Reader) *XXXToken {
  token := NewToken()

  char, _, err := line.ReadRune()

  if char == '0' {
    char, _, _ := line.ReadRune()
    switch {
    case char == 'b':
      token.tag = XXX_BIN
    case char == 'x':
      token.tag = XXX_HEX
    case unicode.IsSpace(char):
      token.tag = XXX_DEC
      token = token.Add('0')
      return token
    case unicode.IsNumber(char):
      token.tag = XXX_OCT
      token = token.Add(char)
    default:
      fmt.Fprintln(os.Stderr, "ERROR: not valid number format (2, 8, 10, 16)")
      os.Exit(1)
    }
  } else {
    token.tag = XXX_DEC
    token = token.Add(char)
  }

loop:
  for {
    char, _, err = line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    if unicode.IsSpace(char) {
      break loop
    }

    token = token.Add(char)
    fmt.Printf("%1v %2x %x\n", string(char), char, token)
  }

  fmt.Println()
  return token
}

func parseVar(line *strings.Reader) *XXXToken {
  token := NewToken()

loop:
  for {
    char, _, err := line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsLetter(char):
    case unicode.IsSpace(char):
      break loop
    default:
      fmt.Fprintln(os.Stderr, "ERROR: not valid number format (2, 8, 10, 16)")
      os.Exit(1)
    }
    token.Add(char)
  }

  return token
}

func parseAssign(line *strings.Reader, env *XXXEnv) *XXXEnv {
  var name string
  var token *XXXToken

  char, _, err := line.ReadRune()
  if !unicode.IsSpace(char) {
    fmt.Fprintln(os.Stderr, "ERROR: commands have to be separated by spaces")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

loop:
  for {
    char, _, err = line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsSpace(char):
      continue
    case unicode.IsNumber(char):
      line.UnreadRune()
      token = parseNum(line)
      env.Set(name, token)
      break loop
    case unicode.IsLetter(char):
      line.UnreadRune()
      token = parseVar(line)
      env.Set(name, token)
      break loop
    }
  }

  return env
}

func parseImport(line *strings.Reader, env *XXXEnv) ([]XXXToken, *XXXEnv) {
  char, _, err := line.ReadRune()
  if !unicode.IsSpace(char) {
    fmt.Fprintln(os.Stderr, "ERROR: commands have to be separated by spaces")
    os.Exit(1)
  } else if err != nil {
    panic(err)
  }

  return nil, env
}

func parseLine(line *strings.Reader, env *XXXEnv) ([]*XXXToken, *XXXEnv) {
  var char rune
  var err error

  tokens := make([]*XXXToken, 0)

  char, _, err = line.ReadRune()
  switch char {
  case '.':
    parseImport(line, env)
  case '=':
    parseAssign(line, env)
  }
  line.UnreadRune()

loop:
  for {
    char, _, err = line.ReadRune()
    if err == io.EOF {
      break loop
    } else if err != nil {
      panic(err)
    }

    switch {
    case unicode.IsSpace(char):
      continue
    case unicode.IsNumber(char):
      line.UnreadRune()
      token := parseNum(line)
      tokens = append(tokens, token)
    case unicode.IsLetter(char):
      line.UnreadRune()
      token := parseVar(line)
      tokens = append(tokens, token)
    case char == '#':
      break loop
    case char == '\x00':
      break loop
    case char == '.':
      fmt.Fprintln(os.Stderr,
        "import command (.) have to be issued at start of the line")
      os.Exit(1)
    case char == '=':
      fmt.Fprintln(os.Stderr,
        "assign command (=) have to be issued at start of the line")
      os.Exit(1)
    default:
      fmt.Fprintf(os.Stderr, "%v is not valid input\n", char)
      os.Exit(1)
    }
  }

  return tokens, env
}

func Parse(src string) ([][]*XXXToken, *XXXEnv) {
  var tokens []*XXXToken
  var pool [][]*XXXToken = make([][]*XXXToken, 0)
  var env *XXXEnv = NewEnv()

  var file *os.File
  if src == "" {
    file = os.Stdout
  } else {
    var err error
    file, err = os.Open(src)
    if err != nil {
      panic(err)
    }
    defer file.Close()
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    line := strings.NewReader(text)
    fmt.Printf("%v\n", text)

    tokens, env = parseLine(line, env)
    if len(tokens) > 0 {
      pool = append(pool, tokens)
      fmt.Printf("%v\n\n", tokens)
    }
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return pool, env
}
