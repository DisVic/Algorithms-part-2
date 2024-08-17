package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.TrimSpace(s)
	str := strings.Split(s, " ")
	fmt.Println(PolishCalculator(str))
}

func operations(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return b - a
	case "*":
		return b * a
	case "/":
		if b < a {
			return b % a
		} else {
			return b / a
		}
	}
	return 0
}

func twoOperands(stack *[]int) (int, int) {
	a := (*stack)[len(*stack)-1]
	b := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]
	return a, b
}

func PolishCalculator(str []string) int {
	stack := []int{}
	for _, ch := range str {
		switch ch {
		case "+", "-", "*", "/":
			a, b := twoOperands(&stack)
			stack = append(stack, operations(ch, a, b))
		default:
			num, _ := strconv.Atoi(ch)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
