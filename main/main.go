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
	var text [][]byte
	goreloaded.ModsMap(lexer)
	goreloaded.ModEdit(&text, lexer)
	goreloaded.TextFmt(&goreloaded.FmtText)

	//goreloaded.TextMod(&text, lexer)
	for _, item := range text {
		fmt.Printf("%s", item)
	}
	fmt.Println()
	//lastb := fmt.Sprintf("%v", text[len(text)-1])
	//var modText [][]byte

	// var prepText [][]byte
	// copy(prepText, modText)
	// reModifier := regexp.MustCompile(`\([hex|bin|cap|low|up]+(?:,\s\d)?\)`)

	// for _, item := range modText {
	// 	match := reModifier.FindString(string(item))
	// 	if match != "" {
	// 		continue
	// 	}
	// 	prepText = append(prepText, item)
	// }

	// var finalText [][]byte
	// prep := bytes.Join(prepText, []byte{})
	// lexer2 := goreloaded.NewLexer((string(prep)))
	// lexer2.Scan()

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
	// for _, item := range finalText {
	// 	fmt.Printf("%s", item)
	// }
	// fmt.Println()
	// file, _ := os.Create("./result.txt")
	// for _, item := range prepText {
	// 	file.Write(item)
	// }
	// // // modText print
	// fmt.Println("Modified text")
	// for _, item := range modText {
	// 	fmt.Printf("%s", string(item))
	// }
	// fmt.Println()
	// fmt.Println()

	//Print the Tokens
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
