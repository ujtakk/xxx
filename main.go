package main

import (
  "os"
  "fmt"
  "flag"
  "bufio"
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
  dst_file string
  src_file string
}

func NewOption() *Option {
  opt := new(Option)

  flag.BoolVar(&opt.help, "h", false, "Show usage")
  flag.StringVar(&opt.dst_file, "o", "", "Specify output file")

  flag.Parse()
  if opt.help {
    usage(0)
  }

  opt.src_file = flag.Arg(0)
  if opt.src_file == "" {
    usage(1)
  }

  return opt
}

func main() {
  opt := NewOption()

  file, err := os.Open(opt.src_file)
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
