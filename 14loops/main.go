package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("---------------")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("---------------")

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			// rougeValue++
			// continue
			break
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}
