package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	stack := stack.New() // create a stack variable of type Stack

	if len(s)%2 != 0 {
		return false
	}

	for _, value := range s {
		if stack.Len() == 0 && (string(value) == ")" || string(value) == "]" || string(value) == "}") {
			return false
		} else {

			if (stack.Peek() == "(" && string(value) == ")") || (stack.Peek() == "{" && string(value) == "}") || (stack.Peek() == "[" && string(value) == "]") {
				stack.Pop()
			} else {
				stack.Push(string(value))
			}

		}

	}

	return stack.Len() == 0

}

func main() {
	fmt.Println(isValid("()[]{}"))
}
