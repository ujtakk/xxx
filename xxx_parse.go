package main

import (
	"bufio"
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
		case unicode.IsNumber(char):
			token.tag = XXX_OCT
			token = token.Add(char)
		case unicode.IsSpace(char):
			token.tag = XXX_DEC
			token = token.Add('0')
			return token
		default:
			panic("ERROR: not valid number format (2, 8, 10, 16)")
		}
	} else {
		token.tag = XXX_DEC
		token = token.Add(char)
	}

loop:
	for {
		char, _, err = line.ReadRune()
		if err == io.EOF || unicode.IsSpace(char) {
			break loop
		} else if err != nil {
			panic(err)
		}

		switch {
		case unicode.IsNumber(char):
			token = token.Add(char)
		case char == '_':
		default:
			panic("ERROR: not valid number format (2, 8, 10, 16)")
		}
	}

	return token
}

func parseVar(line *strings.Reader) *XXXToken {
	token := NewToken()
	token.tag = XXX_VAR

loop:
	for {
		char, _, err := line.ReadRune()
		if err == io.EOF || unicode.IsSpace(char) {
			break loop
		} else if err != nil {
			panic(err)
		}
		token = token.Add(char)
	}

	return token
}

func parseAssign(line *strings.Reader, env *XXXEnv) *XXXEnv {
	var name string
	var token *XXXToken

	char, _, err := line.ReadRune()
	if !unicode.IsSpace(char) {
		panic("ERROR: commands have to be separated by spaces")
	} else if err != nil {
		panic(err)
	}

	for unicode.IsSpace(char) {
		char, _, err = line.ReadRune()
		if err != nil {
			panic(err)
		}
	}

	if unicode.IsLetter(char) {
		line.UnreadRune()
		name = parseVar(line).Compile()
	} else {
		panic("ERROR: not valid format for assignment")
	}

	char, _, err = line.ReadRune()
	for unicode.IsSpace(char) {
		char, _, err = line.ReadRune()
		if err != nil {
			panic(err)
		}
	}

	switch {
	case unicode.IsLetter(char):
		line.UnreadRune()
		token = parseVar(line)
	case unicode.IsNumber(char):
		line.UnreadRune()
		token = parseNum(line)
	default:
		panic("ERROR: not valid format for assignment")
	}

	env = env.Set(name, token)

	return env
}

func parseImport(line *strings.Reader, env *XXXEnv) ([][]*XXXToken, *XXXEnv) {
	char, _, err := line.ReadRune()
	if !unicode.IsSpace(char) {
		panic("ERROR: commands have to be separated by spaces")
	} else if err != nil {
		panic(err)
	}

	for unicode.IsSpace(char) {
		char, _, err = line.ReadRune()
		if err != nil {
			panic(err)
		}
	}

	filepath := make([]rune, 0)
	for !unicode.IsSpace(char) && err != io.EOF {
		filepath = append(filepath, char)
		char, _, err = line.ReadRune()
	}

	xpool, xenv := Parse(string(filepath))

	return xpool, xenv
}

func parseCommand(line *strings.Reader, env *XXXEnv) ([][]*XXXToken, *XXXEnv) {
	char, _, err := line.ReadRune()
	switch char {
	case '.':
		xpool, xenv := parseImport(line, env)
		return xpool, xenv
	case '=':
		env = parseAssign(line, env)
		return nil, env
	default:
		if err == io.EOF {
			return nil, nil
		} else if err != nil {
			panic(err)
		}
		line.UnreadRune()
	}

	return nil, nil
}

func parseLine(line *strings.Reader, env *XXXEnv) ([]*XXXToken, *XXXEnv) {
	var char rune
	var err error

	tokens := make([]*XXXToken, 0)

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
		default:
			panic("ERROR: " + string(char) + " is not valid input")
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

		var xpool [][]*XXXToken
		var xenv *XXXEnv
		xpool, xenv = parseCommand(line, env)
		if xpool != nil || xenv != nil {
			if xpool != nil {
				pool = append(pool, xpool...)
			}
			if xenv != nil {
				env = env.Concat(xenv)
			}
			continue
		}

		tokens, env = parseLine(line, env)
		if len(tokens) > 0 {
			pool = append(pool, tokens)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return pool, env
}
