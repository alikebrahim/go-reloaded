package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Token int

const (
	Modifier Token = iota
	Identifier
	Punct
	Quotemark
	Whitespace
	Invalid
)

type Lexer struct {
	input     string
	position  int
	tokens    []Token
	tokenVals []string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:     input,
		position:  0,
		tokens:    []Token{},
		tokenVals: []string{},
	}
}

func (l *Lexer) Scan() {
	reIdentifier := regexp.MustCompile(`[a-zA-Z0-9_]\w*`)
	reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)
	rePunct := regexp.MustCompile(`[.,?!:;]+`)
	reQuotemark := regexp.MustCompile(`['|"]`)
	reWhitespace := regexp.MustCompile(`\s+`)

	for l.position < len(l.input) {
		substr := l.input[l.position:]

		if match := reIdentifier.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.tokens = append(l.tokens, Identifier)
			l.tokenVals = append(l.tokenVals, reIdentifier.FindString(substr))
			l.position += match[1]
		} else if match := reModifier.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.tokens = append(l.tokens, Modifier)
			l.tokenVals = append(l.tokenVals, reModifier.FindString(substr))
			l.position += match[1]
		} else if match := reWhitespace.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.tokens = append(l.tokens, Whitespace)
			l.tokenVals = append(l.tokenVals, reWhitespace.FindString(substr))
			l.position += match[1]
		} else if match := rePunct.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.tokens = append(l.tokens, Punct)
			l.tokenVals = append(l.tokenVals, rePunct.FindString(substr))
			l.position += match[1]
		} else if match := reQuotemark.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.tokens = append(l.tokens, Quotemark)
			l.tokenVals = append(l.tokenVals, reQuotemark.FindString(substr))
			l.position += match[1]
		} else {
			l.tokens = append(l.tokens, Invalid)
			l.tokenVals = append(l.tokenVals, string(substr[0]))
			l.position++
		}
	}
}

