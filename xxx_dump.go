package main

import (
	// "fmt"
	"bufio"
	"math/bits"
	"os"
)

func dumpBin(data *XXXData, token *XXXToken) *XXXData {
	value := uint(0)
	for _, char := range token.body {
		value <<= 1
		switch char {
		case '0':
			value += 0x0
		case '1':
			value += 0x1
		default:
			panic("ERROR: dumpBin failed")
		}
	}

	return data.Join(value, 1*len(token.body))
}

func dumpOct(data *XXXData, token *XXXToken) *XXXData {
	value := uint(0)
	for _, char := range token.body {
		value <<= 3
		switch char {
		case '0':
			value += 0x0
		case '1':
			value += 0x1
		case '2':
			value += 0x2
		case '3':
			value += 0x3
		case '4':
			value += 0x4
		case '5':
			value += 0x5
		case '6':
			value += 0x6
		case '7':
			value += 0x7
		default:
			panic("ERROR: dumpBin failed")
		}
	}

	return data.Join(value, 3*len(token.body))
}

func dumpDec(data *XXXData, token *XXXToken) *XXXData {
	value := uint(0)
	for _, char := range token.body {
		value *= 10
		switch char {
		case '0':
			value += 0x0
		case '1':
			value += 0x1
		case '2':
			value += 0x2
		case '3':
			value += 0x3
		case '4':
			value += 0x4
		case '5':
			value += 0x5
		case '6':
			value += 0x6
		case '7':
			value += 0x7
		case '8':
			value += 0x8
		case '9':
			value += 0x9
		default:
			panic("ERROR: dumpBin failed")
		}
	}

	return data.Join(value, bits.Len(value))
}

func dumpHex(data *XXXData, token *XXXToken) *XXXData {
	value := uint(0)
	for _, char := range token.body {
		value <<= 4
		switch char {
		case '0':
			value += 0x0
		case '1':
			value += 0x1
		case '2':
			value += 0x2
		case '3':
			value += 0x3
		case '4':
			value += 0x4
		case '5':
			value += 0x5
		case '6':
			value += 0x6
		case '7':
			value += 0x7
		case '8':
			value += 0x8
		case '9':
			value += 0x9
		case 'a':
			fallthrough
		case 'A':
			value += 0xA
		case 'b':
			fallthrough
		case 'B':
			value += 0xB
		case 'c':
			fallthrough
		case 'C':
			value += 0xC
		case 'd':
			fallthrough
		case 'D':
			value += 0xD
		case 'e':
			fallthrough
		case 'E':
			value += 0xE
		case 'f':
			fallthrough
		case 'F':
			value += 0xF
		default:
			panic("ERROR: dumpBin failed")
		}
	}

	return data.Join(value, 4*len(token.body))
}

func dumpVar(data *XXXData, token *XXXToken, env *XXXEnv) *XXXData {
	value := env.Get(token.Compile())
	switch value.tag {
	case XXX_BIN:
		data = dumpBin(data, value)
	case XXX_OCT:
		data = dumpOct(data, value)
	case XXX_DEC:
		data = dumpDec(data, value)
	case XXX_HEX:
		data = dumpHex(data, value)
	case XXX_VAR:
		data = dumpVar(data, value, env)
	}

	return data
}

func dumpLine(tokens []*XXXToken, env *XXXEnv) *XXXData {
	data := NewData()

	for _, token := range tokens {
		switch token.tag {
		case XXX_BIN:
			data = dumpBin(data, token)
		case XXX_OCT:
			data = dumpOct(data, token)
		case XXX_DEC:
			data = dumpDec(data, token)
		case XXX_HEX:
			data = dumpHex(data, token)
		case XXX_VAR:
			data = dumpVar(data, token, env)
		default:
			panic("ERROR: invalid token tag")
		}
	}

	return data
}

func Dump(dst string, pool [][]*XXXToken, env *XXXEnv, little bool) {
	var file *os.File
	if dst == "" {
		file = os.Stdout
	} else {
		var err error
		file, err = os.Create(dst)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	writer := bufio.NewWriter(file)
	for _, tokens := range pool {
		data := dumpLine(tokens, env)
		_, err := writer.Write(data.Compile(little))
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
