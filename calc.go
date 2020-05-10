package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Write an addition and press (enter): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println("You wrote: ", operation)
	operator := "+"
	values := strings.Split(operation, operator)
	fmt.Println("Split: ", values)
	fmt.Println("Concat: ", values[0]+values[1])
	operator1, err1 := strconv.Atoi(values[0])
	operator2, err2 := strconv.Atoi(values[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Errores: ", err1, err2)
	} else {
		switch operator {
		case "+":
			fmt.Println("Add: ", operator1+operator2)
		case "-":
			fmt.Println("Sustract: ", operator1-operator2)
		case "*":
			fmt.Println("Multiply: ", operator1+operator2)
		case "/":
			fmt.Println("Division: ", operator1/operator2)
		default:
			fmt.Println("Operation not supported")
		}
	}
}
