package main

import (
	"fmt"
	"unicode/utf8"
)

const minStrLen = 7

func main() {
	s := "Шла Саша по шоссе"
	r := []rune(s)
	if utf8.RuneCount([]byte(s)) < minStrLen {
		fmt.Printf("The length of string is shorter than %d\n", minStrLen)
		return
	}
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s);
	fmt.Printf("After : %s\n", string(r))
}
