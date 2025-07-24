package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// convertToUnicodeBold converts ASCII letters and digits to their Unicode bold equivalents
func convertToUnicodeBold(s string) string {
	var b strings.Builder

	for _, r := range s {
		switch {
		// Uppercase letters A-Z
		case 'A' <= r && r <= 'Z':
			// Unicode Mathematical Bold uppercase letters start at U+1D400
			b.WriteRune(rune(r - 'A' + 0x1D400))

		// Lowercase letters a-z
		case 'a' <= r && r <= 'z':
			// Unicode Mathematical Bold lowercase letters start at U+1D41A
			b.WriteRune(rune(r - 'a' + 0x1D41A))

		// Digits 0-9
		case '0' <= r && r <= '9':
			// Unicode Mathematical Bold digits start at U+1D7CE
			b.WriteRune(rune(r - '0' + 0x1D7CE))

		// Keep spaces as is
		case unicode.IsSpace(r):
			b.WriteRune(r)

		// For any other characters, keep them unchanged
		default:
			b.WriteRune(r)
		}
	}

	return b.String()
}

func main() {
	fmt.Println("Enter text to convert to Unicode Bold:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	boldText := convertToUnicodeBold(input)
	fmt.Println("Unicode Bold output:")
	fmt.Println(boldText)
}
