package lexer

func Parse(content []byte) {
	for i, char := range content {
		if char == '\n' {
			continue
		}
		if char != ' ' {
			print(string(char))
			continue
		}
		if len(content)-1 == i+1 {
			continue
		}

		nextChar := content[i+1]

		if nextChar != ' ' {
			print(string(char))
		}
	}
}
