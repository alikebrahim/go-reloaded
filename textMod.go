package goreloaded

import (
	"bytes"
	"fmt"
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

// for i, item := range lexer2.Tokens {
// 	if i+1 < len(lexer2.Tokens) && item == 4 {
// 		if lexer2.Tokens[i+1] == 2 {
// 			continue
// 		}
// 	}
// 	if i+1 < len(lexer2.Tokens) && item == 2 {
// 		if lexer2.Tokens[i+1] != 4 {
// 			var slice [][]byte
// 			slice = append(slice, []byte(lexer2.TokenVals[i]))
// 			slice = append(slice, []byte(" "))
// 			jbs := bytes.Join(slice, []byte(""))
// 			finalText = append(finalText, jbs)
// 			continue
// 		}

// 	}
// 	if i+1 < len(lexer2.Tokens) && item == 3 {
// 		wrappedText := goreloaded.WrapQoute(i, lexer2)
// 		finalText = append(finalText, wrappedText)
// 		break
// 	}

// 	if item == 4 && len(lexer2.TokenVals[i]) != 1 {
// 		finalText = append(finalText, []byte(" "))
// 		continue
// 	}
// 	if len(lexer2.TokenVals[i]) == 1 && lexer2.TokenVals[i][0] == 'a' {
// 		if lexer2.TokenVals[i+2][0] == 'a' || lexer2.TokenVals[i+2][0] == 'e' || lexer2.TokenVals[i+2][0] == 'i' || lexer2.TokenVals[i+2][0] == 'o' || lexer2.TokenVals[i+2][0] == 'u' || lexer2.TokenVals[i+2][0] == 'h' {
// 			finalText = append(finalText, []byte("an"))
// 			continue
// 		}
// 	}
// 	if len(lexer2.TokenVals[i]) == 1 && lexer2.TokenVals[i][0] == 'A' {
// 		if lexer2.TokenVals[i+2][0] == 'a' || lexer2.TokenVals[i+2][0] == 'e' || lexer2.TokenVals[i+2][0] == 'i' || lexer2.TokenVals[i+2][0] == 'o' || lexer2.TokenVals[i+2][0] == 'u' || lexer2.TokenVals[i+2][0] == 'h' {
// 			finalText = append(finalText, []byte("An"))
// 			continue
// 		}
// 	}
// 	finalText = append(finalText, []byte(lexer2.TokenVals[i]))
// }

func TextFmt(tb *[][]byte) {
	fmtlexer := NewLexer(bytes.Join(Text, []byte("")))
	fmtlexer.Scan()

	for i, token := range fmtlexer.Tokens {
		switch token {
		case Modifier:
			fmt.Printf("Modifier: %s\n", fmtlexer.TokenVals[i])
		case Identifier:
			fmt.Printf("Identifier: %s\n", fmtlexer.TokenVals[i])
		case Whitespace:
			fmt.Printf("Whitespace: %s\n", fmtlexer.TokenVals[i])
		case Punct:
			fmt.Printf("Punct: %s\n", fmtlexer.TokenVals[i])
		case Quotemark:
			fmt.Printf("Quotemark: %s\n", fmtlexer.TokenVals[i])
		case Invalid:
			fmt.Printf("Invalid: %s\n", fmtlexer.TokenVals[i])
		}
	}
}
