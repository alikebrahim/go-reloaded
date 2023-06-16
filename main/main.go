package main

import (
	"fmt"
	"goreloaded"
	"io/fs"
	"os"
)

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

	lexer := goreloaded.NewLexer(reader)
	lexer.Scan()

	// text alteration
	goreloaded.ModsMap(lexer)
	goreloaded.ModEdit(&goreloaded.Text, lexer)
	goreloaded.TextFmt(&goreloaded.FmtText)

	//goreloaded.TextMod(&text, lexer)
	fmt.Println("Before format:")
	fmt.Println()
	for _, item := range goreloaded.Text {
		fmt.Printf("%s", item)
	}
	fmt.Println()

	fmt.Println("after format:")
	fmt.Println()
	for _, item := range goreloaded.FmtText {
		fmt.Printf("%s", item)
	}
	fmt.Println()

	//Print the Tokens
	//prototypr for token printing

	// for i, token := range lexer.Tokens {
	// 	switch token {
	// 	case goreloaded.Modifier:
	// 		fmt.Printf("Modifier: %s\n", lexer.TokenVals[i])
	// 	case goreloaded.Identifier:
	// 		fmt.Printf("Identifier: %s\n", lexer.TokenVals[i])
	// 	case goreloaded.Whitespace:
	// 		fmt.Printf("Whitespace: %s\n", lexer.TokenVals[i])
	// 	case goreloaded.Punct:
	// 		fmt.Printf("Punct: %s\n", lexer.TokenVals[i])
	// 	case goreloaded.Quotemark:
	// 		fmt.Printf("Quotemark: %s\n", lexer.TokenVals[i])
	// 	case goreloaded.Invalid:
	// 		fmt.Printf("Invalid: %s\n", lexer.TokenVals[i])
	// 	}
	// }
}
