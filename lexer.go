package goreloaded

import "regexp"

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
	Input     string
	Position  int
	Tokens    []Token
	TokenVals []string
}

func NewLexer(Input string) *Lexer {
	return &Lexer{
		Input:     Input,
		Position:  0,
		Tokens:    []Token{},
		TokenVals: []string{},
	}
}

func (l *Lexer) Scan() {
	reIdentifier := regexp.MustCompile(`[a-zA-Z0-9_]\w*`)
	reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)
	rePunct := regexp.MustCompile(`[.,?!:;]+`)
	reQuotemark := regexp.MustCompile(`['|"]`)
	reWhitespace := regexp.MustCompile(`\s+`)

	for l.Position < len(l.Input) {
		substr := l.Input[l.Position:]

		if match := reIdentifier.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Identifier)
			l.TokenVals = append(l.TokenVals, reIdentifier.FindString(substr))
			l.Position += match[1]
		} else if match := reModifier.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Modifier)
			l.TokenVals = append(l.TokenVals, reModifier.FindString(substr))
			l.Position += match[1]
		} else if match := reWhitespace.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Whitespace)
			l.TokenVals = append(l.TokenVals, reWhitespace.FindString(substr))
			l.Position += match[1]
		} else if match := rePunct.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Punct)
			l.TokenVals = append(l.TokenVals, rePunct.FindString(substr))
			l.Position += match[1]
		} else if match := reQuotemark.FindStringIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Quotemark)
			l.TokenVals = append(l.TokenVals, reQuotemark.FindString(substr))
			l.Position += match[1]
		} else {
			l.Tokens = append(l.Tokens, Invalid)
			l.TokenVals = append(l.TokenVals, string(substr[0]))
			l.Position++
		}
	}
}
