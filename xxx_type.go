package main

type XXXVar struct{
  name string
  value []byte
}

type XXXEnv []XXXVar

func (e *XXXEnv) Add(name string, value []byte) {
  *e = append(*e, XXXVar{name, value})
}

type XXXData []byte
