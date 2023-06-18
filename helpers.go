package goreloaded

import (
	"bytes"
	"regexp"
)

var modIndxs []int
var ModMap = make(map[int][]byte)
var MapLen int

var QuotemodIndxs []int
var QuoteModMap = make(map[int][]byte)
var QuoteMapLen int

// Analyze modifiers and identify n of words to be modified
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

// Map the index of each modifier
func LocateModifiers(l *Lexer) {
	for i, item := range l.Tokens {
		if item == 0 {
			modIndxs = append(modIndxs, i)
		}
	}
}

// Locate the indicies of the words to be modified
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

// Inner quote handling (Quotetext)
func QuoteLocateModifiers(l *Lexer) {
	for i, item := range l.Tokens {
		if item == 0 {
			QuotemodIndxs = append(QuotemodIndxs, i)
		}
	}
}

func QuoteModsMap(l *Lexer) {
	QuoteLocateModifiers(l)
	for _, item := range QuotemodIndxs {
		forMod, mod := ModAnalyzer(l.TokenVals[item])
		for i := 0; i <= forMod; i++ {
			QuoteModMap[item-i] = mod
			QuoteMapLen++
		}
	}
}

func QuoteHandler(b []byte) []byte {
	// Initialize and scan lexer
	quoteLexer := NewLexer(b)
	quoteLexer.QuoteFmtScan()

	// Edit and format for qouted text
	QuoteModsMap(quoteLexer)
	QuoteModEdit(&QuoteText, quoteLexer)
	QuoteTextFmt(&QuoteFmtText)

	return bytes.Join(QuoteFmtText, []byte(""))
}

func Power(a int, b int) int {
	if b == 0 {
		return 1
	}
	return a * Power(a, b-1)
}
