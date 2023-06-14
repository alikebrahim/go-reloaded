package goreloaded

import (
	"bytes"
	"regexp"
	"strconv"
)

var slice [][]byte

func ModAnalyzer(mod string) int {
	reArg := regexp.MustCompile(`\d+`)
	match := reArg.FindString(mod)
	if match == "" {
		return 1
	}
	NumOfIdens, _ := strconv.Atoi(match)
	return NumOfIdens * 2
}

func WrapQoute(i int, l *Lexer) []byte {
	slice = append(slice, []byte(l.TokenVals[i]))
	for j := i + 1; j < len(l.Tokens); j++ {
		if l.Tokens[j] == 3 {
			slice = append(slice, []byte(l.TokenVals[j]))
			break
		}
		if l.Tokens[j] == 4 {
			if l.Tokens[j-1] == 3 || l.Tokens[j+1] == 3 {
				continue
			}
		}
		slice = append(slice, []byte(l.TokenVals[j]))

	}

	return bytes.Join(slice, []byte(""))
}
