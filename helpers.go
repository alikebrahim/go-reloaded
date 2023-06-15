package goreloaded

import (
	"bytes"
	"regexp"
)

var slice [][]byte

func ModAnalyzer(mod []byte) int {
	reIndx := regexp.MustCompile(`\d+`)
	match := reIndx.Find(mod)
	if match == nil {
		return 1
	}
	return int(match[0]-48) * 2
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
