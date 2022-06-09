package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(3, 5, 9, 34)
	fmt.Println("proResult is: ", proResult)
}

func adder(val1 int, val2 int) int {
	return val1 + val2

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}
