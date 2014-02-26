package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nint, _ := strconv.Atoi(os.Args[1])
	n := byte(nint)

	stack := [][]byte{}

	// initialize stack with first row values
	for i := byte(0); i < n; i++ {
		list := []byte{i}
		stack = append(stack, list)
	}

	count := 0
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := byte(0); i < n; i++ {
			if !bytes.Contains(cur, []byte{i}) && !checkDiagonals(cur, i) {
				if byte(len(cur)+1) == n {
					count++
				} else {
					test := make([]byte, len(cur))
					copy(test, cur)
					test = append(test, i)
					stack = append(stack, test)
				}
			}
		}
	}
	fmt.Println(n, ": ", count)
}

func checkDiagonals(list []byte, toCheck byte) bool {
	u := toCheck + byte(len(list))
	d := toCheck - byte(len(list))
	for _, x := range list {
		if u == x || d == x {
			return true
		}
		u--
		d++
	}
	return false
}
