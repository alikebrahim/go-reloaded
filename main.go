package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

type Token int

const (
	Modifier Token = iota
	Number
	Identifier
	Punct
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
		fmt.Println("Usage requires INPUTFILE.txt and OUTPUTFILE.txt")
	}
	if os.Args[1][len(os.Args[1])-4:] != ".txt" || os.Args[2][len(os.Args[2])-4:] != ".txt" {
		fmt.Println("Files have to be .txt")
	}
	input_file := os.Args[1]
	reader, err := fs.ReadFile(os.DirFS("."), input_file)
	if err != nil {
		fmt.Println("Error opening input file")
	}
}
