package goreloaded

import "regexp"

type Token int

const (
	Modifier Token = iota
	Identifier
	Punct
	Quotedtext
	Whitespace
	Quotemark
	Invalid
)

type Lexer struct {
	Input     []byte
	Position  int
	Tokens    []Token
	TokenVals [][]byte
}

func NewLexer(Input []byte) *Lexer {
	return &Lexer{
		Input:     Input,
		Position:  0,
		Tokens:    []Token{},
		TokenVals: [][]byte{},
	}
}

func (l *Lexer) Scan() {
	reIdentifier := regexp.MustCompile(`[a-zA-Z0-9_]*[']?[a-zA-Z0-9_]\w*[']?`)
	reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)
	rePunct := regexp.MustCompile(`[.,?!:;]+`)
	reQuotedtext := regexp.MustCompile(`(?U)'(.*[a-zA-Z]'[a-zA-Z])([^']*)'`)
	reWhitespace := regexp.MustCompile(`\s+`)

	for l.Position < len(l.Input) {
		substr := l.Input[l.Position:]

		if match := reIdentifier.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Identifier)
			l.TokenVals = append(l.TokenVals, reIdentifier.Find(substr))
			l.Position += match[1]
		} else if match := reModifier.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Modifier)
			l.TokenVals = append(l.TokenVals, reModifier.Find(substr))
			l.Position += match[1]
		} else if match := reWhitespace.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Whitespace)
			l.TokenVals = append(l.TokenVals, reWhitespace.Find(substr))
			l.Position += match[1]
		} else if match := rePunct.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Punct)
			l.TokenVals = append(l.TokenVals, rePunct.Find(substr))
			l.Position += match[1]
		} else if match := reQuotedtext.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Quotedtext)
			l.TokenVals = append(l.TokenVals, reQuotedtext.Find(substr))
			l.Position += match[1]
		} else {
			l.Tokens = append(l.Tokens, Invalid)
			l.TokenVals = append(l.TokenVals, []byte(substr))
			l.Position++
		}
	}
}

func (l *Lexer) QuoteFmtScan() {
	reIdentifier := regexp.MustCompile(`[a-zA-Z0-9_]*[']?[a-zA-Z0-9_]\w*[']?`)
	reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)
	rePunct := regexp.MustCompile(`[.,?!:;]+`)
	reQuotemark := regexp.MustCompile(`['|"]`)
	reWhitespace := regexp.MustCompile(`\s+`)

	for l.Position < len(l.Input) {
		substr := l.Input[l.Position:]

		if match := reIdentifier.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Identifier)
			l.TokenVals = append(l.TokenVals, reIdentifier.Find(substr))
			l.Position += match[1]
		} else if match := reModifier.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Modifier)
			l.TokenVals = append(l.TokenVals, reModifier.Find(substr))
			l.Position += match[1]
		} else if match := reWhitespace.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Whitespace)
			l.TokenVals = append(l.TokenVals, reWhitespace.Find(substr))
			l.Position += match[1]
		} else if match := rePunct.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Punct)
			l.TokenVals = append(l.TokenVals, rePunct.Find(substr))
			l.Position += match[1]
		} else if match := reQuotemark.FindIndex(substr); match != nil && match[0] == 0 {
			l.Tokens = append(l.Tokens, Quotemark)
			l.TokenVals = append(l.TokenVals, reQuotemark.Find(substr))
			l.Position += match[1]
		} else {
			l.Tokens = append(l.Tokens, Invalid)
			l.TokenVals = append(l.TokenVals, []byte(substr))
			l.Position++
		}
	}
}
