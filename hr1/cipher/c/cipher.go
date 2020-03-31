package cipher

import (
	"fmt"
	"strings"
)

func cipher() {
	var input, output string
	var inpLen, key int
	c := make(map[string]string)
    alphabet := "abcdefghijklmnopqrstuvwxyz"
	
	n, e := fmt.Scanf("%d\n%s\n%d\n", &inpLen, &input, &key)
	if e != nil && n != 3 {
		fmt.Printf("Scan of input didn't go well mate, %s", e)
	}
	
	for i := 0; i < len(alphabet); i++ {
		if i < len(alphabet)-key {
			c[string(alphabet[i])] = string(alphabet[i+key])
		} else {
			c[string(alphabet[i])] = string(alphabet[key-(len(alphabet)-i)])
		}
	}

	for _, ch := range input {
		if strings.Contains(string(ch), "-") == true {
			output += string(ch)
		} else if string(ch) == strings.ToUpper(string(ch)) {
			str := c[strings.ToLower(string(ch))]
			output += strings.ToUpper(str)
		} else {
			output += c[(string(ch))]
		}
	}
	fmt.Println(output)
}