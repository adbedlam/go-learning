package main

import (
	"fmt"
	"strings"	
)

func main() {
	var stroka string = "I love Lisa"

	var vowel string = "qeyuioa"

	for _, k := range stroka{
		if !strings.ContainsRune(vowel, k){
			fmt.Printf("%c",k)
		}
	}
}