package main

type XXXVar struct{
  name string
  value []rune
}

type XXXEnv []XXXVar

func (e *XXXEnv) Add(name string, value []rune) {
  *e = append(*e, XXXVar{name, value})
}

type XXXData []rune
