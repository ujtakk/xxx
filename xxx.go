package main

import (
  "os"
  "fmt"
  "flag"

  "./parser"
)

func usage(exitcode int) {
  // message := fmt.Sprintf(`usage: %v [-h] [-o {dst}] {src}`, os.Args[0])
  message := fmt.Sprintf(`usage: xxx [-h] [-o {dst}] {src}`)

  if exitcode == 0 {
    fmt.Println(message)
  } else {
    fmt.Fprintln(os.Stderr, message)
  }

  os.Exit(exitcode)
}

type Option struct {
  help bool
  dst string
  src string
}

func NewOption() *Option {
  opt := new(Option)

  flag.BoolVar(&opt.help, "h", false, "Show usage")
  flag.StringVar(&opt.dst, "o", "", "Specify output file")

  flag.Parse()
  if opt.help {
    usage(0)
  }

  opt.src = flag.Arg(0)
  if opt.src == "" {
    usage(1)
  }

  return opt
}

func main() {
  opt := NewOption()

  file, err := os.Open(opt.src)
  if err != nil {
    panic(err)
  }

  parser.Parse(file)
}
