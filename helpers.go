package goreloaded

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"regexp"
)

var slice [][]byte

func ModAnalyzer(mod []byte) int {
	reIndx := regexp.MustCompile(`\d+`)
	match := reIndx.Find(mod)
	if match == nil {
		return 1
	}
	NumOfIdens := binary.BigEndian.Uint64(match)
	NumOfIdens_int := int(NumOfIdens)
	fmt.Println(NumOfIdens_int)
	return NumOfIdens_int * 2
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
