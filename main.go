package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

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

	split_txt := strings.Split(string(reader), " ")
	for _, item := range split_txt {
		fmt.Println(item)
	}
	fmt.Println(split_txt)
}
