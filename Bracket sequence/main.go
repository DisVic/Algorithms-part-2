package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var str string
	_, _ = fmt.Scan(&str)
	t := time.Now()
	is_correct_bracket_seq(str)
	fmt.Print(time.Since(t))
}

func is_correct_bracket_seq(s string) {
	if len(s) == 0 {
		fmt.Print("True")
		return
	}
	if len(s)%2 != 0 {
		fmt.Print("False")
		return
	}
	s = strings.ReplaceAll(s, "}", "{")
	s = strings.ReplaceAll(s, "]", "[")
	s = strings.ReplaceAll(s, ")", "(")
	for i := 0; i < len(s)/2; i++ {
		go func(i int) {
			if s[i] != s[len(s)-1-i] {
				fmt.Print("False")
				return
			}
		}(i)
	}
	fmt.Print("True")
}
