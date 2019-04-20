package main

import "fmt"

const BYTE = 8

type XXXType int

const (
  XXX_VAR = iota
  XXX_BIN
  XXX_OCT
  XXX_DEC
  XXX_HEX
)

type XXXToken struct{
  tag XXXType
  body []rune
}

func NewToken() *XXXToken {
  t := new(XXXToken)
  t.body = make([]rune, 0)
  return t
}

func (t *XXXToken) Add(r rune) *XXXToken {
  t.body = append(t.body, r)
  return t
}

type XXXData struct{
  capacity uint
  body []byte
}

func NewData() *XXXData {
  d := new(XXXData)
  d.body = make([]byte, 0)
  d.capacity = 0
  return d
}

func (d *XXXData) Join(s uint, l int) *XXXData {
  if d.capacity == 0 {
    d.body = append(d.body, 0x0)
    d.capacity = BYTE
  }

  index := len(d.body) - 1
  length := uint(l)
  if d.capacity < length {
    d.body[index] = byte((s >> (length-d.capacity)) & 0xFF) | d.body[index]
    for length -= d.capacity; length > BYTE; length -= BYTE {
      d.body = append(d.body, byte((s >> (length-BYTE)) & 0xFF))
      index++
    }
    if length > 0 {
      d.body = append(d.body, byte((s << (BYTE-length)) & 0xFF))
      d.capacity = BYTE - length
    } else {
      d.capacity = 0
    }
  } else {
    d.body[index] = byte((s << (d.capacity-length)) & 0xFF) | d.body[index]
    d.capacity -= length
  }
  fmt.Printf("%08b (%2d: %2d) -> (%2d, %2d)%08b\n", s, s, l, index, d.capacity, d.body)

  return d
}

type XXXEnv map[string]*XXXToken

func NewEnv() *XXXEnv {
  return new(XXXEnv)
}

func (e *XXXEnv) Set(name string, value *XXXToken) {
  (*e)[name] = value
}
