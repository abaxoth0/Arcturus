package main

import (
	"arcturus/packages/lexer"
	"fmt"
	"io"
	"os"
)


func main() {
	file, err := os.OpenFile("example.arc", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	tokens := lexer.Parse(content)

	for _, tk := range tokens {
		r := tk.Raw()
		fmt.Printf("%s ", r)
	}

	println()
}
