package goreloaded

import (
	"bytes"
)

var Text [][]byte
var FmtText [][]byte

var QuoteText [][]byte
var QuoteFmtText [][]byte

// Modify words
// (hex|bin|cap|up|low) cases
func ModEdit(t *[][]byte, l *Lexer) {
	for i, item := range l.Tokens {
		if ModMap[i] != nil && item != 0 {
			mod := string(ModMap[i])
			if mod == "c" {
				*t = append(*t, Cap(l.TokenVals[i]))
				continue
			} else if mod == "u" {
				*t = append(*t, Up(l.TokenVals[i]))
				continue
			} else if mod == "l" {
				*t = append(*t, Low(l.TokenVals[i]))
				continue
			} else if mod == "h" {
				*t = append(*t, Hex(l.TokenVals[i]))
				continue
			} else if mod == "b" {
				*t = append(*t, Bin(l.TokenVals[i]))
				continue
			}
		}
		if item == 0 {
			continue
		}
		*t = append(*t, l.TokenVals[i])
	}

}

// Format text
// Punctuation, vowel case and quotes handling
func TextFmt(tb *[][]byte) {
	fmtlexer := NewLexer(bytes.Join(Text, []byte("")))
	fmtlexer.Scan()
	for i, item := range fmtlexer.Tokens {
		if item == 4 {
			if fmtlexer.Tokens[i+1] == 2 {
				continue
			}
		}
		if item == 2 {
			if fmtlexer.Tokens[i+1] != 4 {
				var slice [][]byte
				slice = append(slice, []byte(fmtlexer.TokenVals[i]))
				slice = append(slice, []byte(" "))
				jbs := bytes.Join(slice, []byte(""))
				*tb = append(*tb, jbs)
				continue
			}

		}
		if item == 3 {
			*tb = append(*tb, QuoteHandler(fmtlexer.TokenVals[i]))
			continue
		}

		if item == 4 && len(fmtlexer.TokenVals[i]) != 1 {
			*tb = append(*tb, []byte(" "))
			continue
		}
		if len(fmtlexer.TokenVals[i]) == 1 && fmtlexer.TokenVals[i][0] == 'a' {
			if fmtlexer.TokenVals[i+2][0] == 'a' || fmtlexer.TokenVals[i+2][0] == 'e' || fmtlexer.TokenVals[i+2][0] == 'i' || fmtlexer.TokenVals[i+2][0] == 'o' || fmtlexer.TokenVals[i+2][0] == 'u' || fmtlexer.TokenVals[i+2][0] == 'h' {
				*tb = append(*tb, []byte("an"))
				continue
			}
		}
		if len(fmtlexer.TokenVals[i]) == 1 && fmtlexer.TokenVals[i][0] == 'A' {
			if fmtlexer.TokenVals[i+2][0] == 'a' || fmtlexer.TokenVals[i+2][0] == 'e' || fmtlexer.TokenVals[i+2][0] == 'i' || fmtlexer.TokenVals[i+2][0] == 'o' || fmtlexer.TokenVals[i+2][0] == 'u' || fmtlexer.TokenVals[i+2][0] == 'h' {
				*tb = append(*tb, []byte("An"))
				continue
			}
		}
		*tb = append(*tb, []byte(fmtlexer.TokenVals[i]))
	}

}

// Inner quote handling (Quotetext)
func QuoteModEdit(t *[][]byte, l *Lexer) {
	for i, item := range l.Tokens {
		if QuoteModMap[i] != nil && item != 0 {
			mod := string(QuoteModMap[i])
			if mod == "c" {
				*t = append(*t, Cap(l.TokenVals[i]))
				continue
			} else if mod == "u" {
				*t = append(*t, Up(l.TokenVals[i]))
				continue
			} else if mod == "l" {
				*t = append(*t, Low(l.TokenVals[i]))
				continue
			} else if mod == "h" {
				*t = append(*t, Hex(l.TokenVals[i]))
				continue
			} else if mod == "b" {
				*t = append(*t, Bin(l.TokenVals[i]))
				continue
			}
		}
		if item == 0 {
			continue
		}
		*t = append(*t, l.TokenVals[i])
	}

}

func QuoteTextFmt(tb *[][]byte) {
	fmtlexer := NewLexer(bytes.Join(QuoteText, []byte("")))
	fmtlexer.QuoteFmtScan()
	for i, item := range fmtlexer.Tokens {
		if item == 4 {
			if fmtlexer.Tokens[i+1] == 2 {
				continue
			} else if fmtlexer.Tokens[i+1] == 5 {
				continue
			}
		}
		if item == 4 {
			if fmtlexer.Tokens[i-1] == 5 {
				continue
			}
		}
		if item == 2 {
			if fmtlexer.Tokens[i+1] != 4 {
				var slice [][]byte
				slice = append(slice, []byte(fmtlexer.TokenVals[i]))
				slice = append(slice, []byte(" "))
				jbs := bytes.Join(slice, []byte(""))
				*tb = append(*tb, jbs)
				continue
			}

		}

		if item == 4 && len(fmtlexer.TokenVals[i]) != 1 {
			*tb = append(*tb, []byte(" "))
			continue
		}
		if len(fmtlexer.TokenVals[i]) == 1 && fmtlexer.TokenVals[i][0] == 'a' {
			if fmtlexer.TokenVals[i+2][0] == 'a' || fmtlexer.TokenVals[i+2][0] == 'e' || fmtlexer.TokenVals[i+2][0] == 'i' || fmtlexer.TokenVals[i+2][0] == 'o' || fmtlexer.TokenVals[i+2][0] == 'u' || fmtlexer.TokenVals[i+2][0] == 'h' {
				*tb = append(*tb, []byte("an"))
				continue
			}
		}
		if len(fmtlexer.TokenVals[i]) == 1 && fmtlexer.TokenVals[i][0] == 'A' {
			if fmtlexer.TokenVals[i+2][0] == 'a' || fmtlexer.TokenVals[i+2][0] == 'e' || fmtlexer.TokenVals[i+2][0] == 'i' || fmtlexer.TokenVals[i+2][0] == 'o' || fmtlexer.TokenVals[i+2][0] == 'u' || fmtlexer.TokenVals[i+2][0] == 'h' {
				*tb = append(*tb, []byte("An"))
				continue
			}
		}
		*tb = append(*tb, []byte(fmtlexer.TokenVals[i]))
	}

}
