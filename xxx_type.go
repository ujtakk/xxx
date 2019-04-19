package main

type XXXToken []rune

type XXXData []byte

type XXXEnv map[string]XXXToken

func NewToken() XXXToken {
  return make(XXXToken, 0)
}

func (t XXXToken) AddRune(r rune) XXXToken {
  return append(t, r)
}

func NewEnv() *XXXEnv {
  return new(XXXEnv)
}

func (e *XXXEnv) Add(name string, value []rune) {
  (*e)[name] = value
}
