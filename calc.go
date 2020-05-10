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
	values := strings.Split(operation, "+")
	fmt.Println("Split: ", values)
	fmt.Println("Concat: ", values[0]+values[1])
	operator1, err1 := strconv.Atoi(values[0])
	operator2, err2 := strconv.Atoi(values[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Errores: ", err1, err2)
	} else {
		fmt.Println("Add: ", operator1+operator2)
	}
}
