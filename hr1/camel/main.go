package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	count := 1
	for _, ch := range input {
		str := string(ch)
		if str == strings.ToUpper(str) {
			count++
		}
	}
	fmt.Println(count)
}


