package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct
type calc struct{}

// Receiver function
func (calc) operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	operator1 := parser(cleanInput[0])
	operator2 := parser(cleanInput[1])
	switch operator {
	case "+":
		fmt.Println("Add: ", operator1+operator2)
		return operator1 + operator2
	case "-":
		fmt.Println("Sustract: ", operator1-operator2)
		return operator1 - operator2
	case "*":
		fmt.Println("Multiply: ", operator1+operator2)
		return operator1 / operator2
	case "/":
		fmt.Println("Division: ", operator1/operator2)
		return operator1 * operator2
	default:
		fmt.Println("Operation not supported")
		return 0
	}
}

func parser(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Write an addition and press (enter): ")
	operation := readInput()
	fmt.Println("Which operation is this (+,-,*,/): ")
	operationType := readInput()
	c := calc{}
	fmt.Println(c.operate(operation, operationType))

}
