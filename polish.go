package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	top  *Operand
	size int
}

type Operand struct {
	value int
	next  *Operand
}

func (s *Stack) Push(value int) {
	s.top = &Operand{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value int, err error) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value, nil
	}
	return -1, errors.New("nothing in stack")
}

func isOperation(op string) bool {
	switch op {
	case "*", "/", "+", "-":
		return true
	}
	return false
}

func calc(expr string) (result int) {
	var stack Stack
	for _, token := range strings.Fields(expr) {
		if isOperation(token) {
			op2, e := stack.Pop()
			if e != nil {
				log.Fatal(e)
			}
			op1, e := stack.Pop()
			if e != nil {
				log.Fatal(e)
			}

			var res int
			switch token {
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			}
			stack.Push(res)
		} else {
			op, e := strconv.Atoi(token)
			if e != nil {
				log.Fatal(e)
			}
			stack.Push(op)
		}
	}
	result, e := stack.Pop()
	if e != nil {
		log.Fatal(e)
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression in reverse polish notation: ")
	expr, _ := reader.ReadString('\n')

	fmt.Printf("Result: %v\n", calc(expr))
}
