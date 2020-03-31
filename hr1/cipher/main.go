package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	output := ""
	var l, key int
	n, e := fmt.Scanf("%d\n%s\n%d", &l, &s, &key)
	if n != 3 || e != nil {
		fmt.Printf("Error scanning the input, %s", e)
	}

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	for _, ch := range s {
		if strings.IndexRune(string(alphabet), ch)>0{
			output = output + string(rotate(ch, key, alphabet))
		} else {
			output += string(ch)
		}
	}

	fmt.Println(output)

}

func rotate(s rune, delta int, orig []rune) rune {
	idx := strings.IndexRune(string(orig), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(orig)
	return orig[idx]
}
