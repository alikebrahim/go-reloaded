package goreloaded

import (
	"bytes"
	"fmt"
	"regexp"
)

// var slice [][]byte
var modIndxs []int
var ModMap = make(map[int][]byte)
var MapLen int

var QuotemodIndxs []int
var QuoteModMap = make(map[int][]byte)
var QuoteMapLen int

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

// func WrapQoute(i int, l *Lexer) []byte {
// 	slice = append(slice, []byte(l.TokenVals[i]))
// 	for j := i + 1; j < len(l.Tokens); j++ {
// 		if l.Tokens[j] == 3 {
// 			slice = append(slice, []byte(l.TokenVals[j]))
// 			break
// 		}
// 		if l.Tokens[j] == 4 {
// 			if l.Tokens[j-1] == 3 || l.Tokens[j+1] == 3 {
// 				continue
// 			}
// 		}
// 		slice = append(slice, []byte(l.TokenVals[j]))

// 	}

// 	return bytes.Join(slice, []byte(""))
// }

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

func QuoteLocateModifiers(l *Lexer) {
	for i, item := range l.Tokens {
		if item == 0 {
			QuotemodIndxs = append(QuotemodIndxs, i)
		}
	}
}

func QuoteModsMap(l *Lexer) {
	// a reset of the following to be done as a test -- test case for more than one Quotext:
	// QuotemodIndxs
	// QuoteModMap
	// QuoteMapLen
	QuoteLocateModifiers(l)
	for _, item := range QuotemodIndxs {
		forMod, mod := ModAnalyzer(l.TokenVals[item])
		for i := 0; i <= forMod; i++ {
			QuoteModMap[item-i] = mod
			QuoteMapLen++
		}
	}
}

func Power(a int, b int) int {
	if b == 0 {
		return 1
	}
	return a * Power(a, b-1)
}

func QuoteHandler(b []byte) []byte {
	quoteLexer := NewLexer(b)
	quoteLexer.QuoteFmtScan()
	QuoteModsMap(quoteLexer)
	ModEdit(&QuoteText, quoteLexer)
	fmt.Println("QuoteText", string(bytes.Join(QuoteText, []byte(""))))
	QuoteTextFmt(&QuoteFmtText)
	fmt.Println("QuoteFmtText", string(bytes.Join(QuoteText, []byte(""))))
	fmt.Println("***** Lexer @QuoteHandler *****")
	for i, token := range quoteLexer.Tokens {
		switch token {
		case Modifier:
			fmt.Printf("Modifier: %s\n", quoteLexer.TokenVals[i])
		case Identifier:
			fmt.Printf("Identifier: %s\n", quoteLexer.TokenVals[i])
		case Whitespace:
			fmt.Printf("Whitespace: %s\n", quoteLexer.TokenVals[i])
		case Punct:
			fmt.Printf("Punct: %s\n", quoteLexer.TokenVals[i])
		case Quotemark:
			fmt.Printf("Quotemark: %s\n", quoteLexer.TokenVals[i])
		case Invalid:
			fmt.Printf("Invalid: %s\n", quoteLexer.TokenVals[i])
		}
	}
	return bytes.Join(QuoteFmtText, []byte(""))
}
