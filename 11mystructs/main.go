package main

import "fmt"

func main() {
	fmt.Println("Struts in golang")
	// no inheritance in golang; No super or parent

	deepak := User{"Deepak", "Deepak@dev", true, 16}
	fmt.Println(deepak)
	fmt.Printf("Deepak details are: %+v\n", deepak)
	fmt.Printf("Name is %v and email is %v", deepak.Name, deepak.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
