package main

import (
	"bytes"
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
	reQuotemark := regexp.MustCompile(`'`)
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
	// == two arguments are passed of type .txt
	// & input file == exists
	if len(os.Args) != 3 {
		fmt.Println("Usage requires input.txt and output.txt")
	}
	if os.Args[1][len(os.Args[1])-4:] != ".txt" || os.Args[2][len(os.Args[2])-4:] != ".txt" {
		fmt.Println("Files have to be .txt")
	}
	input_file := os.Args[1]
	reader, err := fs.ReadFile(os.DirFS("."), input_file)
	if err != nil {
		fmt.Println("Error opening input file")
	}

	lexer := NewLexer(string(reader))
	lexer.Scan()

	// TODO: text alteration

	// Original txt print
	fmt.Println("Original text:")
	for _, item := range lexer.tokenVals {
		fmt.Printf("%s", item)
	}
	fmt.Println()
	fmt.Println()

	// modified text assembly and print
	// // mod analyzer

	// // text modification
	var modText [][]byte

	for i, item := range lexer.tokens {
		if item == 0 {
			args := modAnalyzer(lexer.tokenVals[i]) + 1
			//subtxt := modText[i-args : i]
			if strings.Contains(lexer.tokenVals[i], "cap") {
				for j := i - args; j < i; j++ {
					modText[j] = bytes.Title(modText[j])
				}
			} else if strings.Contains(lexer.tokenVals[i], "up") {
				for j := i - args; j < i; j++ {
					modText[j] = bytes.ToUpper(modText[j])
				}
			} else if strings.Contains(lexer.tokenVals[i], "low") {
				for j := i - args; j < i; j++ {
					modText[j] = bytes.ToLower(modText[j])
				}
			}
			continue
		}
		modText = append(modText, []byte(lexer.tokenVals[i]))
	}

	// posAdjust := 0
	// for i, item := range lexer.tokens {
	// 	if item == 0 {
	// 		args := modAnalyzer(lexer.tokenVals[i])
	// 		for j := args; j > 0; j-- {
	// 			fmt.Printf("i:%d <> j:%d <> args:%d <> posAdjust:%d <> modifier:%s\n", i, j, args, posAdjust, lexer.tokenVals[i])
	// 			fmt.Println(string(modText[i-j-1-posAdjust]))
	// 		}
	// 		posAdjust++
	// 		continue
	// 	}
	// 	modText = append(modText, []byte(lexer.tokenVals[i]))
	// }
	//	modText[0] = bytes.ToUpper(modText[0])

	// // modText print
	fmt.Println("Modified text")
	for _, item := range modText {
		fmt.Printf("%s", string(item))
	}
	fmt.Println()
	fmt.Println()

	// Print the tokens
	// for i, token := range lexer.tokens {
	// 	switch token {
	// 	case Modifier:
	// 		fmt.Printf("Modifier: %s\n", lexer.tokenVals[i])
	// 	case Identifier:
	// 		fmt.Printf("Identifier: %s\n", lexer.tokenVals[i])
	// 	case Whitespace:
	// 		fmt.Printf("Whitespace: %s\n", lexer.tokenVals[i])
	// 	case Punct:
	// 		fmt.Printf("Punct: %s\n", lexer.tokenVals[i])
	// 	case Quotemark:
	// 		fmt.Printf("Quotemark: %s\n", lexer.tokenVals[i])
	// 	case Invalid:
	// 		fmt.Printf("Invalid: %s\n", lexer.tokenVals[i])
	// 	}
	// }
}

// mod analyzer
func modAnalyzer(mod string) int {
	reArg := regexp.MustCompile(`\d+`)
	match := reArg.FindString(mod)
	if match == "" {
		return 1
	}
	args, _ := strconv.Atoi(match)
	return args * 2
}
