package main

import (
	"bytes"
	"fmt"
	"strings"
)

// splitNth takes a string and splits it every n characters.
func splitNth(s string, n int) []string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1

	// Insert a space every N characters
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteRune(' ')
		}
	}

	// Split the string on the empty space and return array
	return strings.Split(buffer.String(), " ")
}

func main() {
	fmt.Println(splitNth("0C1F6EA7EEB20C6A625F01919894A74418D306828316CE1B8B2DD1F2D7C6EFB5", 8)) // [0C1F6EA7 EEB20C6A 625F0191 9894A744 18D30682 8316CE1B 8B2DD1F2 D7C6EFB5]
}
