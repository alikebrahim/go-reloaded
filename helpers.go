package goreloaded

import (
	"bytes"
	"regexp"
)

var slice [][]byte
var modIndxs []int
var ModMap = make(map[int][]byte)
var MapLen int

func ModAnalyzer(mod []byte) (int, []byte) {
	reMod := regexp.MustCompile(`[hex|bin|cap|low|up]`)
	reIndx := regexp.MustCompile(`\d+`)
	modMatch := reMod.Find(mod)
	indxMatch := reIndx.Find(mod)
	if indxMatch == nil {
		return 2, modMatch
	}
	return int(indxMatch[0]-48) * 2, modMatch
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

func LocateModifiers(l *Lexer) {
	for i, item := range l.Tokens {
		if item == 0 {
			modIndxs = append(modIndxs, i)
		}
	}
}

func ModsMap(l *Lexer) {
	LocateModifiers(l)
	for _, item := range modIndxs {
		forMod, mod := ModAnalyzer(l.TokenVals[item])
		for i := 0; i <= forMod; i++ {
			ModMap[item-i] = mod
			MapLen++
		}
	}
}
