package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"goreloaded"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	lexer := goreloaded.NewLexer(string(reader))
	lexer.Scan()

	// text alteration

	var modText [][]byte

	for i, item := range lexer.Tokens {
		if item == 0 {
			NumOfIdens := goreloaded.ModAnalyzer(lexer.TokenVals[i]) + 1
			if strings.Contains(lexer.TokenVals[i], "cap") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.Title(modText[j])
				}
			} else if strings.Contains(lexer.TokenVals[i], "up") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.ToUpper(modText[j])
				}
			} else if strings.Contains(lexer.TokenVals[i], "low") {
				for j := i - NumOfIdens; j < i; j++ {
					modText[j] = bytes.ToLower(modText[j])
				}
			} else if strings.Contains(lexer.TokenVals[i], "hex") {
				for j := i - NumOfIdens; j < i; j++ {
					bs, _ := hex.DecodeString(string(modText[j]))
					for _, item := range bs {
						number := fmt.Sprintf("%d", item)
						modText[j] = []byte(number)
					}
				}
			} else if strings.Contains(lexer.TokenVals[i], "bin") {
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
		modText = append(modText, []byte(lexer.TokenVals[i]))
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
	lexer2 := goreloaded.NewLexer((string(prep)))
	lexer2.Scan()

	for i, item := range lexer2.Tokens {
		if i+1 < len(lexer2.Tokens) && item == 4 {
			if lexer2.Tokens[i+1] == 2 {
				continue
			}
		}
		if i+1 < len(lexer2.Tokens) && item == 2 {
			if lexer2.Tokens[i+1] != 4 {
				var slice [][]byte
				slice = append(slice, []byte(lexer2.TokenVals[i]))
				slice = append(slice, []byte(" "))
				jbs := bytes.Join(slice, []byte(""))
				finalText = append(finalText, jbs)
				continue
			}

		}
		if i+1 < len(lexer2.Tokens) && item == 3 {
			wrappedText := goreloaded.WrapQoute(i, lexer2)
			finalText = append(finalText, wrappedText)
			break
		}

		if item == 4 && len(lexer2.TokenVals[i]) != 1 {
			finalText = append(finalText, []byte(" "))
			continue
		}
		if len(lexer2.TokenVals[i]) == 1 && lexer2.TokenVals[i][0] == 'a' {
			if lexer2.TokenVals[i+2][0] == 'a' || lexer2.TokenVals[i+2][0] == 'e' || lexer2.TokenVals[i+2][0] == 'i' || lexer2.TokenVals[i+2][0] == 'o' || lexer2.TokenVals[i+2][0] == 'u' || lexer2.TokenVals[i+2][0] == 'h' {
				finalText = append(finalText, []byte("an"))
				continue
			}
		}
		if len(lexer2.TokenVals[i]) == 1 && lexer2.TokenVals[i][0] == 'A' {
			if lexer2.TokenVals[i+2][0] == 'a' || lexer2.TokenVals[i+2][0] == 'e' || lexer2.TokenVals[i+2][0] == 'i' || lexer2.TokenVals[i+2][0] == 'o' || lexer2.TokenVals[i+2][0] == 'u' || lexer2.TokenVals[i+2][0] == 'h' {
				finalText = append(finalText, []byte("An"))
				continue
			}
		}
		finalText = append(finalText, []byte(lexer2.TokenVals[i]))
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

	//Print the Tokens
	// 	for i, token := range lexer2.Tokens {
	// 		switch token {
	// 		case Modifier:
	// 			fmt.Printf("Modifier: %s\n", lexer2.TokenVals[i])
	// 		case Identifier:
	// 			fmt.Printf("Identifier: %s\n", lexer2.TokenVals[i])
	// 		case Whitespace:
	// 			fmt.Printf("Whitespace: %s\n", lexer2.TokenVals[i])
	// 		case Punct:
	// 			fmt.Printf("Punct: %s\n", lexer2.TokenVals[i])
	// 		case Quotemark:
	// 			fmt.Printf("Quotemark: %s\n", lexer2.TokenVals[i])
	// 		case Invalid:
	// 			fmt.Printf("Invalid: %s\n", lexer2.TokenVals[i])
	// 		}
	// 	}
}

// mod analyzer
