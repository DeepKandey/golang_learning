package main

import "fmt"

func main() {
	fmt.Println("Struts in golang")
	// no inheritance in golang; No super or parent

	deepak := User{"Deepak", "Deepak@dev", true, 16}
	fmt.Println(deepak)
	fmt.Printf("Deepak details are: %+v\n", deepak)
	fmt.Printf("Name is %v and email is %v\n", deepak.Name, deepak.Email)

	deepak.GetStatus()
	deepak.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
