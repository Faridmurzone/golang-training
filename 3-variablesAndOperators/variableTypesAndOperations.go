package main

import "fmt"

func main() {
	// Declaring variables with var and with :=, how are they printed
	var text string = "String variable using text"
	text2 := "String using :="
	var floatNumber float64 = 10.5
	var intNumber int = 10
	floatNumber2 := 10.
	truthy := true
	falsy := false
	myInteger := 50
	fmt.Println(text)
	fmt.Println(text2)
	fmt.Println("Other variable types: ", floatNumber, truthy, falsy, myInteger)
	fmt.Println("Try int operation 10 / 3 = ", intNumber/3)
	fmt.Println("Print a float64 operation 10 / 3 = ", floatNumber2/3)
	// Some conditional operators:
	x := true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)
	// public and Private functions
	Public()
	private()
	// With defer:
	functionWithDefer()
}

func private() {
	fmt.Println("This is a private function")
}

func Public() {
	// Public function
	fmt.Println("This is a public function")
}

func functionWithDefer() {
	defer fmt.Print("And this is wrote first but printed later")
	fmt.Print("This will be printed first ")
}
