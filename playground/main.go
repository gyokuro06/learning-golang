package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	messages := "Hi 👩 and 👨"

	r, _ := utf8.DecodeRuneInString(messages[3:])
	fmt.Println(string(r))
}
