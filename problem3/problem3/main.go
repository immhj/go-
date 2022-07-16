package main

import (
	"fmt"
)

func main() {
	var sa string
	var n int
	fmt.Scanln(&n)
	fmt.Scanln(&sa)
	for i := 0; i < len(sa); i++ {
		if (int(sa[i]) + n) > 'z' {
			fmt.Print(string(int(sa[i]) + n - 1 - 'z' + 'a'))
		} else {
			fmt.Print(string(int(sa[i]) + n))
		}
	}
}
