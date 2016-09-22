package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: \n")
	text, _ := reader.ReadString('\n')
	fmt.Println("Origional String: \n" + text)
	fmt.Println("Reversed String: ", reverse(text))
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
