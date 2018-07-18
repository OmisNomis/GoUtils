package main

import (
	"fmt"
)

func reverseArray(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(reverseArray([]string{"0C1F6EA7", "EEB20C6A", "625F0191", "9894A744", "18D30682", "8316CE1B", "8B2DD1F2", "D7C6EFB5"})) // [D7C6EFB5 8B2DD1F2 8316CE1B 18D30682 9894A744 625F0191 EEB20C6A 0C1F6EA7]
}