func main() {
	// ensure proper usage
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Usage requires input.txt and output.txt")
		return
	}
	if os.Args[1][len(os.Args[1])-4:] != ".txt" || os.Args[2][len(os.Args[2])-4:] != ".txt" {
		fmt.Println("ERROR: Files have to be .txt")
		return
	}
	input_file := os.Args[1]
	reader, err := fs.ReadFile(os.DirFS("."), input_file)
	if err != nil {
		fmt.Println("ERROR: Error opening input file")
		return
	}

	lexer := NewLexer(string(reader))
	lexer.Scan()

	// text alteration

	var modText [][]byte

	for i, item := range lexer.tokens {
		if item == 0 {
			NumOfIdens := modAnalyzer(lexer.tokenVals[i]) + 1
			if strings.Contains(lexer.tokenVals[i], "cap") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.Title(modText[j])
				}
			} else if strings.Contains(lexer.tokenVals[i], "up") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.ToUpper(modText[j])
				}
			} else if strings.Contains(lexer.tokenVals[i], "low") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.ToLower(modText[j])
				}
			} else if strings.Contains(lexer.tokenVals[i], "hex") {
				for j := i - NumOfIdens; j < i; j++ {
					bs, _ := hex.DecodeString(string(modText[j]))
					for _, item := range bs {
						number := fmt.Sprintf("%d", item)
						modText[j] = []byte(number)
					}
				}
			} else if strings.Contains(lexer.tokenVals[i], "bin") {
				for j := i - NumOfIdens; j < i; j++ {
					if string(modText[j]) != " " {
						decimal, _ := strconv.ParseUint(string(modText[j]), 2, 64)
						dec := fmt.Sprintf("%d", decimal)
						modText[j] = []byte(dec)
					} else {
						continue
					}
				}
			}
		}
		modText = append(modText, []byte(lexer.tokenVals[i]))
	}

	var prepText [][]byte
	copy(prepText, modText)
	reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)

	for _, item := range modText {
		match := reModifier.FindString(string(item))
		if match != "" {
			continue
		}
		prepText = append(prepText, item)
	}

	var finalText [][]byte
	prep := bytes.Join(prepText, []byte{})
	lexer2 := NewLexer((string(prep)))
	lexer2.Scan()

	for i, item := range lexer2.tokens {
		if i+1 < len(lexer2.tokens) && item == 4 {
			if lexer2.tokens[i+1] == 2 {
				continue
			}
		}
		if i+1 < len(lexer2.tokens) && item == 2 {
			if lexer2.tokens[i+1] != 4 {
				var slice [][]byte
				slice = append(slice, []byte(lexer2.tokenVals[i]))
				slice = append(slice, []byte(" "))
				jbs := bytes.Join(slice, []byte(""))
				finalText = append(finalText, jbs)
				continue
			}

		}
		if i+1 < len(lexer2.tokens) && item == 3 {
			wrappedText := wrapQoute(i, lexer2)
			finalText = append(finalText, wrappedText)
			break
		}

		if item == 4 && len(lexer2.tokenVals[i]) != 1 {
			finalText = append(finalText, []byte(" "))
			continue
		}
		if len(lexer2.tokenVals[i]) == 1 && lexer2.tokenVals[i][0] == 'a' {
			if lexer2.tokenVals[i+2][0] == 'a' || lexer2.tokenVals[i+2][0] == 'e' || lexer2.tokenVals[i+2][0] == 'i' || lexer2.tokenVals[i+2][0] == 'o' || lexer2.tokenVals[i+2][0] == 'u' {
				finalText = append(finalText, []byte("an"))
				continue
			}
		}
		if len(lexer2.tokenVals[i]) == 1 && lexer2.tokenVals[i][0] == 'A' {
			if lexer2.tokenVals[i+2][0] == 'a' || lexer2.tokenVals[i+2][0] == 'e' || lexer2.tokenVals[i+2][0] == 'i' || lexer2.tokenVals[i+2][0] == 'o' || lexer2.tokenVals[i+2][0] == 'u' {
				finalText = append(finalText, []byte("An"))
				continue
			}
		}
		finalText = append(finalText, []byte(lexer2.tokenVals[i]))
	}
	for _, item := range finalText {
		fmt.Printf("%s", item)
	}
	fmt.Println()
	file, _ := os.Create("./result.txt")
	for _, item := range prepText {
		file.Write(item)
	}
	// // // modText print
	// fmt.Println("Modified text")
	// for _, item := range modText {
	// 	fmt.Printf("%s", string(item))
	// }
	// fmt.Println()
	// fmt.Println()

	//Print the tokens
	// 	for i, token := range lexer2.tokens {
	// 		switch token {
	// 		case Modifier:
	// 			fmt.Printf("Modifier: %s\n", lexer2.tokenVals[i])
	// 		case Identifier:
	// 			fmt.Printf("Identifier: %s\n", lexer2.tokenVals[i])
	// 		case Whitespace:
	// 			fmt.Printf("Whitespace: %s\n", lexer2.tokenVals[i])
	// 		case Punct:
	// 			fmt.Printf("Punct: %s\n", lexer2.tokenVals[i])
	// 		case Quotemark:
	// 			fmt.Printf("Quotemark: %s\n", lexer2.tokenVals[i])
	// 		case Invalid:
	// 			fmt.Printf("Invalid: %s\n", lexer2.tokenVals[i])
	// 		}
	// 	}
}

// mod analyzer

func modAnalyzer(mod string) int {
	reArg := regexp.MustCompile(`\d+`)
	match := reArg.FindString(mod)
	if match == "" {
		return 1
	}
	NumOfIdens, _ := strconv.Atoi(match)
	return NumOfIdens * 2
}

var slice [][]byte

func wrapQoute(i int, l *Lexer) []byte {
	slice = append(slice, []byte(l.tokenVals[i]))
	for j := i + 1; j < len(l.tokens); j++ {
		if l.tokens[j] == 3 {
			slice = append(slice, []byte(l.tokenVals[j]))
			break
		}
		if l.tokens[j] == 4 {
			if l.tokens[j-1] == 3 || l.tokens[j+1] == 3 {
				continue
			}
		}
		slice = append(slice, []byte(l.tokenVals[j]))

	}

	return bytes.Join(slice, []byte(""))
}
