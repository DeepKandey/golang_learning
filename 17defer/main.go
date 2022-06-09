package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer concept")
	printInt()

}

func printInt() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
