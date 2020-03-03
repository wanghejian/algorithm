package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	l := list.New()
	for _, v := range s {
		switch v {
		case '(':
			l.PushBack(v)
		case '{':
			l.PushBack(v)
		case '[':
			l.PushBack(v)
		case ')':
			if l.Len() == 0 {
				return false
			}

			x := l.Back()
			xx := x.Value.(rune)
			if xx != '(' {
				return false
			}
			l.Remove(x)
		case '}':
			if l.Len() == 0 {
				return false
			}

			x := l.Back()
			xx := x.Value.(rune)
			if xx != '{' {
				return false
			}
			l.Remove(x)
		case ']':
			if l.Len() == 0 {
				return false
			}

			x := l.Back()
			xx := x.Value.(rune)
			if xx != '[' {
				return false
			}
			l.Remove(x)
		default:
		}
	}

	if l.Len() != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("(1)"))
	fmt.Println(isValid(" "))
	fmt.Println(isValid("(1[)]"))
	fmt.Println(isValid("([{232}])"))
	fmt.Println(isValid("("))
	fmt.Println(isValid(")"))
}
