package reversestr

import "fmt"

func ReverseStr(s string) {
	for i := len(s); i > 0; i-- {
		fmt.Print(string(s[i-1]))
	}
}
