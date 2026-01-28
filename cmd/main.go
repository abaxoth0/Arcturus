package main

import (
	"arcturus/packages/lexer"
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

	lexer.Parse(content)

	println()
	println("------------------")
	println("DONE")
}
