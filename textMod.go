package goreloaded

import (
	"bytes"
)

var Text [][]byte
var FmtText [][]byte

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

func TextFmt(tb *[][]byte) {
	fmtlexer := NewLexer(bytes.Join(Text, []byte("")))
	fmtlexer.Scan()
	for i, item := range fmtlexer.Tokens {
		if i+1 < len(fmtlexer.Tokens) && item == 4 {
			if fmtlexer.Tokens[i+1] == 2 {
				continue
			}
		}
		if i+1 < len(fmtlexer.Tokens) && item == 2 {
			if fmtlexer.Tokens[i+1] != 4 {
				var slice [][]byte
				slice = append(slice, []byte(fmtlexer.TokenVals[i]))
				slice = append(slice, []byte(" "))
				jbs := bytes.Join(slice, []byte(""))
				*tb = append(*tb, jbs)
				continue
			}

		}
		//quotedit
		if i+1 < len(fmtlexer.Tokens) && item == 3 {
			firstIndex := bytes.IndexByte(fmtlexer.TokenVals[i], '\'')
			lastIndex := bytes.LastIndexByte(fmtlexer.TokenVals[i], '\'')
			trimmedSlice := bytes.TrimSpace(fmtlexer.TokenVals[i][firstIndex+1 : lastIndex])
			*tb = append(*tb, []byte("'"))
			*tb = append(*tb, []byte(trimmedSlice))
			*tb = append(*tb, []byte("'"))

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

	// for i, token := range fmtlexer.Tokens {
	// 	switch token {
	// 	case Modifier:
	// 		fmt.Printf("Modifier: %s\n", fmtlexer.TokenVals[i])
	// 	case Identifier:
	// 		fmt.Printf("Identifier: %s\n", fmtlexer.TokenVals[i])
	// 	case Whitespace:
	// 		fmt.Printf("Whitespace: %s\n", fmtlexer.TokenVals[i])
	// 	case Punct:
	// 		fmt.Printf("Punct: %s\n", fmtlexer.TokenVals[i])
	// 	case Quotedtext:
	// 		fmt.Printf("Quotedtext: %s\n", fmtlexer.TokenVals[i])
	// 	case Invalid:
	// 		fmt.Printf("Invalid: %s\n", fmtlexer.TokenVals[i])
	// 	}
	// }
}
