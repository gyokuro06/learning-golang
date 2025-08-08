package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])
}
