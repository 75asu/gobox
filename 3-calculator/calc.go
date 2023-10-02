package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "Add two numbers")
	sub := flag.Bool("sub", false, "Subtract two numbers")
	mul := flag.Bool("mul", false, "Multiply two numbers")
	div := flag.Bool("div", false, "Divide two numbers")

	flag.Parse()

	var a, b int
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&b)

	switch {
	case *add:
		fmt.Printf("Additon: %d \n", addition(a, b))
	case *sub:
		fmt.Printf("Difference: %d \n", subtract(a, b))
	case *mul:
		fmt.Printf("Product: %d \n", multiply(a, b))
	case *div:
		fmt.Printf("Division: %d \n", division(a, b))
	default:
		fmt.Fprintln(os.Stderr, "Wrong option, try with add, sub, div and mul")
		os.Exit(1)
	}

}

func addition(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Can't Divide By Zero")
		os.Exit(1)
	}
	return a / b
}
