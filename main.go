package main

import (
	"fmt"
	"strings"
)

var string_1 = "Nikita"

func main() {
	string_2 := "Liza"
	vowels := "euioa"
	for _, val := range string_1 {
		if !strings.ContainsRune(vowels, val) {
			fmt.Println(string(val))
		}

	}
	for _, val := range string_2 {
		fmt.Printf("%c", val)
	}
}
