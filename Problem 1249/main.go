package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0)
	var result string = s

	for i := 0; i < len(result); i++ {

		value := result[i]

		if value != '(' && rune(value) != ')' {
			continue
		}

		if value == '(' {
			stack = append(stack, i)
		} else if value == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				result = result[:i] + result[i+1:]
				i--
			}
		}
	}

	for len(stack) > 0 {
		index := stack[len(stack)-1]
		result = result[:index] + result[index+1:]
		stack = stack[:len(stack)-1]
	}

	return result
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
}
